package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shohinsan/gingonic/db"
	"github.com/shohinsan/gingonic/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
