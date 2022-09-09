package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
}
