package service

import (
	"esther/models"
	repository "esther/repositories"

	"github.com/go-playground/validator/v10"
)

type RolesServiceImpl struct {
	RolesRepository repository.RolesRepository
	Validate        *validator.Validate
}

// Create implements UsersService.
func (r *RolesServiceImpl) Create(name string) models.Role {

	userModel := models.Role{
		Name: name,
	}

	r.RolesRepository.Save(userModel)

	return userModel
}

// Delete implements UsersService.
func (r *RolesServiceImpl) Delete(id int) {
	r.RolesRepository.Delete(id)
}

// FindAll implements UsersService.
func (r *RolesServiceImpl) FindAll() []models.Role {
	result := r.RolesRepository.FindAll()

	return result
}

// FindById implements UsersService.
func (r *RolesServiceImpl) FindById(id int) models.Role {
	foundUser, err := r.RolesRepository.FindById(id)

	if err != nil {
		panic(err)
	}

	return foundUser
}

// Update implements UsersService.
func (r *RolesServiceImpl) Update(role models.Role) {
	foundRole, err := r.RolesRepository.FindById(int(role.ID))
	if err != nil {
		panic(err)
	}

	foundRole.Name = role.Name

	r.RolesRepository.Update(foundRole)
}

func InitRolesServiceImpl(roleRepository repository.RolesRepository, validate *validator.Validate) RolesService {
	return &RolesServiceImpl{
		RolesRepository: roleRepository,
		Validate:        validate,
	}
}
