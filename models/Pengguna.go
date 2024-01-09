package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pengguna struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Username string    `gorm:"unique;not null"`
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Properti []Properti
	// Add other fields as needed
}
