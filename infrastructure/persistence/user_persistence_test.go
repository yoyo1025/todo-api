package persistence

import (
	"fmt"
	"testing"
	"todo-api/database"
	"todo-api/domain/model"
	"todo-api/infrastructure/record"

	"github.com/stretchr/testify/assert"
)

// SignUp テスト
func TestUserPersistence_SignUp(t *testing.T) {
	db := database.SetupTestDB()
	repo := NewCommandUserPersistence(db)

	t.Run("SignUp Success", func(t *testing.T) {
		newUser := model.NewUser("Test User", "test@example.com")
		err := repo.SignUp(newUser)
		assert.NoError(t, err, "ユーザーの登録成功")

		// DB確認
		var rec record.UserRecord
		db.First(&rec, "email = ?", "test@example.com")

		assert.Equal(t, "Test User", rec.Name, "登録されたユーザー名が正しくありません")
		assert.Equal(t, "test@example.com", rec.Email, "登録されたメールアドレスが正しくありません")
	})

	// t.Run("SignUp Failure - Duplicate Email", func(t *testing.T) {
	// 	existingUser := model.NewUser("Existing User", "duplicate@example.com")
	// 	_ = repo.SignUp(existingUser) // 初回登録成功を想定

	// 	// 同じメールアドレスで再度登録
	// 	duplicateUser := model.NewUser("Another User", "duplicate@example.com")
	// 	err := repo.SignUp(duplicateUser)

	// 	assert.Error(t, err, "重複したメールアドレスはエラーが発生するはず")
	// })
}

// FindByEmail テスト
func TestUserPersistence_FindByEmail(t *testing.T) {
	db := database.SetupTestDB()
	cmdRepo := NewCommandUserPersistence(db)
	queryRepo := NewQueryUserPersistence(db)

	// 事前にデータを作成
	user := model.NewUser("Find Test", "find@test.com")
	_ = cmdRepo.SignUp(user)

	t.Run("FindByEmail Success", func(t *testing.T) {
		foundUser, err := queryRepo.FindByEmail("find@test.com")
		fmt.Println(foundUser.ID)
		assert.NoError(t, err, "検索が失敗してはいけない")
		assert.NotNil(t, foundUser, "ユーザーは存在するはず")
		assert.Equal(t, "Find Test", foundUser.Name, "取得したユーザー名が異なります")
		assert.Equal(t, "find@test.com", foundUser.Email, "取得したメールアドレスが異なります")
	})

	t.Run("FindByEmail Failure - Not Found", func(t *testing.T) {
		foundUser, err := queryRepo.FindByEmail("notfound@test.com")
		assert.Nil(t, foundUser, "存在しないユーザーはnilのはず")
		assert.NoError(t, err, "存在しない場合でもエラーが返らない想定")
	})
}