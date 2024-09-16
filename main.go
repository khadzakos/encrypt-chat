package main

import (
	"real-time-chat/internal/controllers"
	"real-time-chat/internal/server/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// protected := r.Group("/api")
	// protected.Use(middleware.AuthMiddleware())
	// {
	//     protected.GET("/profile", controllers.Profile)  // Example of a protected route
	// }

	r.Run()
}
