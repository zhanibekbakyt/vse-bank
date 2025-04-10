package main

import (
	"vse-bank/config"
	"vse-bank/controllers"
	"vse-bank/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to DB and set to controller
	config.ConnectDatabase()
	controllers.SetDB(config.DB)

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
