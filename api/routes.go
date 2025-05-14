package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"userprofile-api/controllers"
)

// SetupRouter configures the API routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	
	// Root handler to help with navigation
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to User Profile API",
			"version": "1.0",
			"endpoints": []string{
				"/api/v1/users - Get all users",
				"/api/v1/users/:id - Get user by ID",
			},
		})
	})
	
	// API version group
	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("", controllers.GetUsers)
			users.GET("/:id", controllers.GetUser)
			users.POST("", controllers.CreateUser)
			users.PUT("/:id", controllers.UpdateUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}
	}
	
	return router
}
