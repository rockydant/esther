package main

import (
	"esther/config"
	_ "esther/docs"
	"esther/models"
	"esther/routes"
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

	// Setup the routes
	r := routes.SetupRouter(validate)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start")
	}
}
