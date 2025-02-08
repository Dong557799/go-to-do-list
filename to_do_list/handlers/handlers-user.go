package handlers

import (
	"net/http"
	"to_do_list/models"

	"github.com/gin-gonic/gin"
)

// Register 用户注册处理函数
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//用户名是否重复
	if _, exists := models.Users[user.Username]; exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}
	models.Users[user.Username] = user
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Login 用户登录处理函数
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//用户名是否存在以及密码是否匹配
	if storedUser, exists := models.Users[user.Username]; exists && storedUser.Password == user.Password {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	}
}
