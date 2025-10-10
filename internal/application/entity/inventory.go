package entity

import (
	"time"
	"gorm.io/gorm"
)

// Inventory tracks product stock levels
type Inventory struct {
	ID            string    `json:"id" gorm:"primaryKey;column:id;type:varchar(36);default:(UUID());comment:'Primary Key'"`
	ProductID     string    `json:"product_id" gorm:"column:product_id;type:varchar(36);uniqueIndex;not null;comment:'FK to product'"`
	Quantity      int       `json:"quantity" gorm:"column:quantity;type:int;not null;default:0;comment:'Available quantity'"`
	Reserved      int       `json:"reserved" gorm:"column:reserved;type:int;not null;default:0;comment:'Reserved quantity'"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime;column:created_at;comment:'Created at'"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;comment:'Updated at'"`
	
	// Available calculates available stock (quantity - reserved)
	// Available     int       `json:"available" gorm:"-"`
}

// TableName specifies the table name for the Inventory model
func (Inventory) TableName() string {
	return "inventory"
}

// CalculateAvailable updates the Available field
func (i *Inventory) CalculateAvailable() {
	i.Available = i.Quantity - i.Reserved
}

// BeforeCreate sets timestamps and performs any pre-creation logic.
func (i *Inventory) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UTC()
	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}
	i.UpdatedAt = now
	return nil
}

// BeforeUpdate updates the UpdatedAt timestamp before updating an existing record.
func (u *Inventory) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now().UTC()
	return nil
}