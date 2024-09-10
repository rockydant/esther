package controllers

import (
	"esther/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// swagger:route GET /api/users users GetUsers
//
// GetUsers returns all users.
//
// Responses:
//
//	200: successResponse
func GetUsers(c *gin.Context) {
	var users []models.User
	models.DB.Preload("Role").Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GetUser handles the GET request to retrieve a single user by ID
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := models.DB.Preload("Role").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser handles the POST request to create a new user
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// UpdateUser handles the PUT request to update an existing user by ID
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := models.DB.Preload("Role").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password if it's being updated
	if updatedUser.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		updatedUser.Password = string(hashedPassword)
	}

	// Update user fields
	user.Username = updatedUser.Username
	user.Password = updatedUser.Password
	user.RoleID = updatedUser.RoleID

	if err := models.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser handles the DELETE request to delete a user by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := models.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
