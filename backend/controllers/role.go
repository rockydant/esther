package controllers

import (
	"esther/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

// GetRole handles the GET request to retrieve a single role by ID
func GetRole(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	if err := models.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	c.JSON(http.StatusOK, role)
}

// CreateRole handles the POST request to create a new role
func CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, role)
}

// UpdateRole handles the PUT request to update an existing role by ID
func UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	if err := models.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Save(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, role)
}

// DeleteRole handles the DELETE request to delete a role by ID
func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if err := models.DB.Delete(&models.Role{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
