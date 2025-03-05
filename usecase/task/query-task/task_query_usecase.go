package querytask

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ITaskQueryUsecase interface {
	FindAllTask(c echo.Context, userId int64) ([]*Task, error)
}

type TaskQueryUsecase struct {
	taskQueryRepository ITaskQueryRepository
}

func NewTaskQueryUsecase(TaskQueryRepository ITaskQueryRepository) ITaskQueryUsecase {
	return &TaskQueryUsecase {
		taskQueryRepository: TaskQueryRepository,
	}
}

func (tqu *TaskQueryUsecase) FindAllTask(c echo.Context, userId int64) ([]*Task, error){
	tasks, err := tqu.taskQueryRepository.FindAll(userId)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "タスクの取得に失敗")
	}
	return tasks, nil
}