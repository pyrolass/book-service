package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pyrolass/book-service/db"
	"github.com/pyrolass/book-service/models"
)

func CreateBook(c *gin.Context) {

	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database := db.GetDB()

	if result := database.Create(&book); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"data": book})
}

func GetBookById(c *gin.Context) {

	println("test")
	id := c.Param("id")

	var book models.Book

	database := db.GetDB()

	if result := database.First(&book, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"data": book})
}
