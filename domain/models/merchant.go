package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MerchantModel struct {
	ID         string `json:"id" gorm:"primary_key"`
	Name       string `json:"name" gorm:"type:varchar; not null"`
	Phone      string `json:"phone" gorm:"type:bigint; unique; not null"`
	Address    string `json:"address" gorm:"type:text; not null"`
	Logo       string `json:"logo" gorm:"type:varchar; not null"`
	Supplier   SupplierModel
	SupplierID string    `json:"supplier_id" gorm:"index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (m *MerchantModel) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.CreatedAt = time.Now()
	return nil
}

func (m *MerchantModel) BeforeUpdate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.UpdatedAt = time.Now()
	return nil
}
