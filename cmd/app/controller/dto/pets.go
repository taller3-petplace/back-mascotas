package dto

import "petplace/back-mascotas/cmd/app/model"

type NewPetRequest struct {
	Name      string           `json:"name" binding:"required"`
	Type      model.AnimalType `json:"type" binding:"required"`
	BirthDate string           `json:"birth_date" binding:"required"`
	OwnerID   int              `json:"owner_id" binding:"required"`
}
