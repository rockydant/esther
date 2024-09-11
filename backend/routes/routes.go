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
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swaggerFiles alias for the files package
	ginSwagger "github.com/swaggo/gin-swagger" // Swagger handler
)

func SetupRouter(roleController *controller.RoleController, userController *controller.UserController) *gin.Engine {
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

	protected.GET("users", userController.GetUsers)
	protected.GET("users/:id", userController.GetUser)
	protected.POST("users", userController.CreateUser)
	protected.PUT("users/:id", userController.UpdateUser)
	protected.DELETE("users/:id", userController.DeleteUser)

	protected.GET("/roles", roleController.GetRoles)
	protected.GET("/roles/:id", roleController.GetRole)
	protected.POST("/roles", roleController.CreateRole)
	protected.PUT("/roles/:id", roleController.UpdateRole)
	protected.DELETE("/roles/:id", roleController.DeleteRole)

	return r
}
