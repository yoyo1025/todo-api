package query

// import (
// 	"errors"
// 	"net/http"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // ----------------------------------------
// // モックリポジトリの定義
// // ----------------------------------------
// type mockTaskQueryRepository struct {
// 	mock.Mock
// }

// func (m *mockTaskQueryRepository) FindAll(userId int64) ([]*Task, error) {
// 	args := m.Called(userId)
// 	tasks, _ := args.Get(0).([]*Task)
// 	return tasks, args.Error(1)
// }

// func (m *mockTaskQueryRepository) FindById(taskId int64) (*Task, error) {
// 	// 今回のインタフェーステストでは未使用
// 	return nil, nil
// }

// // ----------------------------------------
// // テスト
// // ----------------------------------------
// func TestFindAllTask_Success(t *testing.T) {
// 	mockRepo := new(mockTaskQueryRepository)
// 	dummyTasks := []*Task{
// 		{TaskID: 1, UserID: 1, Title: "Task 1"},
// 		{TaskID: 2, UserID: 1, Title: "Task 2"},
// 	}
// 	// モックが呼ばれたらダミータスクと nil エラーを返す
// 	mockRepo.On("FindAll", int64(1)).Return(dummyTasks, nil)

// 	usecase := NewTaskQueryUsecase(mockRepo)

// 	e := echo.New()
// 	c := e.NewContext(nil, nil)

// 	tasks, err := usecase.FindAllTask(c, 1)
// 	assert.NoError(t, err)
// 	assert.Len(t, tasks, 2)
// 	assert.Equal(t, int64(1), tasks[0].TaskID)
// 	assert.Equal(t, "Task 1", tasks[0].Title)
// }

// func TestFindAllTask_Failure(t *testing.T) {
// 	mockRepo := new(mockTaskQueryRepository)
// 	mockRepo.On("FindAll", int64(2)).Return(nil, errors.New("db error"))

// 	usecase := NewTaskQueryUsecase(mockRepo)

// 	e := echo.New()
// 	c := e.NewContext(nil, nil)

// 	tasks, err := usecase.FindAllTask(c, 2)
// 	assert.Nil(t, tasks)
// 	if assert.Error(t, err) {
// 		httpErr, ok := err.(*echo.HTTPError)
// 		assert.True(t, ok)
// 		assert.Equal(t, http.StatusInternalServerError, httpErr.Code)
// 		assert.Equal(t, "タスクの取得に失敗", httpErr.Message)
// 	}
// }
