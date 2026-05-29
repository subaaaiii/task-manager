package controllers

import (
	"backend/database"
	"backend/models"
	requests "backend/structs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {

	var tasks []models.Task

	if err := database.DB.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed get tasks",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success get tasks",
		"data":    tasks,
	})
}

func GetTaskById(c *gin.Context) {

	id := c.Param("id")

	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Task not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success get task",
		"data":    task,
	})
}

func CreateTask(c *gin.Context) {

	var request requests.TaskRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Failed create task:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	task := models.Task{
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
		DueDate:     request.DueDate,
	}

	if err := database.DB.Create(&task).Error; err != nil {
		log.Println("Failed create task:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed create task",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success create task",
		"data":    task,
	})
}

func UpdateTask(c *gin.Context) {

	id := c.Param("id")

	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Task not found",
		})
		return
	}

	var request requests.TaskRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	task.Title = request.Title
	task.Description = request.Description
	task.Status = request.Status
	task.DueDate = request.DueDate

	if err := database.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed update task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success update task",
		"data":    task,
	})
}

func DeleteTask(c *gin.Context) {

	id := c.Param("id")

	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Task not found",
		})
		return
	}

	if err := database.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed delete task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete task",
	})
}
