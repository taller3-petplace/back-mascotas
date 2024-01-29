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

const (
	outputJustApplied = "applied"
	outputJustPending = "pending"
)

type VaccineController struct {
	ABMController[model.Vaccine]
	service services.VaccineService
}

func NewVaccineController(service services.VaccineService) VaccineController {

	temp := VaccineController{}
	temp.service = service
	temp.s = &service
	temp.Validate = ValidateVaccine
	temp.name = "VACCINE"
	return temp
}

func ValidateVaccine(v model.Vaccine) error {

	if !model.ValidAnimalType(v.Animal) {
		return InvalidAnimalType
	}
	return nil
}

// New godoc
//
//	@Summary		Creates a Vaccine
//	@Description	Create a vaccine
//	@Tags			Vaccine
//	@Accept			json
//	@Produce		json
//	@Param			vaccine	body		Vaccine	true	"vaccine info"
//	@Success		201		{object}	model.Vaccine
//	@Failure		400		{object}	APIError
//	@Router			/vaccines/vaccine [post]
func (vs *VaccineController) New(c *gin.Context) {
	vs.ABMController.New(c)
}

// Get godoc
//
//	@Summary		Get a Vaccine
//	@Description	Get vaccine info
//	@Tags			Vaccine
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"id of the vaccine"
//	@Success		200		{object}	model.Vaccine
//	@Failure		400,404	{object}	APIError
//	@Router			/vaccines/vaccine/{id} [get]
func (vs *VaccineController) Get(c *gin.Context) {
	vs.ABMController.Get(c)
}

// Edit godoc
//
//	@Summary		Edit a Vaccine
//	@Description	Edit vaccine info given a pet ID and vaccine info needed to update
//	@Tags			Vaccine
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int		true	"id of the vaccine"
//	@Param			vaccine	body		Vaccine	true	"vaccine info"
//	@Success		200		{object}	model.Vaccine
//	@Failure		400,404	{object}	APIError
//	@Router			/vaccines/vaccine/{id} [put]
func (vs *VaccineController) Edit(c *gin.Context) {
	vs.ABMController.Edit(c)
}

// Delete godoc
//
//	@Summary		Delete a Vaccine
//	@Description	Delete a Vaccine given a pet ID
//	@Tags			Vaccine
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id of the vaccine"
//	@Success		204	{object}	nil
//	@Failure		400	{object}	APIError
//	@Router			/vaccines/vaccine/{id} [delete]
func (vs *VaccineController) Delete(c *gin.Context) {
	vs.ABMController.Delete(c)
}

// ApplyVaccineToPet godoc
//
//	@Summary		Apply vaccine
//	@Description	Apply vaccine to a given pet
//	@Tags			Vaccine
//	@Accept			json
//	@Produce		json
//	@Param			vaccine_id	path		int	true	"vaccine id to apply on pet"
//	@Param			pet_id		path		int	true	"pet id of the target pet"
//	@Success		201			{object}	nil
//	@Failure		400,404		{object}	APIError
//	@Router			/vaccines/apply/{vaccine_id}/to/{pet_id} [post]
func (vs *VaccineController) ApplyVaccineToPet(c *gin.Context) {

	vaccineIDStr, ok := c.Params.Get("id")
	if !ok || vaccineIDStr == "" {
		ReturnError(c, http.StatusBadRequest, MissingParams, "expected vaccine_id")
		return
	}
	vaccineID, err := strconv.Atoi(vaccineIDStr)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, MissingParams, "cannot parse vaccine_id: "+err.Error())
		return
	}

	petIDStr, ok := c.Params.Get("pet_id")
	if !ok || petIDStr == "" {
		ReturnError(c, http.StatusBadRequest, MissingParams, "expected pet_id")
		return
	}
	petID, err := strconv.Atoi(petIDStr)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, MissingParams, "cannot parse pet_id: "+err.Error())
		return
	}

	err = vs.service.ApplyVaccine(uint(petID), uint(vaccineID))
	if err != nil {
		log.Debugf(logTemplate, vs.name, "APPLYVACCINE", err)
		ReturnError(c, http.StatusInternalServerError, RegisterError, err.Error())
		return
	}

	log.Debugf(logTemplate, vs.name, "APPLYVACCINE", "success")

	response := struct {
		PetId     int `json:"pet_id"`
		VaccineId int `json:"vaccine_id"`
	}{
		PetId:     petID,
		VaccineId: vaccineID,
	}

	c.JSON(http.StatusCreated, response)
}

// GetVaccinationPlan godoc
//
//	@Summary		Get vaccination plan
//	@Description	Get the vaccination plan of given pet_id
//	@Tags			Vaccine
//	@Accept			json
//	@Produce		json
//	@Param			pet_id	path		int		true	"pet id to get vaccines"
//	@Param			output	query		string	false	"desired formant for the output"	Enums(applied, pending)
//	@Success		200		{object}	model.VaccinationPlan
//	@Failure		400,404	{object}	APIError
//	@Router			/vaccines/plan/{pet_id} [get]
func (vs *VaccineController) GetVaccinationPlan(c *gin.Context) {

	petIDStr, ok := c.Params.Get("pet_id")
	if !ok || petIDStr == "" {
		ReturnError(c, http.StatusBadRequest, MissingParams, "expected pet_id")
		return
	}
	petID, err := strconv.Atoi(petIDStr)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, MissingParams, "cannot parse pet_id: "+err.Error())
		return
	}

	output := c.Query("output")
	if output != "" && output != outputJustPending && output != outputJustApplied {
		errMsg := fmt.Sprintf("invalid output '%s'. Select '%s', '%s' or avoid it.", output, outputJustApplied, outputJustPending)
		ReturnError(c, http.StatusBadRequest, MissingParams, errMsg)
		return
	}

	plan, err := vs.service.GetPlanVaccination(petID)
	if err != nil {
		log.Debugf(logTemplate, vs.name, "APPLYVACCINE", err)
		ReturnError(c, http.StatusInternalServerError, ServiceError, err.Error())
		return
	}
	if len(plan.Pending) == 0 && len(plan.Applied) == 0 {
		ReturnError(c, http.StatusNotFound, EntityNotFound, fmt.Sprintf("not found vaccines for pet: '%d' ", petID))
		return
	}

	// Formatting response

	switch output {
	case outputJustApplied:
		c.JSON(http.StatusOK, plan.Applied)
		return

	case outputJustPending:
		c.JSON(http.StatusOK, plan.Pending)
		return
	}

	c.JSON(http.StatusOK, plan)
}
