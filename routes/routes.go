package routes

import (
	"vse-bank/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	protected := api.Group("/")
	protected.Use(controllers.AuthMiddleware())

	protected.GET("/banks", controllers.GetBanks)
	protected.POST("/banks", controllers.AddBank)
	protected.PUT("/banks/:id", controllers.UpdateBank)
	protected.DELETE("/banks/:id", controllers.DeleteBank)
	protected.GET("/users", controllers.GetUsers)
	protected.POST("/users", controllers.AddUser)
	protected.PUT("/users/:id", controllers.UpdateUser)
	protected.DELETE("/users/:id", controllers.DeleteUser)
	protected.GET("/loans", controllers.GetLoans)
	protected.POST("/loans", controllers.AddLoan)
	protected.PUT("/loans/:id", controllers.UpdateLoan)
	protected.DELETE("/loans/:id", controllers.DeleteLoan)

}
