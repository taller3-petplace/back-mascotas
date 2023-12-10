package data

import "strings"

type AnimalType string

const (
	Dog     AnimalType = "dog"
	Cat     AnimalType = "cat"
	Bird    AnimalType = "bird"
	Hamster AnimalType = "hamster"
)

var AnimalTypes = []AnimalType{Dog, Cat, Bird, Hamster}

func (t AnimalType) Normalice() AnimalType {
	return AnimalType(strings.ToLower(string(t)))
}
