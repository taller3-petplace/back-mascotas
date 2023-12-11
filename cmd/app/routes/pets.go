package routes

import (
	"petplace/back-mascotas/cmd/app/controller"
	"petplace/back-mascotas/cmd/app/services"
)

func (r *Routes) AddPetRoutes(service services.PetService) error {
	c := controller.NewPetController(service)

	group := r.engine.Group("/pets")

	group.POST("/", c.NewPet)
	group.GET("/:pet_id", c.GetPet)
	group.GET("/owner/:owner_id", c.GetPetsByOwner)
	group.GET("/owner", c.SearchPet)      // TODO
	group.PUT("/", c.EditPet)             // TODO
	group.DELETE("/:pet_id", c.DeletePet) // TODO

	return nil

}
