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

	var request data.NewPetRequest
	err := c.BindJSON(&request)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, pet_errors.EntityFormatError, err.Error())
		return
	}

	if !validAnimalType(request.Type) {
		ReturnError(c, http.StatusBadRequest, pet_errors.InvalidAnimalType, fmt.Sprintf("unexpected animal '%s'", request.Type))
		return
	}

	birthDate, err := time.Parse(time.DateOnly, request.BirthDate)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, pet_errors.InvalidBirthDate, fmt.Sprintf("error parsing birth_date"))
		return
	}

	var newPet data.Pet
	newPet.Type = request.Type.Normalice()
	newPet.Name = request.Name
	newPet.OwnerID = request.OwnerID
	newPet.BirthDate = birthDate

	newPet, err = pc.petPlace.RegisterNewPet(newPet)
	if err != nil {
		ReturnError(c, http.StatusInternalServerError, pet_errors.RegisterError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, newPet)
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
