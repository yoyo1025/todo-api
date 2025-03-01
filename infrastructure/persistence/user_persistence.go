package persistence

import (
	"todo-api/domain/model"
	"todo-api/domain/repository"

	"gorm.io/gorm"
)

type UserPersistence struct {
	db *gorm.DB
}

func NewUserPersistence(db *gorm.DB) repository.IUserRepository  {
	return &UserPersistence {
		db: db,
	}
}

func (up *UserPersistence) FindByEmail(email string) (*model.User, error)  {
	var ur UserRecord
	result := up.db.Where("email = ?", email).Find(&ur)
	if result.Error != nil {
		return nil, result.Error
	}

	user := model.NewUser(ur.Name, ur.Email)
	user.ID = ur.ID
	return user, nil
}


func (up *UserPersistence) SignUp(user *model.User) (*model.User, error) {
	// トランザクション開始
	tx := up.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// UserRecord の作成
	userRecord := &UserRecord{
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}

	// ユーザーを作成
	if err := tx.Create(&userRecord).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// トランザクションをコミット
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// model.User を作成して返す
	newUser := model.NewUser(userRecord.Name, userRecord.Email)
	newUser.ID = userRecord.ID

	return newUser, nil
}
