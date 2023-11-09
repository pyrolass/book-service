package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pyrolass/book-service/db"
	"github.com/pyrolass/book-service/routes"
)

func main() {

	server := gin.Default()

	routes.InitializeRoutes(server)

	db.Connect()

	server.Run(":8081")
}
