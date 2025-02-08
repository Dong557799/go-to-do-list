package models

// Task 任务结构体
type Task struct {
	ID        int    `json:"id"`
	ListID    int    `json:"listID"`
	Title     string `json:"title"`
	Priority  string `json:"priority"`
	DueDate   string `json:"dueDate"`
	Completed bool   `json:"completed"`
}
//Tasks定义
var Tasks = make(map[int]Task)
//用于生成新任务的唯一 ID
var TaskID = 1
