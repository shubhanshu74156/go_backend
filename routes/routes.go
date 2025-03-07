package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		RegisterAuthRoutes(api.Group("/auth"))
	}

	return router
}
