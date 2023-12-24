package db

import "petplace/back-mascotas/cmd/app/model"

type Storable interface {
	Save(pet *model.Pet) error
	Get(id int) (model.Pet, error)
	Delete(id int)
	GetByOwner(OwnerID int) ([]model.Pet, error)
}
