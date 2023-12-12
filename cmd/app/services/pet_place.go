package services

import (
	"errors"
	"petplace/back-mascotas/cmd/app/data"
	"petplace/back-mascotas/cmd/app/db"
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

	pet.RegisterDate = time.Now()
	err := pp.db.Save(&pet)

	return pet, err

}

func (pp *PetPlace) GetPet(petID int) (data.Pet, error) {
	getPet, err := pp.db.Get(petID)
	if err != nil {
		return data.Pet{}, err
	}
	return getPet, nil
}

func (pp *PetPlace) GetPetsByOwner(request data.SearchRequest) (data.SearchResponse, error) {

	pets, err := pp.db.GetByOwner(request.OwnerId)
	if err != nil {
		return data.SearchResponse{}, errors.New("error fetching from db")
	}

	result := data.SearchResponse{
		Paging: data.Paging{
			Total:  uint(len(pets)),
			Offset: request.Offset,
			Limit:  request.Limit,
		},
		Results: []data.Pet{},
	}

	from := min(result.Paging.Offset, result.Paging.Total)
	to := min(result.Paging.Offset+result.Paging.Limit, result.Paging.Total)
	result.Results = pets[from:to]

	return result, nil
}
