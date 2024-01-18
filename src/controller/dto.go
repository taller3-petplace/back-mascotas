package controller

import (
	"petplace/back-mascotas/src/model"
)

// Data transfer types defined for swagger documentation

type Pet struct {
	Name      string           `json:"name" example:"Raaida"`
	Type      model.AnimalType `json:"type" example:"dog"`
	BirthDate string           `json:"birth_date" example:"2013-05-23"`
	OwnerID   int              `json:"owner_id" example:"1"`
}

type Vaccine struct {
	Animal      model.AnimalType `json:"animal" example:"dog"`
	Name        string           `json:"name" example:"anti-rabies"`
	Description string           `json:"description" example:"vaccine to preventing rage"`
	Scheduled   uint             `json:"scheduled" example:"365"`
}
