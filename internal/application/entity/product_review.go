package entity

import (
	"time"

	"gorm.io/gorm"
)

// ProductReview represents a customer's review of a product
type ProductReview struct {
	ID         string    `gorm:"type:char(36);primaryKey;default:(UUID())" json:"id"`
	ProductID  string    `gorm:"type:char(36);column:productId;not null" json:"productId"`
	UserID     *string   `gorm:"type:char(36);column:userId" json:"userId,omitempty"`
	Rating     int8      `gorm:"not null" json:"rating"`
	Title      string    `gorm:"size:255" json:"title,omitempty"`
	Body       string    `gorm:"type:text" json:"body,omitempty"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	
	// Relations
	Product    *Product  `gorm:"foreignKey:ProductID" json:"-"`
	User       *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// TableName specifies the table name for the ProductReview model
func (ProductReview) TableName() string {
	return "productReviews"
}

// BeforeCreate sets timestamps and performs any pre-creation logic.
func (u *ProductReview) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = now
	}
	u.UpdatedAt = now
	return nil
}

// BeforeUpdate updates the UpdatedAt timestamp before updating an existing record.
func (u *ProductReview) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now().UTC()
	return nil
}

