package services

import (
	"math/rand"
	"petplace/back-mascotas/cmd/app/data"
	"petplace/back-mascotas/cmd/app/db"
	"strconv"
	"time"
)

type PetService interface {
	RegisterNewPet(pet data.Pet) (data.Pet, error)
	GetPet(pet int) (data.Pet, error)
	GetPetsByOwner(request data.SearchRequest) (data.SearchResponse, error)
}

type PetPlace struct {
	db db.Storabe
}

func NewPetPlace(db db.Storabe) PetPlace {
	return PetPlace{db: db}
}

func (pp *PetPlace) RegisterNewPet(pet data.Pet) (data.Pet, error) {

	pet.ID = rand.Int()
	pet.RegisterDate = time.Now()

	err := pp.db.Save(strconv.Itoa(pet.ID), pet)

	return pet, err

}

func (pp *PetPlace) GetPet(pet string) (data.Pet, error) {
	return data.Pet{}, nil
}

func (pp *PetPlace) GetPetsByOwner(id int) ([]data.Pet, error) {
	return nil, nil
}
