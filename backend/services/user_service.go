package service

import (
	"esther/models"
)

type UsersService interface {
	Create(user models.RegisterDTO)
	Update(user models.UpdatedUserDTO)
	Delete(id int)
	FindById(id int) models.User
	FindAll() []models.User
}
