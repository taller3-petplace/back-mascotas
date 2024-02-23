package objects

import (
	"encoding/json"
	"gorm.io/gorm"
	"petplace/back-mascotas/src/model"
	"time"
)

type Veterinary struct {
	ID        uint `gorm:"primaryKey;autoIncrement;unique"`
	Name      string
	Address   string
	Phone     string
	Email     string
	WebSite   string
	IMGUrl    string
	City      string
	DayGuard  int
	Doctors   Doctors        `gorm:"type:json"`
	CreatedAt time.Time      `gorm:"type:timestamp"`
	UpdatedAt time.Time      `gorm:"type:timestamptz"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamptz"`
}

type Doctors string

type Doctor struct {
	ID           uint
	Name         string
	Phone        string
	Email        string
	VeterinaryID uint
}

func (ds *Doctors) FromModel(doctors []model.Doctor) {

	result, _ := json.Marshal(doctors)
	*ds = Doctors(result)
}

func (d *Doctor) FromModel(doctor model.Doctor) {

	d.ID = uint(doctor.ID)
	d.Name = doctor.Name
	d.Email = doctor.Email
	d.Phone = doctor.Phone
}

func (v *Veterinary) FromModel(vet model.Veterinary) {
	v.ID = uint(vet.ID)
	v.Name = vet.Name
	v.Address = vet.Address
	v.Phone = vet.Phone
	v.Email = vet.Email
	v.WebSite = vet.WebSite
	v.IMGUrl = vet.IMGUrl
	v.City = vet.City
	v.DayGuard = vet.DayGuard
	v.Doctors.FromModel(vet.Doctors)
}

func (v *Veterinary) ToModel() model.Veterinary {

	var result model.Veterinary
	result.ID = int(v.ID)
	result.Name = v.Name
	result.Address = v.Address
	result.Phone = v.Phone
	result.Email = v.Email
	result.WebSite = v.WebSite
	result.IMGUrl = v.IMGUrl
	result.City = v.City
	result.DayGuard = v.DayGuard
	result.Doctors = v.Doctors.ToModel()
	return result
}

func (ds *Doctors) ToModel() []model.Doctor {
	var result []model.Doctor
	_ = json.Unmarshal([]byte(*ds), &result)
	return result
}

func (d *Doctor) ToModel() model.Doctor {
	var result model.Doctor
	result.ID = int(d.ID)
	result.Name = d.Name
	result.Phone = d.Phone
	result.Email = d.Email
	return result
}
