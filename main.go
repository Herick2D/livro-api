package main

import (
	"livroApi/config"
	"livroApi/models"
	"livroApi/routes"

	"github.com/gin-contrib/cors"
)

func main() {

	config.ConexaoDB()

	config.DB.AutoMigrate(&models.Livro{})

	router := routes.SetupRouter()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.Run(":8080")
}