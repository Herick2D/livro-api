package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"livroApi/models"
	"log"
	"os"
)

var DB *gorm.DB

func ConexaoDB() {

	data := struct {
		Host     string
		User     string
		Password string
		DbName   string
	}{
		Host:     os.Getenv("HOST"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		DbName:   os.Getenv("DBNAME"),
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Sao_Paulo",
		data.Host, data.User, data.Password, data.DbName)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB.AutoMigrate(&models.Livro{})
}
