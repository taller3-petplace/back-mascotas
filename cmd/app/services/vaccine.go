package services

import (
	"fmt"
	"petplace/back-mascotas/cmd/app/db"
	"petplace/back-mascotas/cmd/app/db/objects"
	"petplace/back-mascotas/cmd/app/model"
)

type VaccineService struct {
	db db.Storable
}

func NewVaccineService(db db.Storable) VaccineService {
	return VaccineService{db: db}
}

func (vs *VaccineService) New(vaccine model.Vaccine) (model.Vaccine, error) {

	err := vs.db.Save(&vaccine)
	if err != nil {
		return model.Vaccine{}, err
	}

	return vaccine, err
}

func (vs *VaccineService) Get(id int) (model.Vaccine, error) {

	var object objects.Vaccine
	err := vs.db.Get(id, object)
	if err != nil {
		return model.Vaccine{}, err
	}

	return object.ToModel(), nil
}

func (vs *VaccineService) Edit(id int, vaccine model.Vaccine) (model.Vaccine, error) {
	var object objects.Vaccine
	object.FromModel(vaccine)
	err := vs.db.Save(&object)
	if err != nil {
		fmt.Println(err)
	}
	return object.ToModel(), nil
}

func (vs *VaccineService) Delete(id int) {
	var object objects.Vaccine
	err := vs.db.Delete(id, &object)
	if err != nil {
		fmt.Println(err)
	}
}
