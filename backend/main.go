package main

import (
	"esther/config"
	_ "esther/docs"
	"esther/models"
	"esther/routes"
	"log"
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

	// Assign the database instance to the models package
	models.InitDB()

	// Migrate the schema
	models.DB.AutoMigrate(&models.User{}, &models.Role{})

	// Setup the routes
	r := routes.SetupRouter()

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start")
	}
}
