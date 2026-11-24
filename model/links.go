package model

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Slug      string    `gorm:"unique;not null"`
	URL       string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `gorm:"not null;default:now()"`
	// Set by the dashboard on delete; gorm adds `deleted_at IS NULL` to every query.
	DeletedAt gorm.DeletedAt

	Clicks []Click `gorm:"foreignKey:LinkID"`
}
