package controllers

import (
	"esther/models"
	"net/http"

	"github.com/gin-gonic/gin"
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

// swagger:route GET /api/roles roles GetRoles
//
// GetRoles returns all roles.
//
// Responses:
//
//	200: successResponse
func GetRoles(c *gin.Context) {
	var roles []models.Role
	models.DB.Find(&roles)
	c.JSON(http.StatusOK, gin.H{"data": roles})
}
