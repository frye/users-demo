package api

import (
	"path/filepath"
	"runtime"
	"github.com/gin-gonic/gin"
	"userprofile-api/controllers"
)

// SetupRouter configures the API routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	
	// Get the absolute path to the templates directory
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(filepath.Dir(b))
	templatesPath := filepath.Join(basePath, "templates/*")
	
	// Setup template rendering
	router.LoadHTMLGlob(templatesPath)
	
	// Root handler shows a nice HTML table of all users
	router.GET("/", controllers.HomePageHandler)
	
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
