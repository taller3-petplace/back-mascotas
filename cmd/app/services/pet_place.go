package services

import (
	"errors"
	"petplace/back-mascotas/cmd/app/db"
	"petplace/back-mascotas/cmd/app/model"
	"slices"
	"time"
)

type PetService interface {
	New(pet model.Pet) (model.Pet, error)
	Get(petID int) (model.Pet, error)
	Edit(petID int, pet model.Pet) (model.Pet, error)
	Delete(petID int)
	GetPetsByOwner(request model.SearchRequest) (model.SearchResponse, error)
}

type PetPlace struct {
	ABMService[model.Pet]
	db db.Storable
}

func NewPetPlace(db db.Storable) PetPlace {
	return PetPlace{db: db}
}

func (pp *PetPlace) New(pet model.Pet) (model.Pet, error) {

	pet.RegisterDate = time.Now()
	err := pp.db.Save(&pet)

	return pet, err

}

func (pp *PetPlace) Get(petID int) (model.Pet, error) {
	getPet, err := pp.db.Get(petID)
	if err != nil && errors.Is(err, errors.New("not found")) {
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

	slices.SortFunc(pets, func(a, b model.Pet) int {
		return a.ID - b.ID
	})

	from := min(result.Paging.Offset, result.Paging.Total)
	to := min(result.Paging.Offset+result.Paging.Limit, result.Paging.Total)
	result.Results = pets[from:to]

	return result, nil
}

func (pp *PetPlace) Edit(id int, pet model.Pet) (model.Pet, error) {
	err := pp.db.Save(&pet)
	return pet, err
}

func (pp *PetPlace) Delete(petID int) {
	pp.db.Delete(petID)
}
