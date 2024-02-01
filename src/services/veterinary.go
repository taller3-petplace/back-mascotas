package services

import (
	"errors"
	"petplace/back-mascotas/src/db"
	"petplace/back-mascotas/src/db/objects"
	"petplace/back-mascotas/src/model"
)

type VeterinaryService interface {
	New(veterinary model.Veterinary) (model.Veterinary, error)
	Get(id int) (model.Veterinary, error)
	Edit(id int, veterinary model.Veterinary) (model.Veterinary, error)
	GetVeterinaries(filters map[string]string, searchParams *model.SearchParams) (*model.SearchResponse[model.Veterinary], error)
	Delete(id int)
}

type Veterinary struct {
	db db.Storable
}

func NewVeterinaryService(db db.Storable) *Veterinary {
	return &Veterinary{
		db: db,
	}
}

func (v *Veterinary) New(veterinary model.Veterinary) (model.Veterinary, error) {

	var object objects.Veterinary
	object.FromModel(veterinary)
	err := v.db.Save(&object)
	if err != nil {
		return model.Veterinary{}, err
	}
	return object.ToModel(), nil
}

func (v *Veterinary) Get(id int) (model.Veterinary, error) {
	var object objects.Veterinary
	err := v.db.Get(id, &object)
	if err != nil {
		return model.Veterinary{}, err
	}
	return object.ToModel(), nil
}

func (v *Veterinary) Edit(id int, veterinary model.Veterinary) (model.Veterinary, error) {

	veterinary.ID = id

	var object objects.Veterinary
	object.FromModel(veterinary)

	err := v.db.Save(&object)
	if err != nil {
		return model.Veterinary{}, err
	}
	return object.ToModel(), nil
}

func (v *Veterinary) Delete(d int) {
	var object objects.Veterinary
	err := v.db.Delete(d, &object)
	if err != nil {
		return
	}
}

func (v *Veterinary) GetVeterinaries(filters map[string]string, searchParams *model.SearchParams) (*model.SearchResponse[model.Veterinary], error) {

	var objects []objects.Veterinary
	total, err := v.db.GetFiltered(&objects, filters, "Name ASC", int(searchParams.Limit), int(searchParams.Offset))

	if err != nil {
		return nil, errors.New("error fetching from db")
	}

	result := model.SearchResponse[model.Veterinary]{
		Paging: model.Paging{
			Total:  uint(total),
			Offset: searchParams.Offset,
			Limit:  searchParams.Limit,
		},
		Results: []model.Veterinary{},
	}

	for _, object := range objects {
		result.Results = append(result.Results, object.ToModel())
	}

	return &result, nil
}
