package model

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

func ValidAnimalType(animalType AnimalType) bool {

	var normalized = animalType.Normalice()
	for _, t := range AnimalTypes {
		if t == normalized {
			return true
		}
	}
	return false
}
