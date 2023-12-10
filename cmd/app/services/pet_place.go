package services

import (
	"math/rand"
	"petplace/back-mascotas/cmd/app/data"
	"petplace/back-mascotas/cmd/app/db"
	"strconv"
	"time"
)

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
