package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"petplace/back-mascotas/cmd/app/controller/pet_errors"
	"petplace/back-mascotas/cmd/app/data"
	"petplace/back-mascotas/cmd/app/services"
)

type PetController struct {
	petPlace services.PetPlace
}

func NewPetController(service services.PetPlace) PetController {
	return PetController{petPlace: service}
}

func (pc *PetController) NewPet(c *gin.Context) {

	pet := data.Pet{}
	if c.ShouldBind(&pet) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": pet_errors.EntityFormatError.Error()})
	}

	err := pc.petPlace.RegisterNewPet(pet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": pet_errors.RegisterError.Error()})
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
