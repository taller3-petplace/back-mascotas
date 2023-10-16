package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Routes struct {
	engine *gin.Engine
	port   string
}

func NewRouter(port string) Routes {
	return Routes{
		engine: gin.Default(),
		port:   port,
	}
}

func (r *Routes) Run() {

	err := r.engine.Run(r.port)
	if err != nil {
		panic(err)
	}
}

func (r Routes) AddPingRoute() {

	r.engine.GET("/ping",
		func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "pong"}) })
}
