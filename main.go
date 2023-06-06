package main

import (
	"net/http"

	"tasks/controllers"
	"tasks/models"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "Main web"})
	})

	r.GET("/api/tasks", controllers.FindTasks)
	r.POST("/api/tasks", controllers.CreateTask)
	r.GET("/api/tasks/:id", controllers.FindTask)
	r.PATCH("/api/tasks/:id", controllers.UpdateTask)
	r.DELETE("/api/tasks/:id", controllers.DeleteTask)

	return r
}

func main() {
	r := setupRouter()

	models.ConnectDatabase() // new
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
