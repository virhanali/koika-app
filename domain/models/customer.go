package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerModel struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"type:varchar; not null"`
	Phone     uint64    `json:"phone" gorm:"type:bigint; unique; not null"`
	Address   string    `json:"address" gorm:"type:text; not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *CustomerModel) BeforeCreate(db *gorm.DB) error {
	c.ID = uuid.NewString()
	c.CreatedAt = time.Now()
	return nil
}

func (c *CustomerModel) BeforeUpdate(db *gorm.DB) error {
	c.ID = uuid.NewString()
	c.UpdatedAt = time.Now()
	return nil
}
