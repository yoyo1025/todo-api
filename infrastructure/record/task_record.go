package record

import "gorm.io/gorm"

type TaskRecord struct {
	gorm.Model
	UserID int64 `gorm:"column:user_id"`
	Title  string `gorm:"column:title"`
	Detail string `gorm:"column:detail"`
	Status int64 `gorm:"column:status"`
}