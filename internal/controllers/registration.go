package controllers

import (
	"net/http"
	"real-time-chat/internal/database"
	"real-time-chat/internal/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.UserRegister

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

	if database.GetDB().Where("username = ?", user.Username).First(&models.User{}).RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	err := user.HashPassword(user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	var newUser models.User
	newUser.Username = user.Username
	newUser.Password = user.Password

	if database.GetDB().Create(&newUser).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}
