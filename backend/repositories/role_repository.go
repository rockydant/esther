package repository

import (
	"esther/models"
)

type RolesRepository interface {
	Save(role models.Role)
	Update(role models.Role)
	Delete(id int)
	FindById(id int) (role models.Role, err error)
	FindAll() []models.Role
}
