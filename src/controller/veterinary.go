package controller

import (
	"github.com/gin-gonic/gin"
	"petplace/back-mascotas/src/model"
	"petplace/back-mascotas/src/services"
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
//	@Param			veterinary	body		Veterinary	true	"Veterinary info"
//	@Success		201			{object}	model.Veterinary
//	@Failure		400			{object}	APIError
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
//	@Param			id		path		int	true	"id of the veterinary"
//	@Success		200		{object}	model.Veterinary
//	@Failure		400,404	{object}	APIError
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
//	@Param			id			path		int			true	"id of the Veterinary"
//	@Param			veterinary	body		Veterinary	true	"Veterinary info"
//	@Success		200			{object}	model.Veterinary
//	@Failure		400,404		{object}	APIError
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
//	@Param			id	path		int	true	"id of the veterinary"
//	@Success		204	{object}	nil
//	@Failure		400	{object}	APIError
//	@Router			/veterinaries/veterinary/{id} [delete]
func (vc *VeterinaryController) Delete(c *gin.Context) {
	vc.ABMController.Delete(c)
}
