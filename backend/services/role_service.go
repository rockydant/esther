package service

import (
	"esther/models"
)

type RolesService interface {
	Create(name string) models.Role
	Update(role models.Role)
	Delete(id int)
	FindById(id int) models.Role
	FindAll() []models.Role
}
