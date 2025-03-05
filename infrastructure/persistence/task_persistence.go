package persistence

import (
	"time"
	"todo-api/domain/model"
	"todo-api/domain/repository"
	"todo-api/infrastructure/record"
	querytask "todo-api/usecase/task/query-task"

	"gorm.io/gorm"
)

type TaskCommandPersistence struct {
	db *gorm.DB
}

type TaskQueryPersistence struct {
	db *gorm.DB
}

func NewTaskCommandPersistence(db *gorm.DB) repository.ITaskCommandRepository {
	return &TaskCommandPersistence{
		db: db,
	}
}

func NewTaskQueryPersistence(db *gorm.DB) querytask.ITaskQueryRepository {
	return &TaskQueryPersistence{
		db: db,
	}
}

func (tp *TaskQueryPersistence) FindAll(userId int64) ([]*querytask.Task, error)  {
	var tasks []*querytask.Task
	result := tp.db.Table("task_records").Where("user_id = ?", userId).Scan(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (tp *TaskQueryPersistence) FindById(taskId int64) (*querytask.Task, error) {
	var task *querytask.Task
	result := tp.db.Table("task_records").Where("task_id = ?", taskId).Scan(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (tp *TaskCommandPersistence) Create(task *model.Task) error {
	rec := record.TaskRecord{
		UserID: task.GetUserId(),
		Title: task.GetTitle(),
		Detail: task.GetDetail(),
		Status: task.GetStatus(),
	}
	return tp.db.Create(&rec).Error
}

func (tp *TaskCommandPersistence) Update(taskId int64, task *model.Task) error {
	rec := record.TaskRecord{
		Title: task.GetTitle(),
		Detail: task.GetDetail(),
		Status: task.GetStatus(),
	}
	return tp.db.Model(&record.TaskRecord{}).Where("id = ?", uint(taskId)).Updates(map[string]interface{}{
		"updated_at": time.Now(),
		"title": rec.Title,
		"detail": rec.Detail,
		"status": rec.Status,
	}).Error
}