package main

import (
	"log"
	"real-time-chat/internal/controllers"
	"real-time-chat/internal/database"
	"real-time-chat/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.Init()
	chatRoom := controllers.NewChatRoom()
	go chatRoom.Run()

	r := gin.Default()
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	// r.POST("/chatrooms", controllers.CreateChatroom)
	// r.GET("/chatrooms", controllers.GetChatrooms)
	// r.GET("/chatrooms/:id", controllers.GetMessages)

	r.GET("/ws", middleware.AuthMiddleware(), func(c *gin.Context) {
		chatRoom.HandleConnection(c.Writer, c.Request)
	})

	r.Run(":8080")
}
