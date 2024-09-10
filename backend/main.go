package main

import (
	"esther/config"
	_ "esther/docs"
	"esther/models"
	"esther/routes"
	"log"
)

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
