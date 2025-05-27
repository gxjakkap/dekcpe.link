package model

import (
	"time"
)

type GeoLocation struct {
	Country  string
	Region   string
	City     string
	Timezone string
}

type Click struct {
	ID          int    `gorm:"primaryKey;autoIncrement"`
	LinkID      int    `gorm:"not null"`
	UserAgent   string `gorm:"not null"`
	UTMSource   *string
	Timestamp   time.Time   `gorm:"not null;default:now()"`
	Geolocation GeoLocation `gorm:"type:jsonb"`
}
