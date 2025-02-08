package models

// List 清单结构体
type List struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
//Lists定义
var Lists = make(map[int]List)
//用于生成新清单的唯一 ID
var ListID = 1
