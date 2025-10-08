package entity

import (
	"time"

	"gorm.io/gorm"
)

// Product represents an item for sale
type Product struct {
	ID           string         `json:"id" gorm:"primaryKey;column:id;type:varchar(36);default:(UUID());comment:'Primary Key'"`
	SKU          string         `json:"sku,omitempty" gorm:"column:sku;type:varchar(100);uniqueIndex;comment:'Stock Keeping Unit'"`
	Name         string         `json:"name" gorm:"column:name;type:varchar(255);not null;comment:'Product name'"`
	Slug         string         `json:"slug" gorm:"column:slug;type:varchar(255);uniqueIndex;not null;comment:'URL-friendly name'"`
	Description  string         `json:"description,omitempty" gorm:"column:description;type:text;comment:'Product description'"`
	Price        float64        `json:"price" gorm:"column:price;type:decimal(12,2);not null;comment:'Product price'"`
	Currency     string         `json:"currency" gorm:"column:currency;type:varchar(10);not null;default:'USD';comment:'Currency code'"`
	CategoryID   *string        `json:"category_id,omitempty" gorm:"column:category_id;type:varchar(36);comment:'FK to category'"`
	IsActive     bool           `json:"is_active" gorm:"column:is_active;type:boolean;default:true;comment:'Is product active'"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime;column:created_at;comment:'Created at'"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;comment:'Updated at'"`
	
	// Relations
	// Category     *Category      `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	Images       []ProductImage `json:"images,omitempty" gorm:"foreignKey:ProductID"`
	// Inventory    *Inventory     `json:"inventory,omitempty" gorm:"foreignKey:ProductID"`
}

// TableName specifies the table name for the Product model
func (Product) TableName() string {
	return "products"
}

// BeforeCreate sets timestamps and performs any pre-creation logic.
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UTC()
	if p.CreatedAt.IsZero() {
		p.CreatedAt = now
	}
	p.UpdatedAt = now
	return nil
}

// BeforeUpdate updates the UpdatedAt timestamp before updating an existing record.
func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	p.UpdatedAt = time.Now().UTC()
	return nil
}


