package store

import (
	"github.com/gxjakkap/dekcpe.link/model"
	"gorm.io/gorm"
)

type ClicksStore struct {
	db *gorm.DB
}

func NewClicksStore(db *gorm.DB) *ClicksStore {
	return &ClicksStore{
		db: db,
	}
}

func (s *ClicksStore) Create(click *model.Click) error {
	if err := s.db.Create(click).Error; err != nil {
		return err
	}
	return nil
}
