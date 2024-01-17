package objects

import (
	model2 "petplace/back-mascotas/model"
)

type Vaccine struct {
	ID          int `gorm:"primaryKey;autoIncrement;unique"`
	Animal      string
	Name        string
	Description string
	Scheduled   uint
}

func (v *Vaccine) ToModel() model2.Vaccine {
	return model2.Vaccine{
		ID:          v.ID,
		Animal:      model2.AnimalType(v.Animal),
		Name:        v.Name,
		Description: v.Description,
		Scheduled:   v.Scheduled,
	}
}

func (v *Vaccine) FromModel(vaccine model2.Vaccine) {
	v.ID = vaccine.ID
	v.Animal = string(vaccine.Animal)
	v.Name = vaccine.Name
	v.Description = vaccine.Description
	v.Scheduled = vaccine.Scheduled

}
