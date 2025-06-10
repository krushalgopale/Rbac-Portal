package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/krushalgopale/internal/database"
	"github.com/krushalgopale/internal/routes"
)

func main() {
	loadEnv()
	loadDatabse()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}

func loadDatabse() {
	database.ConnDB()
}

func serveApplication(){
	r := gin.Default()
	routes.Routes(r)
	r.Run(":8080")
	fmt.Println("Server running on port 8080")
}
