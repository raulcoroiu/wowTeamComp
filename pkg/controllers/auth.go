package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raulcoroiu/wowTeamComp/database"
	"github.com/raulcoroiu/wowTeamComp/pkg/models"
)

func RegisterHandler(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Check if the user already exists
	collection := database.Client.Database("your-database-name").Collection("users") // Replace with your database name
	filter := models.User{Email: user.Email}
	count, _ := collection.CountDocuments(c, filter)
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// Hash the user's password and store in the database
	// You should use a library like bcrypt to securely hash passwords
	// Replace the following line with your password hashing code
	user.Password = "hashed-password"

	_, err = collection.InsertOne(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func LoginHandler(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Check if the user exists in the database
	collection := database.Client.Database("your-database-name").Collection("users") // Replace with your database name
	filter := models.User{Email: user.Email, Password: "hashed-password"}            // Replace with the hashed password
	err = collection.FindOne(c, filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Authenticate user successfully
	// You can generate a JWT token and return it here for further authentication

	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
}
