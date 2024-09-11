package repository

import (
	"errors"
	"esther/models"

	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func InitUsersRepositoryImpl(db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: db}
}

func (u UsersRepositoryImpl) Save(users models.User) {
	result := u.Db.Create(&users)
	panic(result.Error)

}

func (u UsersRepositoryImpl) Update(users models.User) {
	var updateUser = models.UpdatedUserDTO{
		Username: users.Username,
		Password: users.Password,
		RoleID:   users.RoleID,
	}
	result := u.Db.Model(&users).Updates(updateUser)
	panic(result.Error)
}

func (u UsersRepositoryImpl) Delete(id int) {
	var user models.User
	result := u.Db.Where("id = ?", id).Delete(&user)
	panic(result.Error)
}

func (u UsersRepositoryImpl) FindById(id int) (models.User, error) {
	var user models.User
	result := u.Db.Find(&user, id)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

func (u UsersRepositoryImpl) FindAll() []models.User {
	var users []models.User
	results := u.Db.Find(&users)

	if results != nil {
		return users
	} else {
		panic(results.Error)
	}
}
