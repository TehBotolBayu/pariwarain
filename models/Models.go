package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setup() (*gorm.DB, error) {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	// Auto-migrate the schemas
	if err := db.AutoMigrate(&Pengguna{}, &Properti{}); err != nil {
		return nil, err
	}

	return db, nil
}
