package objects

import (
	"gorm.io/gorm"
	model2 "petplace/back-mascotas/model"
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
	OwnerID   int
}

func (p *Pet) FromModel(pet model2.Pet) {
	p.ID = uint(pet.ID)
	p.Name = pet.Name
	p.Type = string(pet.Type)
	p.CreatedAt = pet.RegisterDate
	p.BirthDate = pet.BirthDate.Time
	p.OwnerID = pet.OwnerID
}

func (p *Pet) ToModel() model2.Pet {

	var pet model2.Pet
	pet.ID = int(p.ID)
	pet.Name = p.Name
	pet.Type = model2.AnimalType(p.Type)
	pet.RegisterDate = p.CreatedAt
	pet.BirthDate = model2.Date{Time: p.BirthDate}
	pet.OwnerID = p.OwnerID
	return pet
}
