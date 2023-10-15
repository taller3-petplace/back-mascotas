package services

import (
	"petplace/back-mascotas/cmd/app/data"
	"petplace/back-mascotas/cmd/app/db"
)

type PetPlace struct {
	db db.Storabe
}

func NewPetPlace(db db.Storabe) PetPlace {
	return PetPlace{db: db}
}

func (pp *PetPlace) RegisterNewPet(pet data.Pet) error {

	return pp.db.Save(string(rune(pet.Id)), pet)

}
