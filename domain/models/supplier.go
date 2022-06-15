package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SupplierModel struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"type:varchar; not null"`
	Phone     uint64    `json:"phone" gorm:"type:bigint; unique; not null"`
	Address   string    `json:"address" gorm:"type:text; not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *SupplierModel) BeforeCreate(db *gorm.DB) error {
	s.ID = uuid.NewString()
	s.CreatedAt = time.Now()
	return nil
}

func (s *SupplierModel) BeforeUpdate(db *gorm.DB) error {
	s.ID = uuid.NewString()
	s.UpdatedAt = time.Now()
	return nil
}
