package routes

import (
	"petplace/back-mascotas/src/controller"
	"petplace/back-mascotas/src/routes/internal"
	"petplace/back-mascotas/src/services"
)

func (r *Routes) AddVeterinaryRoutes(service services.VeterinaryService) error {
	c := controller.NewVeterinaryController(service)

	group := r.engine.Group("/veterinaries", internal.AppContextCreator())

	group.POST("/veterinary", c.New)
	group.GET("/veterinary/:id", c.Get)
	group.PUT("/veterinary/:id", c.Edit)
	group.DELETE("/veterinary/:id", c.Delete)
	group.GET("", c.GetAll)

	return nil

}
