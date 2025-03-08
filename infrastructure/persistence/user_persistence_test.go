package persistence

import (
	"testing"
	"todo-api/database"
	"todo-api/domain/model"

	"github.com/stretchr/testify/assert"
)

func TestUserPersistence_SignUp(t *testing.T) {
	db := database.SetupTestDB()

	userCmdRepo := NewCommandUserPersistence(db)

	t.Run("SignUp Success", func(t *testing.T) {
		newUser := model.NewUser("Test User", "test@example.com")
		createdUser, err := userCmdRepo.SignUp(newUser)

		assert.NoError(t, err, "ユーザー作成に失敗してはいけない")
		assert.NotNil(t, createdUser, "作成されたユーザーがnilであってはいけない")
		assert.Equal(t, "Test User", createdUser.GetName())
		assert.Equal(t, "test@example.com", createdUser.GetEmail())
		assert.NotZero(t, createdUser.ID)
	})
}

func TestUserPersistence_FindByEmail(t *testing.T) {
	db := database.SetupTestDB()

	userCmdRepo := NewCommandUserPersistence(db)
	userQueryRepo := NewQueryUserPersistence(db)

	// データ準備
	_, _ = userCmdRepo.SignUp(model.NewUser("Find Test", "find@test.com"))

	t.Run("FindByEmail Success", func(t *testing.T) {
		user, err := userQueryRepo.FindByEmail("find@test.com")
		assert.NoError(t, err, "検索が失敗してはいけない")
		assert.NotNil(t, user, "ユーザーは存在するはず")
		assert.Equal(t, "Find Test", user.GetName())
		assert.Equal(t, "find@test.com", user.GetEmail())
	})
}