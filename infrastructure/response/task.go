package response

type Task struct {
	TaskID int64 `json:"taskId" gorm:"column:id"`
	UserID int64 `json:"userId" gorm:"column:user_id"`
	Title string `json:"title" gorm:"column:title"`
	Detail string `json:"detail" gorm:"column:detail"`
	Status uint8 `json:"status" gorm:"column:status"` 
}