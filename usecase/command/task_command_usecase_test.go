package command

import (
	"errors"
	"net/http"
	"testing"
	"todo-api/domain/model"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ----------------------------------------
// モックリポジトリの定義
// ----------------------------------------
type mockTaskCommandRepository struct {
	mock.Mock
}

func (m *mockTaskCommandRepository) Create(task *model.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *mockTaskCommandRepository) Update(taskId int64, task *model.Task) error {
	args := m.Called(taskId, task)
	return args.Error(0)
}

// ----------------------------------------
// テスト
// ----------------------------------------
func TestCreateTask_Success(t *testing.T) {
	// モック設定
	mockRepo := new(mockTaskCommandRepository)
	mockRepo.On("Create", mock.Anything).Return(nil)

	usecase := NewTaskCommandUsecase(mockRepo)

	// テスト用 Echo Context
	e := echo.New()
	c := e.NewContext(nil, nil)

	err := usecase.CreateTask(c, 1, "title", "detail", 0)
	assert.NoError(t, err) // エラーが返らないことを期待
	mockRepo.AssertCalled(t, "Create", mock.Anything)
}

func TestCreateTask_Failure(t *testing.T) {
	mockRepo := new(mockTaskCommandRepository)
	mockRepo.On("Create", mock.Anything).Return(errors.New("some db error"))

	usecase := NewTaskCommandUsecase(mockRepo)

	e := echo.New()
	c := e.NewContext(nil, nil)

	err := usecase.CreateTask(c, 1, "title", "detail", 0)
	if assert.Error(t, err) {
		httpErr, ok := err.(*echo.HTTPError)
		assert.True(t, ok, "error should be an HTTPError")
		assert.Equal(t, http.StatusInternalServerError, httpErr.Code)
		assert.Equal(t, "タスクの作成に失敗", httpErr.Message)
	}
}

func TestUpdateTask_Success(t *testing.T) {
	mockRepo := new(mockTaskCommandRepository)
	mockRepo.On("Update", int64(10), mock.Anything).Return(nil)

	usecase := NewTaskCommandUsecase(mockRepo)

	e := echo.New()
	c := e.NewContext(nil, nil)

	err := usecase.UpdateTask(c, 10, 1, "new title", "new detail", 1)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Update", int64(10), mock.Anything)
}

func TestUpdateTask_Failure(t *testing.T) {
	mockRepo := new(mockTaskCommandRepository)
	mockRepo.On("Update", int64(10), mock.Anything).Return(errors.New("some db error"))

	usecase := NewTaskCommandUsecase(mockRepo)

	e := echo.New()
	c := e.NewContext(nil, nil)

	err := usecase.UpdateTask(c, 10, 1, "new title", "new detail", 1)
	if assert.Error(t, err) {
		httpErr, ok := err.(*echo.HTTPError)
		assert.True(t, ok, "error should be an HTTPError")
		assert.Equal(t, http.StatusInternalServerError, httpErr.Code)
		assert.Equal(t, "タスクの更新に失敗", httpErr.Message)
	}
}
