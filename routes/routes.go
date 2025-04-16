package routes

import (
	"vse-bank/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Public routes
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	// Protected routes
	protected := api.Group("/")
	protected.Use(controllers.AuthMiddleware())

	protected.GET("/banks", controllers.GetBanks)
	protected.POST("/banks", controllers.AddBank)
	protected.GET("/users", controllers.GetUsers)
	protected.POST("/users", controllers.AddUser)
	protected.GET("/loans", controllers.GetLoans)
	protected.POST("/loans", controllers.AddLoan)
	protected.PUT("/loans/:id", controllers.UpdateLoan)
	protected.DELETE("/loans/:id", controllers.DeleteLoan)

}
