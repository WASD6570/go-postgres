package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Email         string    `gorm:"unique_index,not null" json:"email"`
	Password      string    `json:"password"`
	UserID        uuid.UUID `json:"user_id" gorm:"type:uuid;not null,unique_index,foreignkey:UserID"`
	LoginAttempts int       `json:"login_attempts"`
}
