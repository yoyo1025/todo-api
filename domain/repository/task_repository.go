package commandRepo

import (
	"todo-api/domain/model"
)

type ITaskCommandRepository interface {
	// タスクを新規作成する
	Create(task *model.Task) error
	// タスクの進捗状況を更新する
	Update(taskId int64, task *model.Task) error
}