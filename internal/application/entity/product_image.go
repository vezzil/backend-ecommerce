package entity

import (
	"time"

	"gorm.io/gorm"
)

// ProductImage represents an image for a product
type ProductImage struct {
	ID        string    `json:"id" gorm:"primaryKey;column:id;type:varchar(36);default:(UUID());comment:'Primary Key'"`
	ProductID string    `json:"product_id" gorm:"column:product_id;type:varchar(36);not null;comment:'FK to product'"`
	URL       string    `json:"url" gorm:"column:url;type:varchar(1000);not null;comment:'Image URL'"`
	AltText   string    `json:"alt_text,omitempty" gorm:"column:alt_text;type:varchar(255);comment:'Alternative text'"`
	SortOrder int       `json:"sort_order" gorm:"column:sort_order;type:int;default:0;comment:'Sort order'"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;column:created_at;comment:'Created at'"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;comment:'Updated at'"`
	
	// Relation
	Product    *Product    `json:"-" gorm:"foreignKey:ProductID"`
}

// TableName specifies the table name for the ProductImage model
func (ProductImage) TableName() string {
	return "productImages"
}

// BeforeCreate sets timestamps and performs any pre-creation logic.
func (u *ProductImage) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = now
	}
	u.UpdatedAt = now
	return nil
}

// BeforeUpdate updates the UpdatedAt timestamp before updating an existing record.
func (u *ProductImage) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now().UTC()
	return nil
}