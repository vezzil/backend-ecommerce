package service

import (
	"errors"
	"fmt"
	"strconv"

	"backend-ecommerce/internal/application/entity"
	"backend-ecommerce/internal/application/repository"
	"backend-ecommerce/internal/infrastructure/stripe"
	"github.com/stripe/stripe-go/v76"
)

type paymentService struct {
	repo         repository.PaymentRepository
	stripeClient *stripe.StripeManager
}

// NewPaymentService creates a new payment service
func NewPaymentService(repo repository.PaymentRepository, stripeClient *stripe.StripeManager) PaymentService {
	return &paymentService{
		repo:         repo,
		stripeClient: stripeClient,
	}
}

// ProcessPayment processes a payment using Stripe
func (s *paymentService) ProcessPayment(payment entity.Payment) (*entity.Payment, error) {
	// Convert amount to cents (Stripe uses smallest currency unit)
	amount := int64(payment.Amount * 100)

	// Create a payment intent
	intent, err := s.stripeClient.CreatePaymentIntent(amount, string(payment.Currency), payment.OrderID)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment intent: %w", err)
	}

	// Update payment with Stripe details
	payment.TransactionID = intent.ID
	payment.Status = string(intent.Status)
	payment.Metadata = map[string]string{
		"stripe_payment_intent_id": intent.ID,
		"client_secret":           intent.ClientSecret,
	}

	// Save payment to database
	return s.repo.Create(payment)
}

// GetPaymentByID retrieves a payment by its ID
func (s *paymentService) GetPaymentByID(id string) (*entity.Payment, error) {
	return s.repo.FindByID(id)
}

// GetPaymentsByOrderID retrieves all payments for a specific order
func (s *paymentService) GetPaymentsByOrderID(orderID string) ([]entity.Payment, error) {
	return s.repo.FindByOrderID(orderID)
}

// RefundPayment processes a refund for a payment
func (s *paymentService) RefundPayment(paymentID string, amount float64) (*entity.Payment, error) {
	payment, err := s.repo.FindByID(paymentID)
	if err != nil {
		return nil, err
	}

	// Check if payment is eligible for refund
	if payment.Status != "succeeded" {
		return nil, errors.New("only succeeded payments can be refunded")
	}

	// Convert amount to cents for Stripe
	amountInCents := int64(amount * 100)

	// Process refund with Stripe
	_, err = s.stripeClient.CreateRefund(payment.TransactionID, amountInCents)
	if err != nil {
		return nil, fmt.Errorf("failed to process refund: %w", err)
	}

	// Update payment status
	payment.Status = "refunded"
	if amount > 0 && amount < payment.Amount {
		payment.Status = "partially_refunded"
	}

	// Save updated payment
	return s.repo.Update(*payment)
}

// CreateCheckoutSession creates a new Stripe Checkout session
func (s *paymentService) CreateCheckoutSession(orderID string, amount float64, currency string, customerEmail string) (map[string]interface{}, error) {
	// Convert amount to cents for Stripe
	amountInCents := int64(amount * 100)

	// Create checkout session with Stripe
	session, err := s.stripeClient.CreateCheckoutSession(
		amountInCents,
		currency,
		"http://localhost:8080/api/payments/success", // TODO: Make these URLs configurable
		"http://localhost:8080/api/payments/cancel",
		orderID,
		customerEmail,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create checkout session: %w", err)
	}

	// Return session URL and ID for client-side redirection
	return map[string]interface{}{
		"session_id":       session.ID,
		"session_url":      session.URL,
		"payment_intent":   session.PaymentIntent,
		"customer_email":   customerEmail,
		"amount":           amount,
		"currency":         currency,
		"order_id":         orderID,
		"payment_status":   session.PaymentStatus,
		"client_reference": session.ClientReferenceID,
	}, nil
}

// HandleWebhook handles Stripe webhook events
func (s *paymentService) HandleWebhook(payload []byte, signature string) error {
	event, err := s.stripeClient.HandleWebhook(payload, signature)
	if err != nil {
		return fmt.Errorf("error verifying webhook signature: %w", err)
	}

	switch event.Type {
	case "payment_intent.succeeded":
		return s.handlePaymentIntentSucceeded(event)
	case "payment_intent.payment_failed":
		return s.handlePaymentIntentFailed(event)
	case "charge.refunded":
		return s.handleChargeRefunded(event)
	}

	return nil
}

// handlePaymentIntentSucceeded handles successful payment intents
func (s *paymentService) handlePaymentIntentSucceeded(event stripe.Event) error {
	var paymentIntent stripe.PaymentIntent
	err := json.Unmarshal(event.Data.Raw, &paymentIntent)
	if err != nil {
		return fmt.Errorf("error parsing payment intent: %w", err)
	}

	// Update payment status in database
	payment, err := s.repo.FindByTransactionID(paymentIntent.ID)
	if err != nil {
		return fmt.Errorf("error finding payment: %w", err)
	}

	payment.Status = string(paymentIntent.Status)
	_, err = s.repo.Update(*payment)
	return err
}

// handlePaymentIntentFailed handles failed payment intents
func (s *paymentService) handlePaymentIntentFailed(event stripe.Event) error {
	var paymentIntent stripe.PaymentIntent
	err := json.Unmarshal(event.Data.Raw, &paymentIntent)
	if err != nil {
		return fmt.Errorf("error parsing payment intent: %w", err)
	}

	// Update payment status in database
	payment, err := s.repo.FindByTransactionID(paymentIntent.ID)
	if err != nil {
		return fmt.Errorf("error finding payment: %w", err)
	}

	payment.Status = "failed"
	if len(paymentIntent.LastPaymentError.Message) > 0 {
		payment.Metadata["error"] = paymentIntent.LastPaymentError.Message
	}

	_, err = s.repo.Update(*payment)
	return err
}

// handleChargeRefunded handles refund events
func (s *paymentService) handleChargeRefunded(event stripe.Event) error {
	var charge stripe.Charge
	err := json.Unmarshal(event.Data.Raw, &charge)
	if err != nil {
		return fmt.Errorf("error parsing charge: %w", err)
	}

	// Update payment status in database
	payment, err := s.repo.FindByTransactionID(charge.PaymentIntent.ID)
	if err != nil {
		return fmt.Errorf("error finding payment: %w", err)
	}

	if charge.Refunded {
		payment.Status = "refunded"
	} else if len(charge.Refunds.Data) > 0 {
		payment.Status = "partially_refunded"
	}

	_, err = s.repo.Update(*payment)
	return err
}

// ConvertToStripeAmount converts a float amount to Stripe's smallest currency unit (e.g., cents)
func (s *paymentService) ConvertToStripeAmount(amount float64, currency string) (int64, error) {
	// Handle zero-decimal currencies
	zeroDecimalCurrencies := map[string]bool{
		"BIF": true, "CLP": true, "DJF": true, "GNF": true,
		"JPY": true, "KMF": true, "KRW": true, "MGA": true,
		"PYG": true, "RWF": true, "UGX": true, "VND": true,
		"VUV": true, "XAF": true, "XOF": true, "XPF": true,
	}

	if zeroDecimalCurrencies[currency] {
		return int64(amount), nil
	}

	// For other currencies, convert to cents
	return int64(amount * 100), nil
}

// ConvertFromStripeAmount converts from Stripe's smallest currency unit to a float amount
func (s *paymentService) ConvertFromStripeAmount(amount int64, currency string) float64 {
	// Handle zero-decimal currencies
	zeroDecimalCurrencies := map[string]bool{
		"BIF": true, "CLP": true, "DJF": true, "GNF": true,
		"JPY": true, "KMF": true, "KRW": true, "MGA": true,
		"PYG": true, "RWF": true, "UGX": true, "VND": true,
		"VUV": true, "XAF": true, "XOF": true, "XPF": true,
	}

	if zeroDecimalCurrencies[currency] {
		return float64(amount)
	}

	// For other currencies, convert from cents
	return float64(amount) / 100.0
}
