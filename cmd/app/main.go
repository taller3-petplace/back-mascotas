package main

import (
	"petplace/back-mascotas/cmd/app/db"
	"petplace/back-mascotas/cmd/app/routes"
	"petplace/back-mascotas/cmd/app/services"
)

func main() {

	s := CreateService()
	r := routes.NewRouter(":8712")

	r.AddPingRoute()
	err := r.AddPetRoutes(&s)
	if err != nil {
		panic(err)
	}

	r.Run()
}

func CreateService() services.PetPlace {
	fakeDB := db.NewFakeDB()
	fakeDB.Init()
	service := services.NewPetPlace(&fakeDB)
	return service
}
