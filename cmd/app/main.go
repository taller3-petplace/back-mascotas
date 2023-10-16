package main

import (
	"petplace/back-mascotas/cmd/app/routes"
)

func main() {

	r := routes.NewRouter(":8001")
	
	r.AddPingRoute()
	err := r.AddPetRoutes()
	if err != nil {
		panic(err)
	}

	r.Run()
}
