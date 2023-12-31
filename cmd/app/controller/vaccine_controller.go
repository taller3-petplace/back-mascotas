package controller

import (
	"petplace/back-mascotas/cmd/app/model"
	"petplace/back-mascotas/cmd/app/services"
)

type VaccineController struct {
	ABMController[model.Vaccine]
	service services.VaccineService
}

func NewVaccineController(service services.VaccineService) VaccineController {

	temp := VaccineController{}
	temp.service = service
	temp.s = &service
	temp.Validate = ValidateVaccine

	return temp
}

func ValidateVaccine(v model.Vaccine) error {

	if !model.ValidAnimalType(v.Animal) {
		return InvalidAnimalType
	}
	return nil
}
