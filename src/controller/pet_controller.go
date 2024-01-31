package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"petplace/back-mascotas/src/model"
	"petplace/back-mascotas/src/requester"
	"petplace/back-mascotas/src/services"
	"strconv"
)

type PremiumPetController struct {
	ABMController[model.Pet]
	service services.PetService
	users   *requester.Requester
}

func NewPetController(service services.PetService, usersService *requester.Requester) PremiumPetController {

	temp := PremiumPetController{}
	temp.service = service
	temp.s = service
	temp.Validate = ValidateNewAnimal
	temp.name = "PET"
	temp.users = usersService
	return temp
}

func ValidateNewAnimal(pet model.Pet) error {

	if !model.ValidAnimalType(pet.Type) {
		return InvalidAnimalType
	}
	return nil
}

// New godoc
//
//	@Summary		Creates a Pet
//	@Description	Create a pet for a given user
//	@Tags			Pet
//	@Accept			json
//	@Produce		json
//	@Param			pet	body		Pet	true	"pet info"
//	@Success		201	{object}	model.Pet
//	@Failure		400	{object}	APIError
//	@Router			/pets/pet [post]
func (pc *PremiumPetController) New(c *gin.Context) {

	telegramIDStr := c.Request.Header.Get("X-Telegram-Id")
	if telegramIDStr != "" {
		apiErr := pc.handleTelegramID(c, telegramIDStr)
		if apiErr != nil {
			ReturnError(c, http.StatusBadRequest, apiErr.error, apiErr.Message)
			return
		}
	}

	pc.ABMController.New(c)
}

// Get godoc
//
//	@Summary		Get a Pet
//	@Description	Get pet info given a pet ID
//	@Tags			Pet
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"id of the pet"
//	@Success		200		{object}	model.Pet
//	@Failure		400,404	{object}	APIError
//	@Router			/pets/pet/{id} [get]
func (pc *PremiumPetController) Get(c *gin.Context) {
	pc.ABMController.Get(c)
}

// Edit godoc
//
//	@Summary		Edit a Pet
//	@Description	Edit pet info given a pet ID and pet info needed to update
//	@Tags			Pet
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"id of the pet"
//	@Param			pet		body		Pet	true	"pet info"
//	@Success		200		{object}	model.Pet
//	@Failure		400,404	{object}	APIError
//	@Router			/pets/pet/{id} [put]
func (pc *PremiumPetController) Edit(c *gin.Context) {

	telegramIDStr := c.Request.Header.Get("X-Telegram-Id")
	if telegramIDStr != "" {
		apiErr := pc.handleTelegramID(c, telegramIDStr)
		if apiErr != nil {
			ReturnError(c, http.StatusBadRequest, apiErr.error, apiErr.Message)
			return
		}
	}

	pc.ABMController.Edit(c)
}

// Delete godoc
//
//	@Summary		Delete a Pet
//	@Description	Delete a pet given a pet ID
//	@Tags			Pet
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id of the pet"
//	@Success		204	{object}	nil
//	@Failure		400	{object}	APIError
//	@Router			/pets/pet/{id} [delete]
func (pc *PremiumPetController) Delete(c *gin.Context) {
	pc.ABMController.Delete(c)
}

// GetPetsByOwner godoc
//
//	@Summary		Get pets of owner
//	@Description	Get a pet list given the owner ID
//	@Tags			Pet
//	@Accept			json
//	@Produce		json
//	@Param			owner_id	path		string	true	"owner id to get pets"
//	@Param			offset		query		int		false	"offset of the results"
//	@Param			limit		query		int		false	"limit of the results "
//	@Success		200			{object}	model.SearchResponse
//	@Failure		400,404		{object}	APIError
//	@Router			/pets/owner/{owner_id} [get]
func (pc *PremiumPetController) GetPetsByOwner(c *gin.Context) {

	ownerID, ok := c.Params.Get("owner_id")
	if !ok || ownerID == "" {
		ReturnError(c, http.StatusBadRequest, MissingParams, "expected owner_id")
		return
	}

	telegramIDStr := c.Request.Header.Get("X-Telegram-Id")
	if telegramIDStr != "" {

		if ownerID != telegramIDStr {
			ReturnError(c, http.StatusBadRequest, MissingParams, "mismatch between owner_id and X-Telegram-Id")
			return
		}

		telegramID, err := strconv.Atoi(telegramIDStr)
		if err != nil {
			ReturnError(c, http.StatusBadRequest, MissingParams, "cannot parse owner_id: "+err.Error())
			return
		}

		user, err := pc.users.GetUserData(telegramID)
		if err != nil {
			ReturnError(c, http.StatusBadRequest, MissingParams, "expected owner_id")
			return
		}
		ownerID = user.ID
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
		ReturnError(c, http.StatusNotFound, EntityNotFound, fmt.Sprintf("not found pets for owner: '%v' ", ownerID))
		return
	}

	c.JSON(http.StatusOK, response)

}

func (pc *PremiumPetController) handleTelegramID(c *gin.Context, telegramIDStr string) *APIError {

	// Parse the telegram ID to int
	telegramID, err := strconv.Atoi(telegramIDStr)
	if err != nil {
		return NewApiError(fmt.Errorf("asdasdf"), http.StatusBadRequest)
	}

	// Get the user data
	user, err := pc.users.GetUserData(telegramID)
	if err != nil {
		return NewApiError(fmt.Errorf("asdasdf"), http.StatusBadRequest)
	}

	// Set the owner ID in the request body
	err = setOwnerIDRequest(c, user.ID)
	if err != nil {
		return NewApiError(fmt.Errorf("asdasdf"), http.StatusBadRequest)
	}

	return nil
}

func setOwnerIDRequest(ctx *gin.Context, ownerID string) error {

	var pet model.Pet
	body, err := getBody(ctx)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &pet)
	if err != nil {
		return err
	}
	pet.OwnerID = ownerID

	rawPet, err := json.Marshal(pet)
	if err != nil {
		return err
	}
	reWriteBody(ctx, rawPet)

	return nil
}
