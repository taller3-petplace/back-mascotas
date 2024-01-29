package objects

import (
	"gorm.io/gorm"
	"petplace/back-mascotas/src/model"
	"time"
)

type Pet struct {
	ID        uint `gorm:"primaryKey;autoIncrement;unique"`
	Name      string
	Type      string
	CreatedAt time.Time      `gorm:"type:timestamp"`
	UpdatedAt time.Time      `gorm:"type:timestamptz"`
	BirthDate time.Time      `gorm:"type:date"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamptz"`
	OwnerID   string
	IMGUrl    string
}

func (p *Pet) FromModel(pet model.Pet) {
	p.ID = uint(pet.ID)
	p.Name = pet.Name
	p.Type = string(pet.Type)
	p.CreatedAt = pet.RegisterDate
	p.BirthDate = pet.BirthDate.Time
	p.OwnerID = pet.OwnerID
	p.IMGUrl = pet.IMGUrl
}

func (p *Pet) ToModel() model.Pet {

	var pet model.Pet
	pet.ID = int(p.ID)
	pet.Name = p.Name
	pet.Type = model.AnimalType(p.Type)
	pet.RegisterDate = p.CreatedAt
	pet.BirthDate = model.Date{Time: p.BirthDate}
	pet.OwnerID = p.OwnerID
	pet.IMGUrl = p.IMGUrl
	return pet
}
