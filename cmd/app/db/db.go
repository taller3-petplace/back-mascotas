package db

import "petplace/back-mascotas/cmd/app/data"

type Storabe interface {
	Save(pet *data.Pet) error
	Get(id int) (data.Pet, error)
	Delete(id int)
	GetByOwner(OwnerID int) ([]data.Pet, error)
}
