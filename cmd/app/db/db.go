package db

import "petplace/back-mascotas/cmd/app/data"

type Storabe interface {
	Save(id string, pet data.Pet) error
	Get(id string) (data.Pet, error)
	Delete(id string) error
}
