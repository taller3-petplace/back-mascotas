package main

import (
	"petplace/back-mascotas/cmd/app/db"
	"petplace/back-mascotas/cmd/app/db/objects"
	"petplace/back-mascotas/cmd/app/routes"
	"petplace/back-mascotas/cmd/app/services"
)

func main() {

	s := CreateService()
	r := routes.NewRouter(":8001")

	r.AddPingRoute()
	err := r.AddPetRoutes(&s)
	if err != nil {
		panic(err)
	}

	r.Run()
}

func CreateService() services.PetPlace {
	r, err := db.NewRepository("admin:admin@tcp(localhost:3306)/pets?parseTime=true")
	if err != nil {
		panic(err)
	}
	err = r.Init([]interface{}{objects.Pet{}})
	if err != nil {
		panic(err)
	}
	service := services.NewPetPlace(&r)
	return service
}
