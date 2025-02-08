package models

// User 用户结构体
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
//Users定义
var Users = make(map[string]User)
