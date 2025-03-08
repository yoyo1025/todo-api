package queryRepo

import (
	"todo-api/infrastructure/response"
)


type ITaskQueryRepository interface {
	// 全てのタスクを取得する
	FindAllTask(userId int64) ([]*response.Task, error)
	// タスクIDで指定して取得
	FindTaskById(taskId int64) (*response.Task, error)
}