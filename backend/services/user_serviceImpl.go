package service

import (
	"esther/models"
	repository "esther/repositories"

	"github.com/go-playground/validator/v10"
)

type UsersServiceImpl struct {
	UsersRepository repository.UsersRepository
	Validate        *validator.Validate
}

// Create implements UsersService.
func (u *UsersServiceImpl) Create(user models.RegisterDTO) {
	err := u.Validate.Struct(user)

	if err != nil {
		panic(err)
	}

	userModel := models.User{
		Username: user.Username,
		Password: user.Password,
		RoleID:   user.RoleID,
	}

	u.UsersRepository.Save(userModel)
}

// Delete implements UsersService.
func (u *UsersServiceImpl) Delete(id int) {
	u.UsersRepository.Delete(id)
}

// FindAll implements UsersService.
func (u *UsersServiceImpl) FindAll() []models.User {
	result := u.UsersRepository.FindAll()

	return result
}

// FindById implements UsersService.
func (u *UsersServiceImpl) FindById(id int) models.User {
	foundUser, err := u.UsersRepository.FindById(id)

	if err != nil {
		panic(err)
	}

	return foundUser
}

// Update implements UsersService.
func (u *UsersServiceImpl) Update(user models.UpdatedUserDTO) {
	foundUser, err := u.UsersRepository.FindById(int(user.UserID))
	if err != nil {
		panic(err)
	}

	foundUser.Username = user.Username
	foundUser.Password = user.Password
	foundUser.RoleID = user.RoleID

	u.UsersRepository.Update(foundUser)
}

func InitUsersServiceImpl(userReppository repository.UsersRepository, validate *validator.Validate) UsersService {
	return &UsersServiceImpl{
		UsersRepository: userReppository,
		Validate:        validate,
	}
}
