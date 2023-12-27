package services

import (
	"encoding/json"
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
	item, err := pp.db.Save(tableName, pet)
	pet.ID = item.ID

	newItem := db.StorableItem{
		ID:   item.ID,
		Data: pet,
	}

	// Update de ID
	_, err = pp.db.Save(tableName, newItem)
	if err != nil {
		return model.Pet{}, err
	}

	return pet, err

}

func (pp *PetPlace) Get(petID int) (model.Pet, error) {
	item, err := pp.db.Get(tableName, petID)
	if err != nil && errors.Is(err, errors.New("not found")) {
		return model.Pet{}, err
	}
	if item == nil {
		return model.Pet{}, nil
	}

	return getPetFromItem(*item), nil
}

func (pp *PetPlace) GetPetsByOwner(request model.SearchRequest) (model.SearchResponse, error) {

	items, err := pp.db.GetFiltered(tableName, func(item db.StorableItem) bool {
		return getPetFromItem(item).OwnerID == request.OwnerId
	})

	var pets []model.Pet
	for _, item := range items {
		pets = append(pets, getPetFromItem(item))
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
		return a.ID - b.ID
	})

	from := min(result.Paging.Offset, result.Paging.Total)
	to := min(result.Paging.Offset+result.Paging.Limit, result.Paging.Total)
	result.Results = pets[from:to]

	return result, nil
}

func (pp *PetPlace) Edit(petID int, pet model.Pet) (model.Pet, error) {
	pet.ID = petID
	item := db.StorableItem{
		ID:   petID,
		Data: pet,
	}
	_, err := pp.db.Save(tableName, item)
	return pet, err
}

func (pp *PetPlace) Delete(petID int) {
	pp.db.Delete(tableName, petID)
}

func getPetFromItem(item db.StorableItem) model.Pet {

	switch v := item.Data.(type) {
	case model.Pet:
		return v
	}

	var result model.Pet
	d, _ := json.Marshal(item.Data)
	err := json.Unmarshal(d, &result)
	if err != nil {
		return model.Pet{}
	}

	return result
}
