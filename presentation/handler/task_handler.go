package handler

import (
	"net/http"
	"strconv"
	"todo-api/presentation/dto"
	"todo-api/usecase/command"
	"todo-api/usecase/query"

	"github.com/labstack/echo/v4"
)

type ITaskHandler interface {
	HandleGetAllTasks(c echo.Context) error
	HandleCreateTask(c echo.Context) error
	HandleUpdateTask(c echo.Context) error
}

type TaskHandler struct {
	taskCommandUsecase command.ITaskCommandUsecase
	taskQueryUsecase query.ITaskQueryUsecase
}

func NewTaskHandler(taskCommandUsecase command.ITaskCommandUsecase, taskQueryUsecase query.ITaskQueryUsecase) ITaskHandler {
	return &TaskHandler {
		taskCommandUsecase: taskCommandUsecase,
		taskQueryUsecase: taskQueryUsecase,
	}
}

func (th TaskHandler) HandleGetAllTasks(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return err
	}
	tasks, err := th.taskQueryUsecase.FindAllTask(c, int64(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "データ取得に失敗")
	}
	return c.JSON(http.StatusOK, tasks)
}

func (th TaskHandler) HandleCreateTask(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return err
	}
	var task dto.TaskRequest
	if err := c.Bind(&task); err != nil {
		return err
	}
	err = th.taskCommandUsecase.CreateTask(c, int64(userId), task.Title, task.Detail, task.Status)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, "タスクを作成しました")
}

func (th TaskHandler) HandleUpdateTask(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return err
	}
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		return err
	}
	var task dto.TaskRequest
	if err := c.Bind(&task); err != nil {
		return err
	}
	err = th.taskCommandUsecase.UpdateTask(c, int64(taskId), int64(userId), task.Title, task.Detail, task.Status)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, "タスクを更新しました")
}