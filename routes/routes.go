package routes

import (
	"github.com/gin-gonic/gin"
	"livroApi/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/livros", controllers.CreateLivro)
		api.GET("/livros", controllers.GetLivros)
	}

	return router
}
