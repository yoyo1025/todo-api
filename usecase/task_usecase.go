package usecase

import (
	"net/http"
	"todo-api/domain/model"
	"todo-api/domain/repository"

	"github.com/labstack/echo/v4"
)

type ITaskUsecase interface {
	FindAllTask(c echo.Context, userId int64) ([]*model.Task, error)
	CreateTask(c echo.Context, userId int64, title, detail string, status int64) error
	UpdateTask(c echo.Context, taskId, userId int64, title, detail string, status int64) error
}

type TaskUsecase struct {
	taskRepository repository.ITaskRepository	
}

func NewTaskUsecase(taskRepository repository.ITaskRepository) ITaskUsecase{
	return &TaskUsecase {
		taskRepository: taskRepository,
	}
}

// 指定したユーザのタスクをすべて取得する
func (tu *TaskUsecase) FindAllTask(c echo.Context, userId int64) ([]*model.Task, error) {
	tasks, err := tu.taskRepository.FindAll(userId)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "タスクの取得に失敗")
	}
	return tasks, nil
}

// タスクを新規作成する
func (tu *TaskUsecase) CreateTask(c echo.Context, userId int64, title, detail string, status int64) error {
	task := model.NewTask(userId, title, detail, status)
	if err := tu.taskRepository.Create(task); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "タスクの作成に失敗")
	}
	return nil
}

// タスク情報を更新する
func (tu *TaskUsecase) UpdateTask(c echo.Context, taskId, userId int64, title, detail string, status int64) error {
	exsistingTask, err := tu.taskRepository.FindById(taskId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "対称のタスクが見つかりません")
	}

	updatedTask := exsistingTask.Update(title, detail, status)

	if err := tu.taskRepository.Update(taskId, updatedTask); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "タスクの更新に失敗")
	}
	return nil
}