package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"petplace/back-mascotas/cmd/app/model"
	"petplace/back-mascotas/cmd/app/services"
	"strconv"
)

type PremiumPetController struct {
	ABMController[model.Pet]
	service services.PetService
}

func NewPetController(service services.PetService) PremiumPetController {

	temp := PremiumPetController{}
	temp.service = service
	temp.s = service
	temp.Validate = ValidateNewAnimal

	return temp
}

func ValidateNewAnimal(pet model.Pet) error {

	if !validAnimalType(pet.Type) {
		return InvalidAnimalType
	}
	return nil
}

func (pc *PremiumPetController) GetPetsByOwner(c *gin.Context) {

	ownerIDStr, ok := c.Params.Get("owner_id")
	if !ok || ownerIDStr == "" {
		ReturnError(c, http.StatusBadRequest, MissingParams, "expected owner_id")
		return
	}
	ownerID, err := strconv.Atoi(ownerIDStr)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, MissingParams, "cannot parse owner_id: "+err.Error())
		return
	}

	searchRequest := model.NewSearchRequest()
	offsetStr := c.Query("offset")
	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			ReturnError(c, http.StatusBadRequest, MissingParams, "cannot parse offset: "+err.Error())
			return
		}
		searchRequest.Offset = uint(offset)
	}

	limitStr := c.Query("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			ReturnError(c, http.StatusBadRequest, MissingParams, "cannot parse limit: "+err.Error())
			return
		}
		searchRequest.Limit = uint(limit)
	}

	searchRequest.OwnerId = ownerID
	response, err := pc.service.GetPetsByOwner(searchRequest)
	if err != nil {
		ReturnError(c, http.StatusInternalServerError, ServiceError, err.Error())
		return
	}

	if len(response.Results) == 0 {
		ReturnError(c, http.StatusNotFound, EntityNotFound, fmt.Sprintf("not found pets for owner: '%d' ", ownerID))
		return
	}

	c.JSON(http.StatusOK, response)

}
