package routes

import (
	"petplace/back-mascotas/src/controller"
	"petplace/back-mascotas/src/services"
)

func (r *Routes) AddVaccineRoutes(service services.VaccineService) error {
	c := controller.NewVaccineController(service)

	group := r.engine.Group("/vaccines")

	group.POST("/vaccine", c.New)
	group.GET("/vaccine/:id", c.Get)
	group.PUT("/vaccine/:id", c.Edit)
	group.DELETE("/vaccine/:id", c.Delete)

	return nil

}
