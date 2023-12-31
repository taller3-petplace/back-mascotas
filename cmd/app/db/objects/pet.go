package objects

import (
	"gorm.io/gorm"
	"petplace/back-mascotas/cmd/app/model"
	"time"
)

type Pet struct {
	ID        uint `gorm:"primaryKey;autoIncrement;unique"`
	Name      string
	Type      string
	CreatedAt time.Time `gorm:"type:datetime"`
	UpdatedAt time.Time
	BirthDate time.Time `gorm:"type:datetime"`
	DeletedAt gorm.DeletedAt
	OwnerID   int
}

func (p *Pet) FromModel(pet model.Pet) {
	p.ID = pet.ID
	p.Name = pet.Name
	p.Type = string(pet.Type)
	p.CreatedAt = pet.RegisterDate
	p.BirthDate = pet.BirthDate.Time
	p.OwnerID = pet.OwnerID
}

func (p *Pet) ToModel() model.Pet {

	var pet model.Pet
	pet.ID = p.ID
	pet.Name = p.Name
	pet.Type = model.AnimalType(p.Type)
	pet.RegisterDate = p.CreatedAt
	pet.BirthDate = model.Date{Time: p.BirthDate}
	pet.OwnerID = p.OwnerID
	return pet
}
