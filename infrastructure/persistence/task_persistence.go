package persistence

import (
	"errors"
	"time"
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


// TaskRecord -> model/task
func toDomain(r TaskRecord) (*model.Task, error) {
	task := model.NewTask(r.UserID, r.Title, r.Detail, r.Status)
	if task == nil {
		return nil, errors.New("タスクの生成に失敗しました")
	}
	return task, nil
}

// model/task -> TaskRecord
func toRecord(t *model.Task) TaskRecord {
	return TaskRecord{
		UserID: t.GetUserId(),
		Title: t.GetTitle(),
		Detail: t.GetDetail(),
		Status: t.GetStatus(),
	}
}


func (tp *TaskPersistence) FindAll(userId int64) ([]*model.Task, error)  {
	var records []TaskRecord
	result := tp.db.Where("user_id = ?", userId).Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}

	// []TaskRecord -> []*model.Task
	tasks := make([]*model.Task, 0, len(records))
	for _, rec := range records {
		dom, err := toDomain(rec)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, dom)
	}
	return tasks, nil
}

func (tp *TaskPersistence) FindById(taskId int64) (*model.Task, error) {
	var record TaskRecord
	result := tp.db.First(&record, taskId)
	if result.Error != nil {
		return nil, result.Error
	}

	task, err := toDomain(record)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (tp *TaskPersistence) Create(task *model.Task) error {
	rec := toRecord(task)
	return tp.db.Create(&rec).Error
}

func (tp *TaskPersistence) Update(taskId int64, task *model.Task) error {
	rec := toRecord(task)
	return tp.db.Model(&TaskRecord{}).Where("id = ?", uint(taskId)).Updates(map[string]interface{}{
		"updated_at": time.Now(),
		"title": rec.Title,
		"detail": rec.Detail,
		"status": rec.Status,
	}).Error
}