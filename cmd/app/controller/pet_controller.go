package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"petplace/back-mascotas/cmd/app/controller/pet_errors"
	"petplace/back-mascotas/cmd/app/data"
	"petplace/back-mascotas/cmd/app/services"
	"time"
)

type PetController struct {
	petPlace services.PetPlace
}

func NewPetController(service services.PetPlace) PetController {
	return PetController{petPlace: service}
}

func (pc *PetController) NewPet(c *gin.Context) {

	var pet data.Pet
	err := c.BindJSON(&pet)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, pet_errors.EntityFormatError, err.Error())
		return
	}

	if !validAnimalType(pet.Type) {
		ReturnError(c, http.StatusBadRequest, pet_errors.InvalidAnimalType, fmt.Sprintf("unexpected animal '%s'", pet.Type))
		return
	}

	pet.RegisterDate = time.Now()
	pet.Type = pet.Type.Normalice()

	err = pc.petPlace.RegisterNewPet(pet)
	if err != nil {
		ReturnError(c, http.StatusInternalServerError, pet_errors.RegisterError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, pet)
}

func (pc *PetController) GetPet(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{"message": "to be implemented"})
}

func (pc *PetController) GetPetsByOwner(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{"message": "to be implemented"})
}

func (pc *PetController) SearchPet(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{"message": "to be implemented"})
}

func (pc *PetController) EditPet(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{"message": "to be implemented"})
}

func (pc *PetController) DeletePet(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{"message": "to be implemented"})
}

func validAnimalType(animalType data.AnimalType) bool {

	var normalized = animalType.Normalice()
	for _, t := range data.AnimalTypes {
		if t == normalized {
			return true
		}
	}
	return false
}
