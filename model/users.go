package model

import (
	"time"
)

type User struct {
	ID            string `gorm:"primaryKey"`
	Name          string `gorm:"not null"`
	Email         string `gorm:"not null;unique"`
	EmailVerified bool   `gorm:"not null;default:false"`
	Image         *string
	CreatedAt     time.Time `gorm:"not null;default:now()"`
	UpdatedAt     time.Time `gorm:"not null;default:now()"`

	Links []Link `gorm:"foreignKey:Owner;references:ID"`
}
