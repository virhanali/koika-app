package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/virhanali/koika-app/config/pkg"
	"gorm.io/gorm"
)

type UserModel struct {
	ID        string    `json:"id" gorm:"primary_key"`
	FirstName string    `json:"first_name" gorm:"type:varchar; not null"`
	LastName  string    `json:"last_name" gorm:"type:varchar; not null"`
	Email     string    `json:"email" gorm:"type:varchar; not null"`
	Password  string    `json:"password" gorm:"type:varchar; not null"`
	Role      string    `json:"role" gorm:"type:varchar; not null"`
	Active    bool      `json:"active" gorm:"type:boolean; not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *UserModel) BeforeCreate(db *gorm.DB) error {
	u.ID = uuid.NewString()
	u.Password = pkg.HashPassword(u.Password)
	u.CreatedAt = time.Now()
	return nil
}

func (u *UserModel) BeforeUpdate(db *gorm.DB) error {
	u.ID = uuid.NewString()
	u.Password = pkg.HashPassword(u.Password)
	u.UpdatedAt = time.Now()
	return nil
}
