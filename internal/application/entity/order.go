package entity

import (
	"time"

	"gorm.io/gorm"
)

// OrderStatus represents the status of an order
type OrderStatus string

const (
	OrderStatusPending     OrderStatus = "pending"
	OrderStatusPaid        OrderStatus = "paid"
	OrderStatusProcessing  OrderStatus = "processing"
	OrderStatusShipped     OrderStatus = "shipped"
	OrderStatusCancelled   OrderStatus = "cancelled"
	OrderStatusCompleted   OrderStatus = "completed"
)

// Order represents a customer order
type Order struct {
	ID                string         `json:"id" gorm:"primaryKey;column:id;type:varchar(36);default:(UUID());comment:'Primary Key'"`
	UserID            *string        `json:"user_id,omitempty" gorm:"column:user_id;type:varchar(36);comment:'FK to user'"`
	Status            OrderStatus    `json:"status" gorm:"column:status;type:ENUM('pending','paid','processing','shipped','cancelled','completed');default:'pending';comment:'Order status'"`
	TotalAmount       float64        `json:"total_amount" gorm:"column:total_amount;type:decimal(12,2);not null;comment:'Total order amount'"`
	Currency          string         `json:"currency" gorm:"column:currency;type:varchar(10);not null;default:'USD';comment:'Currency code'"`
	ShippingAddressID *string        `json:"shipping_address_id,omitempty" gorm:"column:shipping_address_id;type:varchar(36);comment:'FK to shipping address'"`
	BillingAddressID  *string        `json:"billing_address_id,omitempty" gorm:"column:billing_address_id;type:varchar(36);comment:'FK to billing address'"`
	CreatedAt         time.Time      `json:"created_at" gorm:"autoCreateTime;column:created_at;comment:'Created at'"`
	UpdatedAt         time.Time      `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;comment:'Updated at'"`
	
	// Relations
	User              *User          `json:"user,omitempty" gorm:"foreignKey:UserID"`
	ShippingAddress   *Address       `json:"shipping_address,omitempty" gorm:"foreignKey:ShippingAddressID"`
	BillingAddress    *Address       `json:"billing_address,omitempty" gorm:"foreignKey:BillingAddressID"`
	Items             []OrderItem    `json:"items,omitempty" gorm:"foreignKey:OrderID"`
	Payment           *Payment       `json:"payment,omitempty" gorm:"foreignKey:OrderID"`
}

// TableName specifies the table name for the Order model
func (Order) TableName() string {
	return "orders"
}

// BeforeCreate sets timestamps and performs any pre-creation logic.
func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UTC()
	if o.CreatedAt.IsZero() {
		o.CreatedAt = now
	}
	o.UpdatedAt = now
	return nil
}

// BeforeUpdate updates the UpdatedAt timestamp before updating an existing record.
func (o *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	o.UpdatedAt = time.Now().UTC()
	return nil
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID           string    `json:"id" gorm:"primaryKey;column:id;type:varchar(36);default:(UUID());comment:'Primary Key'"`
	OrderID      string    `json:"order_id" gorm:"column:order_id;type:varchar(36);not null;comment:'FK to order'"`
	ProductID    *string   `json:"product_id,omitempty" gorm:"column:product_id;type:varchar(36);comment:'FK to product'"`
	ProductName  string    `json:"product_name" gorm:"column:product_name;type:varchar(255);not null;comment:'Product name at time of order'"`
	SKU          string    `json:"sku,omitempty" gorm:"column:sku;type:varchar(100);comment:'Product SKU'"`
	Quantity     int       `json:"quantity" gorm:"column:quantity;type:int;not null;comment:'Item quantity'"`
	UnitPrice    float64   `json:"unit_price" gorm:"column:unit_price;type:decimal(12,2);not null;comment:'Price per unit'"`
	TotalPrice   float64   `json:"total_price" gorm:"column:total_price;type:decimal(12,2);not null;comment:'Total price (quantity * unit_price)'"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime;column:created_at;comment:'Created at'"`
	
	// Relations
	Order        *Order     `json:"-" gorm:"foreignKey:OrderID"`
	Product      *Product   `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

// TableName specifies the table name for the OrderItem model
func (OrderItem) TableName() string {
	return "order_items"
}

// BeforeCreate sets timestamps and performs any pre-creation logic.
func (oi *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UTC()
	if oi.CreatedAt.IsZero() {
		oi.CreatedAt = now
	}
	return nil
}
