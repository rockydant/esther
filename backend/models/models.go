package models

import (
	"esther/config"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserDTO struct {
	Username string `json:"Username" gorm:"unique"`
	Password string `json:"Password"`
}

type RegisterDTO struct {
	Username string `json:"Username" binding:"required"`
	Password string `json:"Password" binding:"required"`
	RoleID   uint   `json:"Role_id" binding:"required"`
}

// User represents a blog post with a Username, Password, RoleID, Role.
// swagger:model
type User struct {
	gorm.Model
	Username string `json:"Username" gorm:"unique"`
	Password string `json:"-"`
	RoleID   uint   `json:"Role_id"`
	Role     Role   `json:"Role" gorm:"foreignKey:RoleID"`
}

type UpdatedUserDTO struct {
	UserID   uint   `json:"User_id"`
	Username string `json:"Username" gorm:"unique"`
	Password string `json:"Password"`
	RoleID   uint   `json:"Role_id"`
}

// Role represents a blog post with a Name.
// swagger:model
type Role struct {
	gorm.Model
	Name string `json:"Name"`
}

var DB *gorm.DB

func InitDB() {
	DB = config.DB
}

func CreateDefaultRoleAndAdmin() {
	db := DB

	var adminRole Role
	var adminUser User

	// Check if the admin role exists
	if err := db.Where("name = ?", "admin").First(&adminRole).Error; err != nil {
		// Role does not exist, create it
		adminRole = Role{Name: "admin"}
		if err := db.Create(&adminRole).Error; err != nil {
			log.Fatalf("failed to create admin role: %v", err)
		}
		fmt.Println("Default admin role created")
	}

	// Check if the admin user exists
	if err := db.Where("username = ?", "admin").Preload("Role").First(&adminUser).Error; err != nil {
		// User does not exist, create it
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("M00nlight!$"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("failed to hash password: %v", err)
		}

		adminUser = User{
			Username: "admin",
			Password: string(hashedPassword),
			RoleID:   adminRole.ID,
			Role:     adminRole,
		}
		if err := db.Create(&adminUser).Error; err != nil {
			log.Fatalf("failed to create default admin user: %v", err)
		}

		fmt.Println("Default admin user created")
	}
}
