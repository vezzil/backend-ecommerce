package entity

import (
	"time"
)

// Cart represents a shopping cart
type Cart struct {
	ID         string     `json:"id" gorm:"primaryKey;column:id;type:varchar(36);default:(UUID());comment:'Primary Key'"`
	UserID     *string    `json:"user_id,omitempty" gorm:"column:user_id;type:varchar(36);comment:'FK to user entity'"`
	GuestToken string     `json:"guest_token,omitempty" gorm:"column:guest_token;type:varchar(255);comment:'Guest session token'"`
	CreatedAt  time.Time  `json:"created_at" gorm:"autoCreateTime;column:created_at;comment:'Created at'"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;comment:'Updated at'"`
	
	// Relations
	Items      []CartItem `json:"items,omitempty" gorm:"foreignKey:CartID"`
}

// TableName specifies the table name for the Cart model
func (Cart) TableName() string {
	return "carts"
}

// CartItem represents an item in a shopping cart
type CartItem struct {
	ID         string    `json:"id" gorm:"primaryKey;column:id;type:varchar(36);default:(UUID());comment:'Primary Key'"`
	CartID     string    `json:"cart_id" gorm:"column:cart_id;type:varchar(36);not null;comment:'FK to cart'"`
	ProductID  string    `json:"product_id" gorm:"column:product_id;type:varchar(36);not null;comment:'FK to product'"`
	Quantity   int       `json:"quantity" gorm:"column:quantity;type:int;not null;default:1;comment:'Item quantity'"`
	PriceAtAdd float64   `json:"price_at_add" gorm:"column:price_at_add;type:decimal(12,2);not null;comment:'Price when added to cart'"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime;column:created_at;comment:'Created at'"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;comment:'Updated at'"`
	
	// Relations
	Cart       *Cart     `json:"-" gorm:"foreignKey:CartID"`
	Product    *Product  `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

// TableName specifies the table name for the CartItem model
func (CartItem) TableName() string {
	return "cartItems"
}
