package handlers

import (
	"net/http"
	"strconv"
	"to_do_list/models"

	"github.com/gin-gonic/gin"
)

// CreateList 创建清单处理函数
func CreateList(c *gin.Context) {
	var list models.List
	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	list.ID = models.ListID
	models.Lists[models.ListID] = list
	models.ListID++
	c.JSON(http.StatusOK, list)
}

// GetLists 获取清单列表处理函数
func GetLists(c *gin.Context) {
	var listArray []models.List
	for _, list := range models.Lists {
		listArray = append(listArray, list)
	}
	c.JSON(http.StatusOK, listArray)
}

// UpdateList 更新清单处理函数
func UpdateList(c *gin.Context) {
	// 修改为 :listID
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
	var updatedList models.List
	if err := c.ShouldBindJSON(&updatedList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedList.ID = listID
	models.Lists[listID] = updatedList
	c.JSON(http.StatusOK, updatedList)
}

// DeleteList 删除清单处理函数
func DeleteList(c *gin.Context) {
	// 修改为 :listID
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
	delete(models.Lists, listID)
	c.JSON(http.StatusOK, gin.H{"message": "List deleted successfully"})
}
