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
