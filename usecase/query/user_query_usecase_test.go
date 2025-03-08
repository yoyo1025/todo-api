package query

import (
	"testing"
	"todo-api/infrastructure/response"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockQueryUserRepository struct {
	mock.Mock
}

func (m *mockQueryUserRepository) FindByEmail(email string) (*response.User, error) {
	args := m.Called(email)
	user, _ := args.Get(0).(*response.User)
	return user, args.Error(1)
}

func TestFindByEmail_Success(t *testing.T)  {
	mockRepo := new(mockQueryUserRepository)
	dummyuser := &response.User{
		ID: 1,
		Name: "test user",
		Email: "test@gmail.com",
	}
	mockRepo.On("FindByEmail", "test@gmail.com").Return(dummyuser, nil)

	queryUserUsecase := NewQueryUserUsecase(mockRepo)

	e := echo.New()
	c := e.NewContext(nil, nil)

	user, err := queryUserUsecase.FindByEmail(c, "test@gmail.com")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), user.ID)
	assert.Equal(t, "test user", user.Name)
	assert.Equal(t, "test@gmail.com", user.Email)
}