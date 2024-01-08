package controller

import (
	"github.com/gin-gonic/gin"
	"petplace/back-mascotas/cmd/app/model"
	"petplace/back-mascotas/cmd/app/services"
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
//	@Param			vaccine	body   model.Vaccine	true	"vaccine info"
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
//	@Param			id	path   int	true	"id of the vaccine"
//	@Success		200		{object}	model.Vaccine
//	@Failure		400		{object}	APIError
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
//	@Param			id	path   int	true	"id of the vaccine"
//	@Param			vaccine	body   model.Vaccine	true	"vaccine info"
//	@Success		200		{object}	model.Vaccine
//	@Failure		400		{object}	APIError
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
//	@Param			id	path   int	true	"id of the vaccine"
//	@Success		204		{object}	nil
//	@Failure		400		{object}	APIError
//	@Router			/vaccines/vaccine/{id} [delete]
func (vs *VaccineController) Delete(c *gin.Context) {
	vs.ABMController.Delete(c)
}
