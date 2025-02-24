package persistence

import (
	"todo-api/domain/model"
	"todo-api/domain/repository"

	"gorm.io/gorm"
)

type TaskPersistence struct {
	db *gorm.DB
}

func NewTaskPersistence(db *gorm.DB) repository.ITaskRepository {
	return &TaskPersistence{
		db: db,
	}
}


func (tp *TaskPersistence) FindAll(userId int64) ([]*model.Task, error)  {
	var tasks []*model.Task
	result := tp.db.Find(&tasks, userId)
	return tasks, result.Error
}

func (tp *TaskPersistence) Create(task *model.Task) error {
	return tp.db.Create(task).Error
}

func (tp *TaskPersistence) Update(task *model.Task) error {
	return tp.db.Save(task).Error
}