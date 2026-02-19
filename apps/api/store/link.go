package store

import (
	"errors"

	"github.com/gxjakkap/dekcpe.link/model"
	"gorm.io/gorm"
)

type LinkStore struct {
	db *gorm.DB
}

func NewLinkStore(db *gorm.DB) *LinkStore {
	return &LinkStore{
		db: db,
	}
}

func (s *LinkStore) GetLinkBySlug(sl string) (*model.Link, error) {
	var l model.Link
	if err := s.db.Where(&model.Link{Slug: sl}).First(&l).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &l, nil
}
