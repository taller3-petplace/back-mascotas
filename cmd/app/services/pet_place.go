package services

import (
	"errors"
	"fmt"
	"petplace/back-mascotas/cmd/app/db"
	"petplace/back-mascotas/cmd/app/db/objects"
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

const tableName = "pets"

type PetPlace struct {
	ABMService[model.Pet]
	db db.Storable
}

func NewPetPlace(db db.Storable) PetPlace {
	return PetPlace{db: db}
}

func (pp *PetPlace) New(pet model.Pet) (model.Pet, error) {

	pet.RegisterDate = time.Now()

	var object objects.Pet
	object.FromModel(pet)
	err := pp.db.Save(&object)
	if err != nil {
		return model.Pet{}, err
	}
	pet.ID = object.ID
	return pet, nil

}

func (pp *PetPlace) Get(petID int) (model.Pet, error) {

	var object objects.Pet
	err := pp.db.Get(petID, &object)
	if err != nil && errors.Is(err, errors.New("not found")) {
		return model.Pet{}, err
	}

	return object.ToModel(), nil
}

func (pp *PetPlace) GetPetsByOwner(request model.SearchRequest) (model.SearchResponse, error) {

	err := pp.db.GetFiltered(func(item interface{}) bool {
		return item.(model.Pet).OwnerID == request.OwnerId
	})

	var pets []model.Pet
	for _, item := range pets {
		pets = append(pets, item)
	}

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
		return int(a.ID) - int(b.ID)
	})

	from := min(result.Paging.Offset, result.Paging.Total)
	to := min(result.Paging.Offset+result.Paging.Limit, result.Paging.Total)
	result.Results = pets[from:to]

	return result, nil
}

func (pp *PetPlace) Edit(petID int, pet model.Pet) (model.Pet, error) {

	var object objects.Pet
	object.FromModel(pet)
	err := pp.db.Save(&object)
	if err != nil {
		fmt.Println(err)
	}
	return object.ToModel(), nil
}

func (pp *PetPlace) Delete(petID int) {
	var object objects.Pet
	err := pp.db.Delete(petID, &object)
	if err != nil {
		fmt.Println(err)
	}
}
