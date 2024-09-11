package main

import (
	"esther/config"
	controller "esther/controllers"
	_ "esther/docs"
	"esther/models"
	repository "esther/repositories"
	"esther/routes"
	service "esther/services"
	"log"

	"github.com/go-playground/validator/v10"
)

// @title Esther Server API
// @version 1.0
// @description This is api server for esther project
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey JWT
// @in                          header
// @name                        Authorization

// @host localhost:8080
// @BasePath /
func main() {
	// Initialize the database connection
	config.ConnectDatabase()
	validate := validator.New()

	// Assign the database instance to the models package
	models.InitDB()

	// Migrate the schema
	models.DB.AutoMigrate(&models.User{}, &models.Role{})

	models.CreateDefaultRoleAndAdmin()

	//Init Role Repository
	roleRepository := repository.InitRolesRepositoryImpl(models.DB)

	//Init Role Service
	roleService := service.InitRolesServiceImpl(roleRepository, validate)

	//Init Role controller
	roleController := controller.InitRoleController(roleService)

	//Init Repository
	userRepository := repository.InitUsersRepositoryImpl(models.DB)

	//Init Service
	userService := service.InitUsersServiceImpl(userRepository, validate)

	//Init controller
	userController := controller.InitUserController(userService)

	// Setup the routes
	r := routes.SetupRouter(roleController, userController)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start")
	}
}
