package persistence

import (
	"todo-api/domain/model"
	commandRepo "todo-api/domain/repository"
	"todo-api/infrastructure/record"
	"todo-api/infrastructure/response"
	queryRepo "todo-api/usecase/query/repository"

	"gorm.io/gorm"
)

type UserPersistence struct {
	db *gorm.DB
}

func NewCommandUserPersistence(db *gorm.DB) commandRepo.ICommandUserRepository  {
	return &UserPersistence {
		db: db,
	}
}

func NewQueryUserPersistence(db *gorm.DB) queryRepo.IQueryUserRepository {
	return &UserPersistence{
		db: db,
	}
}

func (up *UserPersistence) FindByEmail(email string) (*response.User, error)  {
	var user *response.User
	result := up.db.Table("user_records").Where("email = ?", email).Scan(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}


func (up *UserPersistence) SignUp(user *model.User) error {
	// トランザクション開始
	tx := up.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// UserRecord の作成
	userRecord := &record.UserRecord{
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}

	// ユーザーを作成
	if err := tx.Create(&userRecord).Error; err != nil {
		tx.Rollback()
		return err
	}

	// トランザクションをコミット
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
