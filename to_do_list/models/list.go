package models

// List 清单结构体
type List struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var Lists = make(map[int]List)
var ListID = 1
