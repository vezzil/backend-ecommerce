package entity

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// PaymentStatus represents the status of a payment
type PaymentStatus string

const (
	PaymentStatusInitiated PaymentStatus = "initiated"
	PaymentStatusSucceeded PaymentStatus = "succeeded"
	PaymentStatusFailed    PaymentStatus = "failed"
	PaymentStatusRefunded  PaymentStatus = "refunded"
)

// RawResponse is a custom type to handle JSON data in the database
type RawResponse map[string]interface{}

// Value implements the driver.Valuer interface
func (r RawResponse) Value() (driver.Value, error) {
	return json.Marshal(r)
}

// Scan implements the sql.Scanner interface
func (r *RawResponse) Scan(value interface{}) error {
	if value == nil {
		*r = nil
		return nil
	}
	return json.Unmarshal(value.([]byte), r)
}

// Payment represents a payment transaction
type Payment struct {
	ID                string        `json:"id" gorm:"primaryKey;column:id;type:varchar(36);default:(UUID());comment:'Primary Key'"`
	OrderID           string        `json:"order_id" gorm:"column:order_id;type:varchar(36);not null;comment:'FK to order'"`
	Provider          string        `json:"provider" gorm:"column:provider;type:varchar(50);comment:'Payment provider (e.g., stripe, paypal)'"`
	ProviderPaymentID string        `json:"provider_payment_id,omitempty" gorm:"column:provider_payment_id;type:varchar(255);comment:'Payment ID from provider'"`
	Amount            float64       `json:"amount" gorm:"column:amount;type:decimal(12,2);not null;comment:'Payment amount'"`
	Currency          string        `json:"currency" gorm:"column:currency;type:varchar(10);not null;default:'USD';comment:'Currency code'"`
	Status            PaymentStatus `json:"status" gorm:"column:status;type:ENUM('initiated','succeeded','failed','refunded');default:'initiated';comment:'Payment status'"`
	RawResponse       *RawResponse  `json:"raw_response,omitempty" gorm:"type:json;column:raw_response;comment:'Raw response from payment provider'"`
	CreatedAt         time.Time     `json:"created_at" gorm:"autoCreateTime;column:created_at;comment:'Created at'"`
	UpdatedAt         time.Time     `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;comment:'Updated at'"`
	
	// Relation
	Order             *Order        `json:"-" gorm:"foreignKey:OrderID"`
}

// TableName specifies the table name for the Payment model
func (Payment) TableName() string {
	return "payments"
}

// BeforeCreate sets timestamps and performs any pre-creation logic.
func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UTC()
	if p.CreatedAt.IsZero() {
		p.CreatedAt = now
	}
	p.UpdatedAt = now
	return nil
}

// BeforeUpdate updates the UpdatedAt timestamp before updating an existing record.
func (p *Payment) BeforeUpdate(tx *gorm.DB) (err error) {
	p.UpdatedAt = time.Now().UTC()
	return nil
}

