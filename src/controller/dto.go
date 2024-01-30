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

type Veterinary struct {
	Name     string   `json:"name" example:"Veterinary 1"`
	Address  string   `json:"address" example:"Av. Siempreviva 123"`
	Phone    string   `json:"phone" example:"123456789"`
	Email    string   `json:"email" example:"veterinary1@gmail.com"`
	WebSite  string   `json:"web_site" example:"www.veterinary1.com"`
	IMGUrl   string   `json:"img_url" example:"www.veterinary1.com/img.png"`
	City     string   `json:"city_id" example:"Buenos Aires"`
	Location Location `json:"location"`
	Doctors  []Doctor `json:"doctors"`
	DayGuard int      `json:"day_guard" example:"1"`
}

type Location struct {
	Latitude  float64 `json:"latitude" example:"-34.603684"`
	Longitude float64 `json:"longitude" example:"-58.381559"`
}

type Doctor struct {
	Name  string `json:"name" example:"Juan Valdez"`
	Phone string `json:"phone" example:"123456789"`
	Email string `json:"email" example:"JuanValdez@gmail.com"`
}
