package controllers

import (
	"esther/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Summary List users
// @Description list all users
// @Accept  json
// @Produce  json
// @Tags Users
//
// @Security JWT
// @Success 200 "ok"
// @Router /api/users [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	models.DB.Preload("Role").Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// @Summary Get user
// @Description Get user by id
// @Accept  json
// @Produce  json
// @Tags Users
// @Param id path string true "User ID"
// @Security JWT
// @Success 200 "ok"
// @Router /api/user [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := models.DB.Preload("Role").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary Create user
// @Description Create new user
// @Accept  json
// @Produce  json
// @Tags Users
// @Param RegisterDTO	body		models.RegisterDTO	true	"Register new user"
// @Security JWT
// @Success 200 "ok"
// @Router /api/user [post]
func CreateUser(c *gin.Context) {
	var input models.RegisterDTO

	// Bind and validate the input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the role exists
	var role models.Role
	if err := models.DB.First(&role, input.RoleID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role not found"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create the new user object
	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
		RoleID:   input.RoleID,
		Role:     role,
	}

	// Save the user to the database
	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Return the created user (without the password)
	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"role":     role.Name,
	})
}

// @Summary Update user
// @Description Update user
// @Accept  json
// @Produce  json
// @Tags Users
// @Param UpdatedUserDTO	body		models.UpdatedUserDTO	true	"Update user"
// @Security JWT
// @Success 200 "ok"
// @Router /api/user [put]
// UpdateUser handles the PUT request to update an existing user by ID
func UpdateUser(c *gin.Context) {
	var updatedUser models.UpdatedUserDTO
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Preload("Role").First(&user, updatedUser.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
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

// @Summary Delete user
// @Description Delete user
// @Accept  json
// @Produce  json
// @Tags Users
// @Param id path string true "User ID"
// @Security JWT
// @Success 200 "ok"
// @Router /api/user [delete]
// DeleteUser handles the DELETE request to delete a user by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := models.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
