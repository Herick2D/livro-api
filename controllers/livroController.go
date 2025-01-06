package controllers

import (
	"gorm.io/gorm"
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

func UpdateLivro(c *gin.Context) {
	var livro models.Livro
	id := c.Param("id")

	if err := config.DB.First(&livro, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}

	var input models.Livro
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos ou incompletos: " + err.Error()})
		return
	}

	livro.Titulo = input.Titulo
	livro.Categoria = input.Categoria
	livro.Autor = input.Autor
	livro.Sinopse = input.Sinopse

	if result := config.DB.Save(&livro); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao atualizar o livro: " + result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, livro)
}

func DeleteLivro(c *gin.Context) {
	id := c.Param("id")
	var livro models.Livro

	if err := config.DB.First(&livro, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar o livro: " + err.Error()})
		}
		return
	}

	if err := config.DB.Delete(&livro).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir o livro: " + err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
