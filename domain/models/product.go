package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductModel struct {
	ID         string `json:"id" gorm:"primary_key"`
	Name       string `json:"name" gorm:"type:varchar; not null"`
	Image      string `json:"image" gorm:"type:varchar; not null"`
	SKU        uint64 `json:"sku" gorm:"type:bigint; not null; default=0"`
	Price      uint64 `json:"price" gorm:"type:bigint; not null; default=0"`
	Outlet     OutletModel
	OutletID   string `json:"outlet_id" gorm:"index"`
	Supplier   SupplierModel
	SupplierID string    `json:"supplier_id" gorm:"index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (p *ProductModel) BeforeCreate(db *gorm.DB) error {
	p.ID = uuid.NewString()
	p.CreatedAt = time.Now()
	return nil
}

func (p *ProductModel) BeforeUpdate(db *gorm.DB) error {
	p.ID = uuid.NewString()
	p.UpdatedAt = time.Now()
	return nil
}
