package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	userId int64
	title string
	detail string
	status int64
}

func NewTask(userId int64, title, detail string, status int64) *Task {
	return &Task{
		userId: userId,
		title: title,
		detail: detail,
		status: status,
	}
}

func (t *Task) Update(title, detail string, status int64) *Task {
	return &Task{
		userId: t.userId,
		title: title,
		detail: detail,
		status: status,
	}
}

func (t *Task) GetId() int64 {
	return int64(t.ID)
}

func (t *Task) GetUserId() int64 {
	return t.userId
}

func (t *Task) GetTitle() string {
	return t.title
}

func (t *Task) GetDetail() string {
	return t.detail
}

func (t *Task) GetStatus() int64 {
	return t.status
}

func (t *Task) SetId(id uint) {
	t.ID = id
}