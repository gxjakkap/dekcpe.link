package link

import (
	"github.com/gxjakkap/dekcpe.link/model"
)

type Store interface {
	GetLinkBySlug(string) (*model.Link, error)
}
