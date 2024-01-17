package main

import (
	"fmt"
	"petplace/back-mascotas/src/config"
	"petplace/back-mascotas/src/db"
	"petplace/back-mascotas/src/db/objects"
	"petplace/back-mascotas/src/routes"
	"petplace/back-mascotas/src/services"
)

func main() {

	appConfig, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	repository := initDB(appConfig.DbURL)

	pp := services.NewPetPlace(&repository)
	vs := services.NewVaccineService(&repository)

	r := routes.NewRouter(fmt.Sprintf(":%d", appConfig.Port))
	r.AddPingRoute()
	err = r.AddPetRoutes(&pp)
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

func initDB(url string) db.Repository {

	r, err := db.NewPostgresRepository(url)
	if err != nil {
		panic(err)
	}
	err = r.Init([]interface{}{objects.Pet{}, objects.Vaccine{}})
	if err != nil {
		panic(err)
	}
	return r
}