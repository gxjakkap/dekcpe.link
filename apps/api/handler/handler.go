package handler

import (
	"github.com/gxjakkap/dekcpe.link/clicks"
	"github.com/gxjakkap/dekcpe.link/link"
)

type Handler struct {
	linkStore   link.Store
	clicksStore clicks.Store
	validator   *Validator
}

func NewHandler(linkStore link.Store, clicksStore clicks.Store) *Handler {
	v := NewValidator()
	return &Handler{
		linkStore:   linkStore,
		clicksStore: clicksStore,
		validator:   v,
	}
}
