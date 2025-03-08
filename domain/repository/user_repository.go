package commandRepo

import "todo-api/domain/model"

type ICommandUserRepository interface {
	// ユーザを新規登録する
	SignUp(user *model.User) (*model.User, error)
}