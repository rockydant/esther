package models

import (
	"esther/config"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserDTO struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

type RegisterDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleID   uint   `json:"role_id" binding:"required"`
}

// User represents a blog post with a Username, Password, RoleID, Role.
// swagger:model
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`
	RoleID   uint   `json:"role_id"`
	Role     Role   `json:"role" gorm:"foreignKey:RoleID"`
}

type UpdatedUserDTO struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`
	RoleID   uint   `json:"role_id"`
}

// Role represents a blog post with a Name.
// swagger:model
type Role struct {
	gorm.Model
	Name string `json:"name"`
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
