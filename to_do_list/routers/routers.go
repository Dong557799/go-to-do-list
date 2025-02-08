package routers

import (
	"github.com/gin-gonic/gin"
	"to_do_list/handlers"
)

//定义SetupRouter 的函数，用于设置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	//注册 handlers.Register handlers.Login 处理函数
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	//处理与清单相关的路由
	listGroup := r.Group("/lists")
	{
		listGroup.POST("/", handlers.CreateList)
		listGroup.GET("/", handlers.GetLists)
		listGroup.PUT("/:listID", handlers.UpdateList)
		listGroup.DELETE("/:listID", handlers.DeleteList)

		//处理与任务相关的路由
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
