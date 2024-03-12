package main

import (
	docs "fetch-service/docs"
	receipts "fetch-service/internal/handlers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	app := gin.New()
	// define basepath
	docs.SwaggerInfo.BasePath = "/api/"
	router := app.Group("/api")
	// add routes
	receipts.AddRoutes(router)
	// add swagger
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// run app
	app.Run("0.0.0.0:8083")
}
