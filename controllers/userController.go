package controllers

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/shubhanshu74156/go-rest-api/database"
// 	"github.com/shubhanshu74156/go-rest-api/models"
// )

// func GetUsers(c *gin.Context) {
// 	var users []models.User
// 	if err := database.DB.Find(&users).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, users)
// }

// func CreateUser(c *gin.Context) {
// 	var user models.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if err := database.DB.Create(&user).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, user)
// }
