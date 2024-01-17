package main

import (
	"petplace/back-mascotas/cmd/app/db"
	"petplace/back-mascotas/cmd/app/db/objects"
	"petplace/back-mascotas/cmd/app/routes"
	"petplace/back-mascotas/cmd/app/services"
)

func main() {

	repository := initDB()

	pp := services.NewPetPlace(&repository)
	vs := services.NewVaccineService(&repository)

	r := routes.NewRouter(":8001")
	r.AddPingRoute()
	err := r.AddPetRoutes(&pp)
	if err != nil {
		panic(err)
	}

	err = r.AddVaccineRoutes(vs)
	if err != nil {
		panic(err)
	}

	err = r.AddSwaggerRoutes()
	if err != nil {
		panic(err)
	}

	r.Run()
}

func initDB() db.Repository {

	r, err := db.NewRepository("admin:admin@tcp(localhost:3306)/pets?parseTime=true")
	if err != nil {
		panic(err)
	}
	err = r.Init([]interface{}{objects.Pet{}, objects.Vaccine{}})
	if err != nil {
		panic(err)
	}
	return r
}
