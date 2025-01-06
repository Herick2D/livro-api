package controllers

import (
	"livroApi/config"
	"livroApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateLivro(c *gin.Context) {
	var livro models.Livro
	if err := c.ShouldBindJSON(&livro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&livro)
	c.JSON(http.StatusCreated, livro)
}

func GetLivros(c *gin.Context) {
	var livros []models.Livro
	config.DB.Find(&livros)
	c.JSON(http.StatusOK, livros)
}

func GetLivro(c *gin.Context) {
	id := c.Param("id")
	var livro models.Livro
	if err := config.DB.First(&livro, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}
	c.JSON(http.StatusOK, livro)
}
