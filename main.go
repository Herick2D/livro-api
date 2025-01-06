package main

import (
	"livroApi/config"
	"livroApi/models"
	"livroApi/routes"
)

func main() {
	config.ConexaoDB()

	config.DB.AutoMigrate(&models.Livro{})

	router := routes.SetupRouter()

	router.Run(":8080")
}
