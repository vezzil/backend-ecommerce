package entity

import (
	"time"

	"gorm.io/gorm"
)

// Category represents a product category
type Category struct {
	ID          string    `json:"id" gorm:"primaryKey;column:id;type:varchar(36);default:(UUID());comment:'Primary Key'"`
	Name        string    `json:"name" gorm:"column:name;type:varchar(150);uniqueIndex;not null;comment:'Category name'"`
	Slug        string    `json:"slug" gorm:"column:slug;type:varchar(150);uniqueIndex;not null;comment:'URL-friendly name'"`
	Description string    `json:"description,omitempty" gorm:"column:description;type:text;comment:'Category description'"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime;column:created_at;comment:'Created at'"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;comment:'Updated at'"`
}

// TableName specifies the table name for the Category model
func (Category) TableName() string {
	return "categories"
}

// BeforeCreate sets timestamps and performs any pre-creation logic.
func (u *Category) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = now
	}
	u.UpdatedAt = now
	return nil
}

// BeforeUpdate updates the UpdatedAt timestamp before updating an existing record.
func (u *Category) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now().UTC()
	return nil
}