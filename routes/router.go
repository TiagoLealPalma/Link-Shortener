package routes

import (
	"Link-Shortener/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Endpoints
	r.POST("/shorten", controllers.ShortenURL)

	r.GET("/:token", controllers.Redirect)

	return r
}
