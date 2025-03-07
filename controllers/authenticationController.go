package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shubhanshu74156/go-rest-api/database"
	"github.com/shubhanshu74156/go-rest-api/helper"
	"github.com/shubhanshu74156/go-rest-api/models"
)

func Signup(c *gin.Context) {
	// Get user details from the request body
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Initialize user and organization models
	var user models.User
	var organization models.Organization

	// Check if the user already exists
	if err := database.DB.Where("email = ?", email).First(&user).Error; err == nil {
		// If user exists, return an error
		helper.SendResponse(c, http.StatusBadRequest, "User already exists", nil)
		return
	}

	// Create a new organization ID (example: you can generate UUID for an organization or fetch from other logic)
	organizationID := uuid.New().String()

	// Create the organization entry
	organization = models.Organization{
		OrganizationID: organizationID,
		Name:           email,
	}

	if err := database.DB.Create(&organization).Error; err != nil {
		helper.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	encryptPassword, err := helper.HashPassword(password)
	if err != nil {
		helper.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Create a new user entry with the assigned organization
	user = models.User{
		Email:          email,
		Password:       encryptPassword, // Make sure to hash the password before saving (e.g., bcrypt)
		Role:           "admin",         // Default role is "admin"
		OrganizationID: organizationID,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		helper.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Respond with the generated token
	helper.SendResponse(c, http.StatusOK, "Signup successfull", nil)
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm(("password"))

	var user models.User

	if err := database.DB.Where("email=?", email).First(&user).Error; err != nil {
		helper.SendResponse(c, http.StatusBadRequest, "User not registerd", nil)
	}

	if err := helper.CheckPassword(user.Password, password); err != nil {
		helper.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Generate JWT token
	token, err := helper.GenerateToken(user)
	if err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.SendResponse(c, http.StatusOK, "Login successfull", token)

}
