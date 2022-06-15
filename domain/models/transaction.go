package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TransactionModel struct {
	ID           string `json:"id" gorm:"primary_key"`
	Customer     CustomerModel
	CustomerID   string `json:"customer_id" gorm:"index"`
	Outlet       OutletModel
	OutletID     string    `json:"outlet_id" gorm:"index"`
	Products     pq.StringArray `json:"products" gorm:"type:text[]"`
	PurchaseDate time.Time `json:"purchase_date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (t *TransactionModel) BeforeCreate(db *gorm.DB) error {
	t.ID = uuid.NewString()
	t.CreatedAt = time.Now()
	return nil
}

func (t *TransactionModel) BeforeUpdate(db *gorm.DB) error {
	t.ID = uuid.NewString()
	t.UpdatedAt = time.Now()
	return nil
}
