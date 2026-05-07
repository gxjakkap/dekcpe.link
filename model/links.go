package model

import (
	"time"
)

type Link struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Slug      string    `gorm:"unique;not null"`
	URL       string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `gorm:"not null;default:now()"`
	Owner     string    `gorm:"not null"`

	User   User    `gorm:"foreignKey:Owner;references:ID;constraint:OnDelete:CASCADE"`
	Clicks []Click `gorm:"foreignKey:LinkID"`
}
