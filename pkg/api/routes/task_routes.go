package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karma/karma-backend/pkg/api/handlers"
)

func SetupTaskRoutes(router *gin.RouterGroup) {
	taskGroup := router.Group("/task")
	{
		taskGroup.GET("/get_task", handlers.GetTasks)
		// Add other user-related routes as needed
	}
}