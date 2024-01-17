package services

import (
	"errors"
	"fmt"
	"petplace/back-mascotas/db"
	"petplace/back-mascotas/db/objects"
	model2 "petplace/back-mascotas/model"
	"strconv"
	"time"
)

type PetService interface {
	New(pet model2.Pet) (model2.Pet, error)
	Get(petID int) (model2.Pet, error)
	Edit(petID int, pet model2.Pet) (model2.Pet, error)
	Delete(petID int)
	GetPetsByOwner(request model2.SearchRequest) (model2.SearchResponse, error)
}

const tableName = "pets"

type PetPlace struct {
	ABMService[model2.Pet]
	db db.Storable
}

func NewPetPlace(db db.Storable) PetPlace {
	return PetPlace{db: db}
}

func (pp *PetPlace) New(pet model2.Pet) (model2.Pet, error) {

	pet.RegisterDate = time.Now()

	var object objects.Pet
	object.FromModel(pet)
	err := pp.db.Save(&object)
	if err != nil {
		return model2.Pet{}, err
	}
	pet.ID = int(object.ID)
	return pet, nil

}

func (pp *PetPlace) Get(petID int) (model2.Pet, error) {

	var object objects.Pet
	err := pp.db.Get(petID, &object)
	if err != nil && errors.Is(err, errors.New("not found")) {
		return model2.Pet{}, err
	}

	return object.ToModel(), nil
}

func (pp *PetPlace) GetPetsByOwner(request model2.SearchRequest) (model2.SearchResponse, error) {

	var objects []objects.Pet
	total, err := pp.db.GetFiltered(&objects, map[string]string{
		"owner_id": strconv.Itoa(request.OwnerId),
	}, "Name ASC", int(request.Limit), int(request.Offset))

	if err != nil {
		return model2.SearchResponse{}, errors.New("error fetching from db")
	}

	result := model2.SearchResponse{
		Paging: model2.Paging{
			Total:  uint(total),
			Offset: request.Offset,
			Limit:  request.Limit,
		},
		Results: []model2.Pet{},
	}

	for _, object := range objects {
		result.Results = append(result.Results, object.ToModel())
	}

	return result, nil
}

func (pp *PetPlace) Edit(petID int, pet model2.Pet) (model2.Pet, error) {

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
