package controller

import (
	"esther/models"
	service "esther/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleService service.RolesService
}

func InitRoleController(service service.RolesService) *RoleController {
	return &RoleController{roleService: service}
}

// @Summary List roles
// @Description list all roles
// @Accept  json
// @Produce  json
// @Tags Roles
//
// @Security JWT
// @Success 200 "ok"
// @Router /roles [get]
func (controller *RoleController) GetRoles(c *gin.Context) {
	roles := controller.roleService.FindAll()

	c.JSON(http.StatusOK, gin.H{"data": roles})
}

// @Summary Get role
// @Description Get role by id
// @Accept  json
// @Produce  json
// @Tags Roles
// @Param id path string true "Role ID"
// @Security JWT
// @Success 200 "ok"
// @Router /roles/{id} [get]
func (controller *RoleController) GetRole(c *gin.Context) {
	id := c.Param("id")
	roleId, _ := strconv.Atoi(id)

	role := controller.roleService.FindById(roleId)

	c.JSON(http.StatusOK, role)
}

// @Summary Create role
// @Description Create new role
// @Accept  json
// @Produce  json
// @Tags Roles
// @Param name	path		string	true	"Create new role"
// @Security JWT
// @Success 200 "ok"
// @Router /roles/{name} [post]
func (controller *RoleController) CreateRole(c *gin.Context) {
	name := c.Param("name")

	newRole := controller.roleService.Create(name)
	c.JSON(http.StatusCreated, newRole)
}

// @Summary Update role
// @Description Update role
// @Accept  json
// @Produce  json
// @Tags Roles
// @Param Role	body		models.Role	true	"Update role"
// @Security JWT
// @Success 200 "ok"
// @Router /roles [put]
func (controller *RoleController) UpdateRole(c *gin.Context) {
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

// @Summary Delete role
// @Description Delete role
// @Accept  json
// @Produce  json
// @Tags Roles
// @Param id path string true "Role ID"
// @Security JWT
// @Success 200 "ok"
// @Router /roles [delete]
func (controller *RoleController) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if err := models.DB.Delete(&models.Role{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
