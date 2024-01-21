package services

import (
	"errors"
	"fmt"
	"petplace/back-mascotas/src/db"
	"petplace/back-mascotas/src/db/objects"
	"petplace/back-mascotas/src/model"
	"time"
)

type VaccineService struct {
	db db.Storable
}

func NewVaccineService(db db.Storable) VaccineService {
	return VaccineService{db: db}
}

func (vs *VaccineService) New(vaccine model.Vaccine) (model.Vaccine, error) {

	var object objects.Vaccine
	object.FromModel(vaccine)
	err := vs.db.Save(&object)
	if err != nil {
		return model.Vaccine{}, err
	}

	return object.ToModel(), err
}

func (vs *VaccineService) Get(id int) (model.Vaccine, error) {

	var object objects.Vaccine
	err := vs.db.Get(id, &object)
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

func (vs *VaccineService) ApplyVaccine(petID uint, vaccineID uint) error {

	var object objects.Application
	object.PetID = petID
	object.VaccineID = vaccineID
	object.AppliedAt = time.Now()
	err := vs.db.Save(&object)

	return err
}

func (vs *VaccineService) GetPlanVaccination(petID int) (model.VaccinationPlan, error) {

	// pegarle a la tabla de mascotas para obtener el tipo
	var pet objects.Pet
	err := vs.db.Get(petID, &pet)
	if err != nil && errors.Is(err, errors.New("not found")) {
		return model.VaccinationPlan{}, err
	}

	if pet.ID == 0 || pet.OwnerID == 0 {
		return model.VaccinationPlan{}, nil
	}

	// pegarle a la tabla de aplicaciones con el petID
	var applications []objects.Application
	_, err = vs.db.GetFiltered(&applications, map[string]string{
		"pet_id": fmt.Sprintf("%d", petID),
	}, "applied_at desc", 100, 0)
	if err != nil {
		return model.VaccinationPlan{}, err
	}

	// pegarle a la tabla de vaccines para evaluar cuales tengo y cuales me faltan
	var vaccines []objects.Vaccine
	_, err = vs.db.GetFiltered(&vaccines, map[string]string{
		"animal": pet.Type,
	}, "scheduled", 100, 0)

	return getVaccinationPlan(pet, vaccines, applications), nil
}

func getVaccinationPlan(pet objects.Pet, vaccines []objects.Vaccine, applications []objects.Application) model.VaccinationPlan {

	var result model.VaccinationPlan
	result.Name = pet.Name
	result.Type = pet.Type
	result.OwnerID = pet.OwnerID

	for _, app := range applications {
		tmp := getVaccine(app.VaccineID, vaccines)
		tmp.AppliedAt = &app.AppliedAt
		result.Applied = append(result.Applied, tmp)
	}

	for _, v := range vaccines {
		if !applied(v, applications) {
			result.Pending = append(result.Pending, v.ToModel())
		}
	}

	return result
}

func applied(vaccine objects.Vaccine, apps []objects.Application) bool {
	for _, app := range apps {
		if app.VaccineID == vaccine.ID {
			return true
		}
	}
	return false
}

func getVaccine(id uint, vs []objects.Vaccine) model.Vaccine {
	for _, v := range vs {
		if v.ID == id {
			return v.ToModel()
		}
	}
	return model.Vaccine{
		ID:          id,
		Animal:      "unknown",
		Name:        "unknown",
		Description: "unknown",
		Scheduled:   0,
	}
}
