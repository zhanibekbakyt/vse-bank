package routes

import (
	"vse-bank/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.GET("/banks", controllers.GetBanks)
	api.POST("/banks", controllers.AddBank)
	api.GET("/users", controllers.GetUsers)
	api.POST("/users", controllers.AddUser)
	api.GET("/loans", controllers.GetLoans)
	api.POST("/loans", controllers.AddLoan)
}
