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

var Tasks = make(map[int]Task)
var TaskID = 1
