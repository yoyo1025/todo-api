package command

import (
	"testing"
	"todo-api/domain/model"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCommandUserRepository struct {
	mock.Mock
}


// SignUp メソッドのモック実装
func (m *mockCommandUserRepository) SignUp(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestSignUp_Success(t *testing.T)  {
	// モック設定
	mockRepo := new(mockCommandUserRepository)
	mockRepo.On("SignUp", mock.Anything).Return(nil)

	commandUserUsecase := NewCommandUserUsecase(mockRepo)

	e := echo.New()
	c := e.NewContext(nil, nil)

	err := commandUserUsecase.SingUp(c, "test", "test@gmail.com")
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "SignUp", mock.Anything)
}

