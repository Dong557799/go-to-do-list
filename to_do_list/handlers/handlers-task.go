package handlers

import (
	"net/http"
	"strconv"
	"to_do_list/models"

	"github.com/gin-gonic/gin"
)

// CreateTask 创建任务处理函数
func CreateTask(c *gin.Context) {
	listIDStr := c.Param("listID")
	listID, err := strconv.Atoi(listIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid list ID"})
		return
	}
	if _, exists := models.Lists[listID]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found"})
		return
	}
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.ID = models.TaskID
	task.ListID = listID
	if task.Priority == "" {
		task.Priority = "p2"
	}
	models.Tasks[models.TaskID] = task
	models.TaskID++
	c.JSON(http.StatusOK, task)
}

// GetTasks 获取任务列表处理函数
func GetTasks(c *gin.Context) {
	listIDStr := c.Param("listID")
	listID, err := strconv.Atoi(listIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid list ID"})
		return
	}
	if _, exists := models.Lists[listID]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found"})
		return
	}
	var taskArray []models.Task
	for _, task := range models.Tasks {
		if task.ListID == listID {
			taskArray = append(taskArray, task)
		}
	}
	c.JSON(http.StatusOK, taskArray)
}

// UpdateTask 更新任务处理函数
func UpdateTask(c *gin.Context) {
	listIDStr := c.Param("listID")
	listID, err := strconv.Atoi(listIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid list ID"})
		return
	}
	if _, exists := models.Lists[listID]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found"})
		return
	}
	taskIDStr := c.Param("taskID")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	if _, exists := models.Tasks[taskID]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTask.ID = taskID
	updatedTask.ListID = listID
	models.Tasks[taskID] = updatedTask
	c.JSON(http.StatusOK, updatedTask)
}

// DeleteTask 删除任务处理函数
func DeleteTask(c *gin.Context) {
	listIDStr := c.Param("listID")
	listID, err := strconv.Atoi(listIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid list ID"})
		return
	}
	if _, exists := models.Lists[listID]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found"})
		return
	}
	taskIDStr := c.Param("taskID")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	if _, exists := models.Tasks[taskID]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	delete(models.Tasks, taskID)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// MarkTaskCompleted 标记任务完成处理函数
func MarkTaskCompleted(c *gin.Context) {
	listIDStr := c.Param("listID")
	listID, err := strconv.Atoi(listIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid list ID"})
		return
	}
	if _, exists := models.Lists[listID]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found"})
		return
	}
	taskIDStr := c.Param("taskID")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	if _, exists := models.Tasks[taskID]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	task := models.Tasks[taskID]
	task.Completed = true
	models.Tasks[taskID] = task
	c.JSON(http.StatusOK, task)
}
