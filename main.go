package main

import (
	"tasks/controllers"
	"tasks/models"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"content-type"},
		AllowMethods:    []string{"POST", "DELETE", "GET", "PATCH"},
	}))
	r.GET("/tasks", controllers.FindTasks)
	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks/:id", controllers.FindTask)
	r.PATCH("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)

	return r
}

func main() {
	// Load the .env file in the current directory
	godotenv.Load()

	r := setupRouter()

	models.ConnectDatabase() // new
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8082")
}
