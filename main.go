package main

import (
	"real-time-chat/internal/controllers"
	"real-time-chat/internal/server/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	chatRoom := controllers.NewChatRoom()
	go chatRoom.Run()

	r := gin.Default()
	// r.POST("/register", controllers.Register)
	// r.POST("/login", controllers.Login)
	// r.POST("/chatrooms", controllers.CreateChatroom)
	// r.GET("/chatrooms", controllers.GetChatrooms)
	// r.GET("/chatrooms/:id", controllers.GetMessages)

	// protected := r.Group("/api")
	// protected.Use(middleware.AuthMiddleware())
	// {
	//     protected.GET("/profile", controllers.Profile)  // Example of a protected route
	// }

	r.GET("/ws", func(c *gin.Context) {
		chatRoom.HandleConnection(c.Writer, c.Request)
	})

	r.Run(":8080")
}
