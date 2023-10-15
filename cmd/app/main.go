package main

import (
	"github.com/gin-gonic/gin"
	"petplace/back-mascotas/cmd/app/controller"
	"petplace/back-mascotas/cmd/app/db"
	"petplace/back-mascotas/cmd/app/services"
)

func main() {

	router := gin.Default()

	c := controller.NewPetController(services.NewPetPlace(db.NewFakeDB()))

	router.POST("/pets/", c.NewPet)
	router.GET("/pets/:pet_id", c.GetPet)       // TODO
	router.GET("/pets/", c.GetPetsByOwner)      // TODO
	router.GET("/pets/Owner", c.SearchPet)      // TODO
	router.PUT("/pets/", c.EditPet)             // TODO
	router.DELETE("/pets/:pet_id", c.DeletePet) // TODO

	err := router.Run(":8001")
	if err != nil {
		return
	}
}
