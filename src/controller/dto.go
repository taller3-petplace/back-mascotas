package controller

import (
	"petplace/back-mascotas/src/model"
	"time"
)

// Data transfer types defined for swagger documentation

type Pet struct {
	Name      string           `json:"name" example:"Raaida"`
	Type      model.AnimalType `json:"type" example:"dog"`
	BirthDate string           `json:"birth_date" example:"2013-05-23"`
	OwnerID   string           `json:"owner_id" example:"aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"`
}

type Vaccine struct {
	Animal      model.AnimalType `json:"animal" example:"dog"`
	Name        string           `json:"name" example:"anti-rabies"`
	Description string           `json:"description" example:"vaccine to preventing rage"`
	Scheduled   uint             `json:"scheduled" example:"365"`
}

type Applications struct {
	PetID    int                   `json:"pet_id"`
	OwnerID  string                `json:"owner_id"`
	PetName  string                `json:"pet_name"`
	Vaccines map[time.Time]Vaccine `json:"vaccines"`
}

type OutputFormat string

var formats = []OutputFormat{" asdf", "asdf"}
