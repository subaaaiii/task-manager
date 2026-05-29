package routes

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "PATCH", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	api := router.Group("/api")

	api.Use(middlewares.SimpleAuth())

	{
		api.GET("/task", controllers.GetTasks)
		api.GET("/task/:id", controllers.GetTaskById)
		api.POST("/task", controllers.CreateTask)
		api.DELETE("/task/:id", controllers.DeleteTask)
		api.PATCH("/task/:id", controllers.UpdateTask)
	}

	return router
}
