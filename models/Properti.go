package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Properti struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	foto       string
	judul      string `gorm:"not null"`
	alamat     string `gorm:"not null"`
	harga      int64  `gorm:"not null"`
	Password   string `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	PenggunaID uuid.UUID
	User       Pengguna `gorm:"foreignKey:UserID"`
}
