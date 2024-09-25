package controllers

import (
	"go-jwt/initializers"
	"go-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body: " + err.Error(),
		})
		return
	}

	if body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email cannot be empty",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user := models.User{
		Email:    body.Email,
		Password: string(hash),
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}
