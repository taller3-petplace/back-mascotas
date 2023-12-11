package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"petplace/back-mascotas/cmd/app/controller/pet_errors"
	"petplace/back-mascotas/cmd/app/data"
	"petplace/back-mascotas/cmd/app/services"
	"strconv"
	"time"
)

type PetController struct {
	petPlace services.PetService
}

func NewPetController(service services.PetService) PetController {
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

	petIDStr, ok := c.Params.Get("pet_id")
	if !ok || petIDStr == "" {
		ReturnError(c, http.StatusBadRequest, pet_errors.MissingParams, "expected pet_id")
		return
	}

	petID, err := strconv.Atoi(petIDStr)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, pet_errors.MissingParams, "cannot parse pet_id: "+err.Error())
		return
	}

	newPet, err := pc.petPlace.GetPet(petID)
	if err != nil {
		ReturnError(c, http.StatusInternalServerError, pet_errors.ServiceError, err.Error())
		return
	}

	if newPet.IsZeroValue() {
		ReturnError(c, http.StatusNotFound, pet_errors.PetNotFound, fmt.Sprintf("pet with id '%d' not found", petID))
		return
	}

	c.JSON(http.StatusOK, newPet)
}

func (pc *PetController) GetPetsByOwner(c *gin.Context) {

	ownerID, ok := c.Params.Get("owner_id")
	if !ok || ownerID == "" {
		ReturnError(c, http.StatusBadRequest, pet_errors.MissingParams, "expected owner_id")
		return
	}

	searchRequest := data.NewSearchRequest()
	offsetStr := c.Query("offset")
	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			ReturnError(c, http.StatusBadRequest, pet_errors.MissingParams, "cannot parse offset: "+err.Error())
			return
		}
		searchRequest.Offset = uint(offset)
	}

	limitStr := c.Query("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			ReturnError(c, http.StatusBadRequest, pet_errors.MissingParams, "cannot parse limit: "+err.Error())
			return
		}
		searchRequest.Limit = uint(limit)
	}

	searchRequest.OwnerId = ownerID
	response, err := pc.petPlace.GetPetsByOwner(searchRequest)
	if err != nil {
		ReturnError(c, http.StatusInternalServerError, pet_errors.ServiceError, err.Error())
		return
	}

	if len(response.Results) == 0 {
		ReturnError(c, http.StatusNotFound, pet_errors.PetNotFound, fmt.Sprintf("not found pets for owner: '%s' ", ownerID))
		return
	}

	c.JSON(http.StatusOK, response)

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
