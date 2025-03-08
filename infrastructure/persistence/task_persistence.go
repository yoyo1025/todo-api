package persistence

import (
	"time"
	"todo-api/domain/model"
	"todo-api/domain/repository"
	"todo-api/infrastructure/record"
	"todo-api/infrastructure/response"
	queryRepo "todo-api/usecase/query/repository"

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

func NewTaskQueryPersistence(db *gorm.DB) queryRepo.ITaskQueryRepository{
	return &TaskQueryPersistence{
		db: db,
	}
}

func (tp *TaskQueryPersistence) FindAllTask(userId int64) ([]*response.Task, error) {
	var tasks []*response.Task
	result := tp.db.Table("task_records").Where("user_id = ?", userId).Scan(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (tp *TaskQueryPersistence) FindTaskById(taskId int64) (*response.Task, error){
	var task *response.Task
	result := tp.db.Table("task_records").Where("id = ?", taskId).Scan(&task)
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