package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type RoleAllowed string

const (
	admin    RoleAllowed = "admin"
	merchant RoleAllowed = "merchant"
	customer RoleAllowed = "customer"
	supplier RoleAllowed = "supplier"
)

type RoleModel struct {
	ID         string         `json:"id" gorm:"primary_key"`
	RoleName   string         `json:"role_name" sql:"type:role_name"`
	RoleAccess pq.StringArray `json:"role_access" gorm:"type:text[]; not null"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

func (r *RoleModel) BeforeCreate(db *gorm.DB) error {
	r.ID = uuid.New().String()
	r.CreatedAt = time.Now()
	return nil
}

func (r *RoleModel) BeforeUpdate(db *gorm.DB) error {
	r.ID = uuid.New().String()
	r.UpdatedAt = time.Now()
	return nil
}
