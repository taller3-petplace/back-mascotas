package routes

import (
	"petplace/back-mascotas/src/controller"
	"petplace/back-mascotas/src/services"
)

func (r *Routes) AddPetRoutes(service services.PetService) error {
	c := controller.NewPetController(service)

	group := r.engine.Group("/pets")

	group.POST("/pet", c.New)
	group.GET("/pet/:id", c.Get)
	group.GET("/owner/:owner_id", c.GetPetsByOwner)
	group.PUT("/pet/:id", c.Edit)
	group.DELETE("/pet/:id", c.Delete)

	return nil

}