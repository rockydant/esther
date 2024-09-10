package controllers

import (
	"esther/models"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Login
// @Description login
// @Accept  json
// @Produce  json
//
//	@Param			UserDTO	body		models.UserDTO	true	"Login user"
//
// @Success 200 "ok"
// @Router /login [post]
func Login(c *gin.Context) {
	var input models.UserDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		fmt.Println(err)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role.Name,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString([]byte("secret"))
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// @Summary Register
// @Description register
// @Accept  json
// @Produce  json
//
//	@Param			RegisterDTO	body		models.RegisterDTO	true	"Register new user"
//
// @Success 200 "ok"
// @Router /register [post]
func Register(c *gin.Context) {
	var input models.RegisterDTO

	// Bind and validate the input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the role exists
	var role models.Role
	if err := models.DB.First(&role, input.RoleID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role not found"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create the new user object
	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
		RoleID:   input.RoleID,
		Role:     role,
	}

	// Save the user to the database
	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Return the created user (without the password)
	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"role":     role.Name,
	})
}
