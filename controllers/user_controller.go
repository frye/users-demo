package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"userprofile-api/models"
)

// Sample user data
var users = []models.UserProfile{
	{ID: "1", FullName: "John Doe", Emoji: "😀"},
	{ID: "2", FullName: "Jane Smith", Emoji: "🚀"},
	{ID: "3", FullName: "Robert Johnson", Emoji: "🎸"},
}

// HomePageHandler renders a HTML page displaying users in a table
func HomePageHandler(c *gin.Context) {
	log.Println("GET / endpoint called")
	c.HTML(http.StatusOK, "users.html", gin.H{
		"Users": users,
	})
}

// GetUsers returns all users
func GetUsers(c *gin.Context) {
	log.Println("GET /api/v1/users endpoint called")
	c.JSON(http.StatusOK, users)
}

// GetUser returns a single user by ID
func GetUser(c *gin.Context) {
	id := c.Param("id")
	
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// CreateUser adds a new user
func CreateUser(c *gin.Context) {
	var newUser models.UserProfile
	
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// For simplicity, we're just appending to the slice
	// In a real application, you would use a database
	users = append(users, newUser)
	
	c.JSON(http.StatusCreated, newUser)
}

// UpdateUser updates an existing user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser models.UserProfile
	
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id // Ensure ID doesn't change
			users[i] = updatedUser
			c.JSON(http.StatusOK, updatedUser)
			return
		}
	}
	
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// DeleteUser removes a user by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}
	
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
