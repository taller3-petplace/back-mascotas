package routes

import (
	"petplace/back-mascotas/cmd/app/controller"
	"petplace/back-mascotas/cmd/app/services"
)

func (r *Routes) AddPetRoutes(service services.PetService) error {
	c := controller.NewPetController(service)

	group := r.engine.Group("/pets")

	group.POST("/pet", c.NewPet)
	group.GET("/pet/:pet_id", c.GetPet)
	group.GET("/owner/:owner_id", c.GetPetsByOwner)
	group.GET("/", c.SearchPet)           // TODO
	group.PUT("/", c.EditPet)             // TODO
	group.DELETE("/:pet_id", c.DeletePet) // TODO

	return nil

}
