package handlers

import (
	"net/http"//提供了 HTTP 客户端和服务器的实现
	"strconv"
	"to_do_list/models"

	"github.com/gin-gonic/gin"
)

// CreateList 创建清单处理函数
func CreateList(c *gin.Context) {
	var list models.List
	//使用 c.ShouldBindJSON 方法将请求体中的 JSON 数据解析到 user 变量并提醒错误
	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//为新创建的清单分配一个唯一的 ID
	list.ID = models.ListID
	models.Lists[models.ListID] = list
	//为下一个新创建的清单分配一个不同的 ID
	models.ListID++
	//返回HTTP 响应
	c.JSON(http.StatusOK, list)
}

// GetLists 获取清单列表处理函数
func GetLists(c *gin.Context) {
	//声明models.List类型切片
	var listArray []models.List
	//遍历清单并将其添加到 listArray 切片
	for _, list := range models.Lists {
		listArray = append(listArray, list)
	}
	//返回HTTP响应
	c.JSON(http.StatusOK, listArray)
}

// UpdateList 更新清单处理函数
func UpdateList(c *gin.Context) {
	// 修改为 :listID
	//获取名为 listID 的参数值
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
	
	//存储从请求体中解析出来的更新后的清单数据
	var updatedList models.List
	if err := c.ShouldBindJSON(&updatedList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedList.ID = listID
	//覆盖原来的清单
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
