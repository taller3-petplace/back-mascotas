package model

import "time"

type Pet struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Type         AnimalType `json:"type"`
	RegisterDate time.Time  `json:"register_date"`
	BirthDate    time.Time  `json:"birth_date"`
	OwnerID      int        `json:"owner_id"`
}

func (p Pet) IsZeroValue() bool {

	var zeroValue Pet

	result := p.ID == zeroValue.ID
	result = result && (p.Name == zeroValue.Name)
	result = result && (p.Type == zeroValue.Type)
	result = result && (p.RegisterDate == zeroValue.RegisterDate)
	result = result && (p.BirthDate == zeroValue.BirthDate)
	result = result && (p.OwnerID == zeroValue.OwnerID)

	return result
}

type Race struct {
	Type string `json:"type"`
	Fur  string `json:"fur"`
	Size int32  `json:"size"`
}
