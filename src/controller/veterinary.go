package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"petplace/back-mascotas/src/model"
	"petplace/back-mascotas/src/services"
	"strconv"
)

type VeterinaryController struct {
	s    services.VeterinaryService
	name string
}

func NewVeterinaryController(s services.VeterinaryService) VeterinaryController {

	temp := VeterinaryController{}
	temp.s = s
	temp.name = "VETERINARY"
	return temp
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
	log.Debugf(logTemplate, vc.name, "NEW", fmt.Sprintf("new request | body: %v", getBodyString(c)))

	var e model.Veterinary
	err := c.BindJSON(&e)
	if err != nil {
		log.Debugf(logTemplate, vc.name, "NEW", err)
		ReturnError(c, http.StatusBadRequest, EntityFormatError, err.Error())
		return
	}

	e, err = vc.s.New(e)
	if err != nil {
		log.Debugf(logTemplate, vc.name, "NEW", err)
		ReturnError(c, http.StatusInternalServerError, RegisterError, err.Error())
		return
	}

	log.Debugf(logTemplate, vc.name, "NEW", fmt.Sprintf("success | response: %v", e))

	c.JSON(http.StatusCreated, e)
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
	idStr, ok := c.Params.Get(IDParamName)
	if !ok || idStr == "" {
		log.Debugf(logTemplate, vc.name, "GET", "expected entity id")
		ReturnError(c, http.StatusBadRequest, MissingParams, "expected entity id")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Debugf(logTemplate, vc.name, "GET", "cannot parse id: "+err.Error())
		ReturnError(c, http.StatusBadRequest, MissingParams, "cannot parse id: "+err.Error())
		return
	}

	e, err := vc.s.Get(id)
	if err != nil {
		log.Errorf(logTemplate, vc.name, "GET", err)
		ReturnError(c, http.StatusInternalServerError, ServiceError, err.Error())
		return
	}

	if e.IsZeroValue() {
		log.Debugf(logTemplate, vc.name, "GET", "entity not found")
		ReturnError(c, http.StatusNotFound, EntityNotFound, fmt.Sprintf("entity id '%d' not found", id))
		return
	}
	log.Debugf(logTemplate, vc.name, "NEW", fmt.Sprintf("success | response: %v", e))
	c.JSON(http.StatusOK, e)
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
	log.Debugf(logTemplate, vc.name, "EDIT", fmt.Sprintf("edit request | body: %s", getBodyString(c)))

	idStr, ok := c.Params.Get(IDParamName)
	if !ok || idStr == "" {
		log.Debugf(logTemplate, vc.name, "EDIT", "expected entity id")
		ReturnError(c, http.StatusBadRequest, MissingParams, "expected entity id")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Debugf(logTemplate, vc.name, "EDIT", "cannot parse id: "+err.Error())
		ReturnError(c, http.StatusBadRequest, MissingParams, "cannot parse id: "+err.Error())
		return
	}

	e, err := vc.s.Get(id)
	if err != nil {
		log.Errorf(logTemplate, vc.name, "GET", err)
		ReturnError(c, http.StatusInternalServerError, RegisterError, err.Error())
		return
	}
	if e.IsZeroValue() {
		log.Debugf(logTemplate, vc.name, "EDIT", "entity not found")
		ReturnError(c, http.StatusNotFound, EntityNotFound, fmt.Sprintf("entity id '%d' not found", id))
		return
	}

	err = c.BindJSON(&e)
	if err != nil {
		log.Debugf(logTemplate, vc.name, "EDIT", err)
		ReturnError(c, http.StatusBadRequest, EntityFormatError, err.Error())
		return
	}

	e, err = vc.s.Edit(id, e)
	if err != nil {
		log.Errorf(logTemplate, vc.name, "EDIT", err)
		ReturnError(c, http.StatusInternalServerError, RegisterError, err.Error())
		return
	}

	log.Debugf(logTemplate, vc.name, "NEW", fmt.Sprintf("success | response: %v", e))
	c.JSON(http.StatusOK, e)
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
	idStr, ok := c.Params.Get(IDParamName)
	if !ok || idStr == "" {
		log.Debugf(logTemplate, vc.name, "DELETE", "expected entity id")
		ReturnError(c, http.StatusBadRequest, MissingParams, "expected entity id")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Debugf(logTemplate, vc.name, "DELETE", err)
		ReturnError(c, http.StatusBadRequest, MissingParams, "cannot parse id: "+err.Error())
		return
	}

	vc.s.Delete(id)
	log.Debugf(logTemplate, vc.name, "DELETE", "success")
	c.JSON(http.StatusNoContent, nil)
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

	response, err := vc.s.GetVeterinaries(filters, searchParams)
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
