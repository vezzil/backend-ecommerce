package entity

import (
	"time"

	"gorm.io/gorm"
)

// Address represents a user's shipping/billing address
type Address struct {
	ID         string    `json:"id" gorm:"primaryKey;column:id;type:varchar(36);default:(UUID());comment:'Primary Key'"`
	UserID     string    `json:"user_id" gorm:"column:user_id;type:varchar(36);not null;comment:'FK to user entity'"`
	Label      string    `json:"label" gorm:"column:label;type:varchar(100);comment:'Address label'"`
	Street     string    `json:"street" gorm:"column:street;type:varchar(255);comment:'Street address'"`
	City       string    `json:"city" gorm:"column:city;type:varchar(100);comment:'City'"`
	State      string    `json:"state" gorm:"column:state;type:varchar(100);comment:'State/Province'"`
	PostalCode string    `json:"postal_code" gorm:"column:postal_code;type:varchar(50);comment:'Postal/ZIP code'"`
	Country    string    `json:"country" gorm:"column:country;type:varchar(100);comment:'Country'"`
	Phone      string    `json:"phone" gorm:"column:phone;type:varchar(50);comment:'Contact phone number'"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime;column:created_at;comment:'Created at'"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;comment:'Updated at'"`
	User       User      `json:"-" gorm:"foreignKey:UserID"`
}

// TableName specifies the table name for the Address model
func (Address) TableName() string {
	return "addresses"
}

// BeforeCreate sets timestamps and performs any pre-creation logic.
func (u *Address) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = now
	}
	u.UpdatedAt = now
	return nil
}

// BeforeUpdate updates the UpdatedAt timestamp before updating an existing record.
func (u *Address) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now().UTC()
	return nil
}