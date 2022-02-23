package main

import (
	"gza/user/controllers"
	"gza/user/models"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userAPI := r.Group("/user")
	{
		// GET: http://localhost:8080/user
		userAPI.GET("/", controllers.GetUsers)
		// GET: http://localhost:8080/user/:id
		userAPI.GET("/:id", controllers.GetUser)
		// POST: http://localhost:8080/user
		userAPI.POST("/", controllers.CreateUser)
		// POST: http://localhost:8080/user/login
		userAPI.POST("/login", controllers.Login)
		// PUT: http://localhost:8080/user:id
		userAPI.PUT("/:id", controllers.UpdateUser)
		// DELETE: http://localhost:8080/user
		userAPI.DELETE("/:id", controllers.DeleteUser)
	}
	r.Use()
	r.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})

	DB = models.ConnectDatabase()
	r.Run(":8080")
}
