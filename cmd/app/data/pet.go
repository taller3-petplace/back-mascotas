package data

import "time"

type Pet struct {
	Id           int        `json:"id"`
	Name         string     `json:"name"`
	Type         AnimalType `json:"type"`
	RegisterDate time.Time  `json:"register_date"`
	BirthDate    time.Time  `json:"birth_date"`
	Owner        string     `json:"owner"`
}

type Race struct {
	Type string `json:"type"`
	Fur  string `json:"fur"`
	Size int32  `json:"size"`
}
