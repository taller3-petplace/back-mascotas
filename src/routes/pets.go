package routes

import (
	"petplace/back-mascotas/src/controller"
	"petplace/back-mascotas/src/requester"
	"petplace/back-mascotas/src/routes/internal"
	"petplace/back-mascotas/src/services"
)

func (r *Routes) AddPetRoutes(service services.PetService, userService *requester.Requester) error {
	c := controller.NewPetController(service, userService)

	group := r.engine.Group("/pets", internal.AppContextCreator())

	group.POST("/pet", c.New)
	group.GET("/pet/:id", c.Get)
	group.GET("/owner/:owner_id", c.GetPetsByOwner)
	group.PUT("/pet/:id", c.Edit)
	group.DELETE("/pet/:id", c.Delete)

	return nil

}
