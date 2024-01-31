package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"petplace/back-mascotas/src/config"
	"petplace/back-mascotas/src/db"
	"petplace/back-mascotas/src/db/objects"
	"petplace/back-mascotas/src/requester"
	"petplace/back-mascotas/src/routes"
	"petplace/back-mascotas/src/services"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

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
	vaccinesService := initRequester(appConfig.TreatmentURL)
	usersService := initRequester(appConfig.UsersURL)

	pp := services.NewPetPlace(&repository)
	vs := services.NewVaccineService(&repository, vaccinesService)
	vet := services.NewVeterinaryService(&repository)

	r := routes.NewRouter(fmt.Sprintf(":%d", appConfig.Port))
	r.AddPingRoute()
	r.AddMiddleware(CORSMiddleware())
	err = r.AddPetRoutes(&pp, usersService)
	if err != nil {
		panic(err)
	}

	err = r.AddVaccineRoutes(vs)
	if err != nil {
		panic(err)
	}

	err = r.AddVeterinaryRoutes(vet)
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
	err = r.Init([]interface{}{objects.Pet{}, objects.Vaccine{}, objects.Application{}, objects.Veterinary{}})
	if err != nil {
		panic(err)
	}
	return r
}

func initMockRequester(url string) *requester.Requester {
	return requester.NewRequester(requester.NewMockHttpClient(), url)
}

func initRequester(url string) *requester.Requester {
	return requester.NewRequester(requester.NewHttpClient(), url)
}
