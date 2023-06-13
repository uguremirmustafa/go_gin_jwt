package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uguremirmustafa/go-gin-jwt/controllers"
	"github.com/uguremirmustafa/go-gin-jwt/initializers"
	"github.com/uguremirmustafa/go-gin-jwt/middlewares"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signup", controllers.Singup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)

	r.Run()
}
