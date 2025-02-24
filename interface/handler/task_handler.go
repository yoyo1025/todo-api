package handler

import (
	"net/http"
	"strconv"
	"todo-api/domain/model"
	"todo-api/interface/dto"
	"todo-api/usecase"

	"github.com/labstack/echo/v4"
)

type ITaskHandler interface {
	HandleGetAllTasks(c echo.Context) error
	HandleCreateTask(c echo.Context) error
	HandleUpdateTask(c echo.Context) error
}

type TaskHandler struct {
	taskUsecase usecase.ITaskUsecase
}

func NewTaskHandler(taskUsecase usecase.ITaskUsecase) ITaskHandler {
	return &TaskHandler {
		taskUsecase: taskUsecase,
	}
}

func (th TaskHandler) HandleGetAllTasks(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return err
	}
	tasks, err := th.taskUsecase.FindAllTask(c, int64(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "データ取得に失敗")
	}
	return c.JSON(http.StatusOK, dto.ToTaskResponses(tasks))
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
	err = th.taskUsecase.CreateTask(c, int64(userId), task.Title, task.Detail, task.Status)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, "タスクを作成しました")
}

func (th TaskHandler) HandleUpdateTask(c echo.Context) error {
	var task model.Task
	if err := c.Bind(&task); err != nil {
		return err
	}
	return th.taskUsecase.UpdateTask(c, task.GetId(), task.GetUserId(), task.GetTitle(), task.GetDetail(), task.GetStatus())
}