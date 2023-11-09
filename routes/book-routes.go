package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pyrolass/book-service/controllers"
)

func BookRoutes(router *gin.Engine) {
	bookGroup := router.Group("/")
	{
		bookGroup.POST("/create_book", controllers.CreateBook)
	}
}
