// Package routes Blog API.
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//	Host: localhost:8080
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package routes

import (
	"esther/config"
	controller "esther/controllers"
	_ "esther/docs"
	"esther/middlewares"
	"esther/models"
	repository "esther/repositories"
	service "esther/services"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"     // swaggerFiles alias for the files package
	ginSwagger "github.com/swaggo/gin-swagger" // Swagger handler
)

func InitRoleController(validate *validator.Validate) *controller.RoleController {
	//Init Role Repository
	roleRepository := repository.InitRolesRepositoryImpl(models.DB)

	//Init Role Service
	roleService := service.InitRolesServiceImpl(roleRepository, validate)

	//Init Role controller
	roleController := controller.InitRoleController(roleService)

	return roleController
}

func InitUserController(validate *validator.Validate) *controller.UserController {
	//Init Repository
	userRepository := repository.InitUsersRepositoryImpl(models.DB)

	//Init Service
	userService := service.InitUsersServiceImpl(userRepository, validate)

	//Init controller
	userController := controller.InitUserController(userService)

	return userController
}

func SetupRouter(validate *validator.Validate) *gin.Engine {
	r := gin.Default()
	corsDomains := strings.Split(config.GetEnvVariable("cors"), ",")

	corsConfig := cors.Config{
		AllowOrigins:     corsDomains,                                         // List allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // List allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // List allowed headers
		AllowCredentials: true,                                                // Allow credentials (cookies, authorization headers, etc.)
	}

	r.Use(cors.New(corsConfig))
	r.Use(middlewares.CORSMiddleware())

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public routes
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)

	// Protected routes
	protected := r.Group("/")

	protected.Use(cors.New(corsConfig))
	protected.Use(middlewares.CORSMiddleware())
	protected.Use(middlewares.JWTAuthMiddleware())

	userController := InitUserController(validate)

	protected.GET("users", userController.GetUsers)
	protected.GET("users/:id", userController.GetUser)
	protected.POST("users", userController.CreateUser)
	protected.PUT("users/:id", userController.UpdateUser)
	protected.DELETE("users/:id", userController.DeleteUser)

	roleController := InitRoleController(validate)
	protected.GET("/roles", roleController.GetRoles)
	protected.GET("/roles/:id", roleController.GetRole)
	protected.POST("/roles", roleController.CreateRole)
	protected.PUT("/roles/:id", roleController.UpdateRole)
	protected.DELETE("/roles/:id", roleController.DeleteRole)

	return r
}
