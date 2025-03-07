package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubhanshu74156/go-rest-api/controllers"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	router.POST("/signup", controllers.Signup) // Signup route
	router.POST("/login", controllers.Login)   // Login route
}
