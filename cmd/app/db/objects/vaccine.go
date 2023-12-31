package objects

import (
	"petplace/back-mascotas/cmd/app/model"
)

type Vaccine struct {
	ID          int `gorm:"primaryKey;autoIncrement;unique"`
	Animal      string
	Name        string
	Description string
	Scheduled   uint
}

func (v *Vaccine) ToModel() model.Vaccine {
	return model.Vaccine{
		ID:          v.ID,
		Animal:      model.AnimalType(v.Animal),
		Name:        v.Name,
		Description: v.Description,
		Scheduled:   v.Scheduled,
	}
}

func (v *Vaccine) FromModel(vaccine model.Vaccine) {
	v.ID = vaccine.ID
	v.Animal = string(vaccine.Animal)
	v.Name = vaccine.Name
	v.Description = vaccine.Description
	v.Scheduled = vaccine.Scheduled

}
