package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"livroApi/models"
	"log"
)

var DB *gorm.DB

func ConexaoDB() {
	dsn := "host=postgres dbname=livroapidb user=herick password=123321 port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB.AutoMigrate(&models.Livro{})
}
