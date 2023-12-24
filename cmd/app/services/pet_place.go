package services

import (
	"errors"
	"petplace/back-mascotas/cmd/app/db"
	"petplace/back-mascotas/cmd/app/model"
	"time"
)

type PetService interface {
	RegisterNewPet(pet model.Pet) (model.Pet, error)
	GetPet(petID int) (model.Pet, error)
	GetPetsByOwner(request model.SearchRequest) (model.SearchResponse, error)
	EditPet(pet model.Pet) (model.Pet, error)
	DeletePet(petID int)
}

type PetPlace struct {
	db db.Storable
}

func NewPetPlace(db db.Storable) PetPlace {
	return PetPlace{db: db}
}

func (pp *PetPlace) RegisterNewPet(pet model.Pet) (model.Pet, error) {

	pet.RegisterDate = time.Now()
	err := pp.db.Save(&pet)

	return pet, err

}

func (pp *PetPlace) GetPet(petID int) (model.Pet, error) {
	getPet, err := pp.db.Get(petID)
	if err != nil {
		return model.Pet{}, err
	}
	return getPet, nil
}

func (pp *PetPlace) GetPetsByOwner(request model.SearchRequest) (model.SearchResponse, error) {

	pets, err := pp.db.GetByOwner(request.OwnerId)
	if err != nil {
		return model.SearchResponse{}, errors.New("error fetching from db")
	}

	result := model.SearchResponse{
		Paging: model.Paging{
			Total:  uint(len(pets)),
			Offset: request.Offset,
			Limit:  request.Limit,
		},
		Results: []model.Pet{},
	}

	from := min(result.Paging.Offset, result.Paging.Total)
	to := min(result.Paging.Offset+result.Paging.Limit, result.Paging.Total)
	result.Results = pets[from:to]

	return result, nil
}

func (pp *PetPlace) EditPet(pet model.Pet) (model.Pet, error) {
	err := pp.db.Save(&pet)
	return pet, err
}

func (pp *PetPlace) DeletePet(petID int) {
	pp.db.Delete(petID)
}
