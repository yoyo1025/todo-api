package querytask


type ITaskQueryRepository interface {
	// 全てのタスクを取得する
	FindAll(userId int64) ([]*Task, error)
	// タスクIDで指定して取得
	FindById(taskId int64) (*Task, error)
}