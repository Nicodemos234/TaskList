package controllers

import (
	"net/http"

	"tasks/models"

	"github.com/gin-gonic/gin"
)

// GET /api/tasks
// Get all tasks
func FindTasks(c *gin.Context) {
	var tasks []models.Task
	models.DB.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// GET /api/tasks/:id
// Find a task
func FindTask(c *gin.Context) { // Get model if exist
	var task models.Task

	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

type CreateTaskInput struct {
	Title string `json:"title" binding:"required"`
}

// POST /api/tasks
// Create new task
func CreateTask(c *gin.Context) {
	// Validate input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create task
	task := models.Task{Title: input.Title}
	models.DB.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

type UpdateTaskInput struct {
	Title string `json:"title"`
}

// PATCH /api/tasks/:id
// Update a task
func UpdateTask(c *gin.Context) {
	// Get model if exist
	var task models.Task
	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&task).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// DELETE /api/tasks/:id
// Delete a task
func DeleteTask(c *gin.Context) {
	// Get model if exist
	var task models.Task
	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
