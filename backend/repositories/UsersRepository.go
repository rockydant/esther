package repository

import (
	"esther/models"
)

type UsersRepository interface {
	Save(user models.User)
	Update(user models.User)
	Delete(id int)
	FindById(id int) (user models.User, err error)
	FindAll() []models.User
}
