package clicks

import "github.com/gxjakkap/dekcpe.link/model"

type Store interface {
	Create(click *model.Click) error
}
