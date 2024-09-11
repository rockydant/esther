package repository

import (
	"errors"
	"esther/models"

	"gorm.io/gorm"
)

type RolesRepositoryImpl struct {
	Db *gorm.DB
}

func InitRolesRepositoryImpl(db *gorm.DB) RolesRepository {
	return &RolesRepositoryImpl{Db: db}
}

func (u RolesRepositoryImpl) Save(role models.Role) {
	result := u.Db.Create(&role)
	panic(result.Error)

}

func (u RolesRepositoryImpl) Update(role models.Role) {
	var updatedRole = models.Role{
		Name: role.Name,
	}
	result := u.Db.Model(&role).Updates(updatedRole)
	panic(result.Error)
}

func (u RolesRepositoryImpl) Delete(id int) {
	var role models.Role
	result := u.Db.Where("id = ?", id).Delete(&role)
	panic(result.Error)
}

func (u RolesRepositoryImpl) FindById(id int) (models.Role, error) {
	var role models.Role
	result := u.Db.Find(&role, id)
	if result != nil {
		return role, nil
	} else {
		return role, errors.New("role is not found")
	}
}

func (u RolesRepositoryImpl) FindAll() []models.Role {
	var roles []models.Role
	results := u.Db.Find(&roles)

	if results != nil {
		return roles
	} else {
		panic(results.Error)
	}
}
