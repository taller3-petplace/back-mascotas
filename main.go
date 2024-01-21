package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
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

	err = config.InitLogger(appConfig.LogLevel)
	if err != nil {
		log.Error(err)
	}
	log.Info("Log level: ", log.GetLevel())

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
	err = r.Init([]interface{}{objects.Pet{}, objects.Vaccine{}, objects.Application{}})
	if err != nil {
		panic(err)
	}
	return r
}
