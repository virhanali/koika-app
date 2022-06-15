package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OutletModel struct {
	ID         string `json:"id" gorm:"primary_key"`
	Name       string `json:"name" gorm:"type:varchar; not null"`
	Phone      uint64 `json:"phone" gorm:"type:bigint; unique; not null"`
	Address    string `json:"address" gorm:"type:text; not null"`
	Merchant   MerchantModel
	MerchantID string    `json:"merchant_id" gorm:"index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (u *OutletModel) BeforeCreate(db *gorm.DB) error {
	u.ID = uuid.NewString()
	u.CreatedAt = time.Now()
	return nil
}

func (u *OutletModel) BeforeUpdate(db *gorm.DB) error {
	u.ID = uuid.NewString()
	u.UpdatedAt = time.Now()
	return nil
}
