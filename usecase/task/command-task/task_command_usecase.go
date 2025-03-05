package commandtask

import (
	"fmt"
	"net/http"
	"todo-api/domain/model"
	"todo-api/domain/repository"

	"github.com/labstack/echo/v4"
)


type ITaskCommandUsecase interface {
	// タスクを新規作成する
	CreateTask(c echo.Context, userId int64, title, detail string, status int64) error
	// タスク情報を更新する(タスクの削除もこのユーズケースを用いる status=2 にする)
	UpdateTask(c echo.Context, taskId, userId int64, title, detail string, status int64) error
}

type TaskCommandUsecase struct {
	taskRepository repository.ITaskCommandRepository
}

func NewTaskCommandUsecase(taskRepository repository.ITaskCommandRepository) ITaskCommandUsecase{
	return &TaskCommandUsecase {
		taskRepository: taskRepository,
	}
}

func (tu *TaskCommandUsecase) CreateTask(c echo.Context, userId int64, title, detail string, status int64) error {
	task := model.NewTask(userId, title, detail, status)
	fmt.Println(task)
	if err := tu.taskRepository.Create(task); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "タスクの作成に失敗")
	}
	return nil
}

// タスク情報を更新する(タスクの削除もこのユーズケースを用いる status=2 にする)
func (tu *TaskCommandUsecase) UpdateTask(c echo.Context, taskId, userId int64, title, detail string, status int64) error {
	updatedTask := model.NewTask(userId, title, detail, status)
	if err := tu.taskRepository.Update(taskId, updatedTask); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "タスクの更新に失敗")
	}
	return nil
}