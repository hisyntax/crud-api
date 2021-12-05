package main

import (
	"log"

	"github.com/hisyntax/crud-api/routes"
	"github.com/joho/godotenv"
)

//func init runs before the main function
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	//call the router package to run
	routes.UserRoutes()
}
