package query

import (
	"net/http"
	"todo-api/infrastructure/response"
	queryRepo "todo-api/usecase/query/repository"

	"github.com/labstack/echo/v4"
)

type ITaskQueryUsecase interface {
	FindAllTask(c echo.Context, userId int64) ([]*response.Task, error)
}

type TaskQueryUsecase struct {
	taskQueryRepository queryRepo.ITaskQueryRepository
}

func NewTaskQueryUsecase(TaskQueryRepository queryRepo.ITaskQueryRepository) ITaskQueryUsecase {
	return &TaskQueryUsecase {
		taskQueryRepository: TaskQueryRepository,
	}
}

func (tqu *TaskQueryUsecase) FindAllTask(c echo.Context, userId int64) ([]*response.Task, error){
	tasks, err := tqu.taskQueryRepository.FindAllTask(userId)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "タスクの取得に失敗")
	}
	return tasks, nil
}