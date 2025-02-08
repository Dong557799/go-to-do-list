package routers

import (
	"github.com/gin-gonic/gin"
	"to_do_list/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	listGroup := r.Group("/lists")
	{
		listGroup.POST("/", handlers.CreateList)
		listGroup.GET("/", handlers.GetLists)
		listGroup.PUT("/:listID", handlers.UpdateList)
		listGroup.DELETE("/:listID", handlers.DeleteList)

		taskGroup := listGroup.Group("/:listID/tasks")
		{
			taskGroup.POST("/", handlers.CreateTask)
			taskGroup.GET("/", handlers.GetTasks)
			taskGroup.PUT("/:taskID", handlers.UpdateTask)
			taskGroup.DELETE("/:taskID", handlers.DeleteTask)
			taskGroup.PUT("/:taskID/complete", handlers.MarkTaskCompleted)
		}
	}

	return r
}
