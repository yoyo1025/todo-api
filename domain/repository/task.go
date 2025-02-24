package repository

import "todo-api/domain/model"

type ITaskRepository interface {
	FindAll(userId int64) ([]*model.Task, error)
	Create(task *model.Task) error
	Update(task *model.Task) error
}