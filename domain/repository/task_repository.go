package repository

import "todo-api/domain/model"

type ITaskRepository interface {
	// 全てのタスク取得する
	FindAll(userId int64) ([]*model.Task, error)
	// タスクIDをもとにタスクを取得する
	FindById(taskId int64) (*model.Task, error)
	// タスクを新規作成する
	Create(task *model.Task) error
	// タスクの進捗状況を更新する
	Update(taskId int64, task *model.Task) error
}