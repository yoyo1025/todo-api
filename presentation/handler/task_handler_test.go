package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"todo-api/presentation/dto"
	"todo-api/presentation/handler"
	querytask "todo-api/usecase/task/query-task"
)

// =========================================
// モックの定義
// =========================================

type mockTaskCommandUsecase struct {
	mock.Mock
}

func (m *mockTaskCommandUsecase) CreateTask(c echo.Context, userId int64, title, detail string, status int64) error {
	args := m.Called(c, userId, title, detail, status)
	return args.Error(0)
}
func (m *mockTaskCommandUsecase) UpdateTask(c echo.Context, taskId, userId int64, title, detail string, status int64) error {
	args := m.Called(c, taskId, userId, title, detail, status)
	return args.Error(0)
}

type mockTaskQueryUsecase struct {
	mock.Mock
}

func (m *mockTaskQueryUsecase) FindAllTask(c echo.Context, userId int64) ([]*querytask.Task, error) {
	args := m.Called(c, userId)

	tasks, _ := args.Get(0).([]*querytask.Task)
	return tasks, args.Error(1)
}

// =========================================
// テスト
// =========================================

func TestHandleGetAllTasks_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/1/tasks", nil)
	// URLパラメータを設定
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("userId")
	c.SetParamValues("1")

	mockQuery := new(mockTaskQueryUsecase)
	// ダミーのタスク
	dummyTasks := []*querytask.Task{
		{
			TaskID: 1,
			UserID: 1,
			Title:  "Task 1",
		},
		{
			TaskID: 2,
			UserID: 1,
			Title:  "Task 2",
		},
	}
	// 正常にタスクが取得できる想定
	mockQuery.On("FindAllTask", c, int64(1)).Return(dummyTasks, nil)

	mockCmd := new(mockTaskCommandUsecase)

	h := handler.NewTaskHandler(mockCmd, mockQuery)

	if assert.NoError(t, h.HandleGetAllTasks(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var result []*querytask.Task
		_ = json.Unmarshal(rec.Body.Bytes(), &result)
		assert.Len(t, result, 2)
		assert.Equal(t, int64(1), result[0].TaskID)
		assert.Equal(t, "Task 1", result[0].Title)
	}
}

func TestHandleGetAllTasks_Failure(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/1/tasks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("userId")
	c.SetParamValues("1")

	mockQuery := new(mockTaskQueryUsecase)
	// エラーを返す
	mockQuery.On("FindAllTask", c, int64(1)).Return(nil, errors.New("some error"))

	mockCmd := new(mockTaskCommandUsecase)

	h := handler.NewTaskHandler(mockCmd, mockQuery)

	err := h.HandleGetAllTasks(c)
	// ハンドラ内で c.JSON(http.StatusInternalServerError, ...) を返している
	// Handler自体はnilを返す可能性もあるが、ここでは明示的にerr確認
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, "\"データ取得に失敗\"\n", rec.Body.String())
}

func TestHandleCreateTask_Success(t *testing.T) {
	e := echo.New()

	// テスト用ボディ(JSON)
	taskBody, _ := json.Marshal(dto.TaskRequest{
		Title:  "New Task",
		Detail: "Some Detail",
		Status: 0,
	})

	req := httptest.NewRequest(http.MethodPost, "/users/1/tasks", bytes.NewReader(taskBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("userId")
	c.SetParamValues("1")

	mockCmd := new(mockTaskCommandUsecase)
	// モックの返却設定
	mockCmd.On("CreateTask", c, int64(1), "New Task", "Some Detail", int64(0)).Return(nil)

	mockQuery := new(mockTaskQueryUsecase)

	h := handler.NewTaskHandler(mockCmd, mockQuery)

	// 実行
	err := h.HandleCreateTask(c)

	// 検証
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, "\"タスクを作成しました\"\n", rec.Body.String())
	mockCmd.AssertCalled(t, "CreateTask", c, int64(1), "New Task", "Some Detail", int64(0))
}

func TestHandleCreateTask_Failure(t *testing.T) {
	e := echo.New()

	taskBody, _ := json.Marshal(dto.TaskRequest{
		Title:  "New Task",
		Detail: "Some Detail",
		Status: 0,
	})

	req := httptest.NewRequest(http.MethodPost, "/users/1/tasks", bytes.NewReader(taskBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("userId")
	c.SetParamValues("1")

	mockCmd := new(mockTaskCommandUsecase)
	// エラーを返す
	mockCmd.On("CreateTask", c, int64(1), "New Task", "Some Detail", int64(0)).Return(errors.New("create failed"))

	mockQuery := new(mockTaskQueryUsecase)

	h := handler.NewTaskHandler(mockCmd, mockQuery)

	err := h.HandleCreateTask(c)

	assert.Error(t, err)
}

func TestHandleUpdateTask_Success(t *testing.T) {
	e := echo.New()

	taskBody, _ := json.Marshal(dto.TaskRequest{
		Title:  "Updated Task",
		Detail: "Updated Detail",
		Status: 1,
	})

	req := httptest.NewRequest(http.MethodPut, "/users/1/tasks/10", bytes.NewReader(taskBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("userId", "taskId")
	c.SetParamValues("1", "10")

	mockCmd := new(mockTaskCommandUsecase)
	// 更新に成功するパターン
	mockCmd.On("UpdateTask", c, int64(10), int64(1), "Updated Task", "Updated Detail", int64(1)).Return(nil)

	mockQuery := new(mockTaskQueryUsecase)

	h := handler.NewTaskHandler(mockCmd, mockQuery)

	err := h.HandleUpdateTask(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, "\"タスクを更新しました\"\n", rec.Body.String())
	mockCmd.AssertCalled(t, "UpdateTask", c, int64(10), int64(1), "Updated Task", "Updated Detail", int64(1))
}

func TestHandleUpdateTask_Failure(t *testing.T) {
	e := echo.New()

	taskBody, _ := json.Marshal(dto.TaskRequest{
		Title:  "Updated Task",
		Detail: "Updated Detail",
		Status: 1,
	})

	req := httptest.NewRequest(http.MethodPut, "/users/1/tasks/10", bytes.NewReader(taskBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("userId", "taskId")
	c.SetParamValues("1", "10")

	mockCmd := new(mockTaskCommandUsecase)
	// エラーを返す
	mockCmd.On("UpdateTask", c, int64(10), int64(1), "Updated Task", "Updated Detail", int64(1)).Return(errors.New("update failed"))

	mockQuery := new(mockTaskQueryUsecase)

	h := handler.NewTaskHandler(mockCmd, mockQuery)

	err := h.HandleUpdateTask(c)

	assert.Error(t, err)
}
