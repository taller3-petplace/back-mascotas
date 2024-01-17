package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Routes struct {
	engine  *gin.Engine
	address string
}

func NewRouter(address string) Routes {
	return Routes{
		engine:  gin.Default(),
		address: address,
	}
}

func NewMockRouter() Routes {

	gin.SetMode(gin.TestMode)
	e := gin.Default()

	return Routes{
		engine:  e,
		address: "",
	}
}

func (r *Routes) Run() {

	err := r.engine.Run(r.address)
	if err != nil {
		panic(err)
	}
}

func (r *Routes) AddPingRoute() {

	r.engine.GET("/ping",
		func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "pong"}) })
}

func (r *Routes) ServeRequest(w http.ResponseWriter, req *http.Request) {
	r.engine.ServeHTTP(w, req)
}
