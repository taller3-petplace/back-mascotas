package routes

import (
	"petplace/back-mascotas/cmd/app/controller"
	"petplace/back-mascotas/cmd/app/db"
	"petplace/back-mascotas/cmd/app/services"
)

func (r *Routes) AddPetRoutes() error {

	c := controller.NewPetController(services.NewPetPlace(db.NewFakeDB()))

	group := r.engine.Group("/pets")

	group.POST("/", c.NewPet)
	group.GET("/:pet_id", c.GetPet)       // TODO
	group.GET("/", c.GetPetsByOwner)      // TODO
	group.GET("/owner", c.SearchPet)      // TODO
	group.PUT("/", c.EditPet)             // TODO
	group.DELETE("/:pet_id", c.DeletePet) // TODO

	return nil
}
