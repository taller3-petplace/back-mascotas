package routes

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"petplace/back-mascotas/cmd/app/docs"
)

func (r *Routes) AddSwaggerRoutes() error {

	docs.SwaggerInfo.Title = "Swagger Pets API"
	group := r.engine.Group("/swagger")
	group.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return nil
}
