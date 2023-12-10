package data

import "time"

type NewPetRequest struct {
	Name      string     `json:"name" binding:"required"`
	Type      AnimalType `json:"type" binding:"required"`
	BirthDate string     `json:"birth_date" binding:"required"`
	OwnerID   string     `json:"owner_id" binding:"required"`
}

type Pet struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Type         AnimalType `json:"type"`
	RegisterDate time.Time  `json:"register_date"`
	BirthDate    time.Time  `json:"birth_date"`
	OwnerID      string     `json:"owner_id"`
}

type Race struct {
	Type string `json:"type"`
	Fur  string `json:"fur"`
	Size int32  `json:"size"`
}
