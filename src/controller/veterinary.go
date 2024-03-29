package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"petplace/back-mascotas/src/model"
	"petplace/back-mascotas/src/services"
	"strconv"
)

type VeterinaryController struct {
	ABMController[model.Veterinary]
	service services.VeterinaryService
}

func NewVeterinaryController(s services.VeterinaryService) VeterinaryController {

	temp := VeterinaryController{}
	temp.service = s
	temp.s = s
	temp.Validate = ValidateVeterinary
	temp.name = "VETERINARY"
	return temp
}

func ValidateVeterinary(v model.Veterinary) error {
	return nil
}

// New godoc
//
//	@Summary		Creates a Veterinary
//	@Description	Create a Veterinary
//	@Tags			Veterinary
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string		true	"JWT header"
//	@Param			X-Telegram-App	header		bool		true	"request from telegram"
//	@Param			X-Telegram-Id	header		string		false	"chat id of the telegram user"
//	@Param			veterinary		body		Veterinary	true	"Veterinary info"
//	@Success		201				{object}	model.Veterinary
//	@Failure		400				{object}	APIError
//	@Router			/veterinaries/veterinary [post]
func (vc *VeterinaryController) New(c *gin.Context) {
	vc.ABMController.New(c)
}

// Get godoc
//
//	@Summary		Get a veterinary
//	@Description	Get veterinary info given a veterinary ID
//	@Tags			Veterinary
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"JWT header"
//	@Param			X-Telegram-App	header		bool	true	"request from telegram"
//	@Param			X-Telegram-Id	header		string	false	"chat id of the telegram user"
//	@Param			id				path		int		true	"id of the veterinary"
//	@Success		200				{object}	model.Veterinary
//	@Failure		400,404			{object}	APIError
//	@Router			/veterinaries/veterinary/{id} [get]
func (vc *VeterinaryController) Get(c *gin.Context) {
	vc.ABMController.Get(c)
}

// Edit godoc
//
//	@Summary		Edit a Veterinary
//	@Description	Edit Veterinary info given a veterinary ID
//	@Tags			Veterinary
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string		true	"JWT header"
//	@Param			X-Telegram-App	header		bool		true	"request from telegram"
//	@Param			X-Telegram-Id	header		string		false	"chat id of the telegram user"
//	@Param			id				path		int			true	"id of the Veterinary"
//	@Param			veterinary		body		Veterinary	true	"Veterinary info"
//	@Success		200				{object}	model.Veterinary
//	@Failure		400,404			{object}	APIError
//	@Router			/veterinaries/veterinary/{id} [put]
func (vc *VeterinaryController) Edit(c *gin.Context) {
	vc.ABMController.Edit(c)
}

// Delete godoc
//
//	@Summary		Delete a Veterinary
//	@Description	Delete a Veterinary given a veterinary ID
//	@Tags			Veterinary
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"JWT header"
//	@Param			X-Telegram-App	header		bool	true	"request from telegram"
//	@Param			X-Telegram-Id	header		string	false	"chat id of the telegram user"
//	@Param			id				path		int		true	"id of the veterinary"
//	@Success		204				{object}	nil
//	@Failure		400				{object}	APIError
//	@Router			/veterinaries/veterinary/{id} [delete]
func (vc *VeterinaryController) Delete(c *gin.Context) {
	vc.ABMController.Delete(c)
}

// GetAll godoc
//
//	@Summary		Get veterinaries
//	@Description	Get veterinaries applying filters by city, day_guard, offset and limit
//	@Tags			Veterinary
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"JWT header"
//	@Param			X-Telegram-App	header		bool	true	"request from telegram"
//	@Param			X-Telegram-Id	header		string	false	"chat id of the telegram user"
//	@Param			city			query		string	false	"city of the veterinary"
//	@Param			day_guard		query		int		false	"day_guard of the veterinary"
//	@Param			offset			query		int		false	"offset of the results"
//	@Param			limit			query		int		false	"limit of the results "
//	@Success		200				{object}	model.SearchResponse[model.Veterinary]
//	@Failure		400				{object}	APIError
//	@Router			/veterinaries [get]
func (vc *VeterinaryController) GetAll(c *gin.Context) {

	searchParams, apiErr := getSearchParams(c)
	if apiErr != nil {
		ReturnError(c, apiErr.Status, apiErr.error, apiErr.Message)
		return
	}

	filters := make(map[string]string)

	city := c.Query("city")
	if city != "" {
		filters["city"] = city
	}

	guardDay := c.Query("day_guard")
	if guardDay != "" {
		filters["day_guard"] = guardDay
	}

	response, err := vc.service.GetVeterinaries(filters, searchParams)
	if err != nil {
		ReturnError(c, http.StatusInternalServerError, ServiceError, err.Error())
		return
	}

	if len(response.Results) == 0 {
		ReturnError(c, http.StatusNotFound, EntityNotFound, fmt.Sprintf("veterinaries not found"))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetNearest godoc
//
//	@Summary		Get nearest
//	@Description	Get nearest veterinary given a latitude and longitude and optional radio
//	@Tags			Veterinary
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"JWT header"
//	@Param			X-Telegram-App	header		bool	true	"request from telegram"
//	@Param			X-Telegram-Id	header		string	false	"chat id of the telegram user"
//	@Param			latitude		query		string	true	"latitude of the point of reference"	Example(-34.603684)
//	@Param			longitude		query		string	true	"longitude of the point of reference"	Example(-34.603684)
//	@Param			radius			query		string	false	"radius of the search in km"			Example(1)
//	@Param			offset			query		int		false	"offset of the results"
//	@Param			limit			query		int		false	"limit of the results "
//	@Success		200				{object}	model.SearchResponse[model.Veterinary]
//	@Failure		400				{object}	APIError
//	@Router			/veterinaries/nearest [get]
func (vc *VeterinaryController) GetNearest(c *gin.Context) {

	searchParams, apiErr := getSearchParams(c)
	if apiErr != nil {
		ReturnError(c, apiErr.Status, apiErr.error, apiErr.Message)
		return
	}

	location, apiErr := getLocation(c)
	if apiErr != nil {
		ReturnError(c, apiErr.Status, apiErr.error, apiErr.Message)
		return
	}

	var err error
	var radius float64
	radiusStr := c.Query("radius")
	if radiusStr != "" {
		radius, err = strconv.ParseFloat(radiusStr, 64)
		if err != nil {
			ReturnError(c, http.StatusBadRequest, err, "error parsing radius")
			return
		}
	} else {
		radius = 1
	}

	var filters = map[string]string{}

	// TODO: do it more precise with haversine formula
	kmRadius := radius / 111.32

	filters["latitude <= ?"] = floatToString(location.Latitude + kmRadius)
	filters["latitude >= ?"] = floatToString(location.Latitude - kmRadius)
	filters["longitude <= ?"] = floatToString(location.Longitude + kmRadius)
	filters["longitude >= ?"] = floatToString(location.Longitude - kmRadius)

	response, err := vc.service.GetVeterinaries(filters, searchParams)
	if err != nil {
		ReturnError(c, http.StatusInternalServerError, ServiceError, err.Error())
		return
	}

	if len(response.Results) == 0 {
		ReturnError(c, http.StatusNotFound, EntityNotFound, fmt.Sprintf("veterinaries arround %f, %f not found", location.Latitude, location.Longitude))
		return
	}

	c.JSON(http.StatusOK, response)
}

func floatToString(inputNum float64) string {
	return strconv.FormatFloat(inputNum, 'f', 6, 64)
}
