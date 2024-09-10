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
	"esther/controllers"
	_ "esther/docs"
	"esther/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swaggerFiles alias for the files package
	ginSwagger "github.com/swaggo/gin-swagger" // Swagger handler
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "https://your-frontend-domain.com"}, // List allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                   // List allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},                   // List allowed headers
		AllowCredentials: true,                                                                  // Allow credentials (cookies, authorization headers, etc.)
	}

	r.Use(cors.New(corsConfig))
	r.Use(middlewares.CORSMiddleware())

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public routes
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	// Protected routes
	protected := r.Group("/api")

	protected.Use(cors.New(corsConfig))
	protected.Use(middlewares.CORSMiddleware())
	protected.Use(middlewares.JWTAuthMiddleware())

	protected.GET("/users", controllers.GetUsers)
	protected.GET("/users/:id", controllers.GetUser)
	protected.POST("/users", controllers.CreateUser)
	protected.PUT("/users/:id", controllers.UpdateUser)
	protected.DELETE("/users/:id", controllers.DeleteUser)

	// swagger:route GET /api/roles roles GetRoles
	//
	// GetRoles returns all roles.
	//
	// Responses:
	//
	//	200: successResponse
	protected.GET("/roles", controllers.GetRoles)
	protected.GET("/roles/:id", controllers.GetRole)
	protected.POST("/roles", controllers.CreateRole)
	protected.PUT("/roles/:id", controllers.UpdateRole)
	protected.DELETE("/roles/:id", controllers.DeleteRole)

	return r
}
