package repository

import "todo-api/domain/model"

type IUserRepository interface {
	// 登録されたメールアドレスに基づいてユーザ情報を取得する
	// 未登録のユーザであった場合、model.user.name が空になる
	FindByEmail(email string) (*model.User, error)
	// ユーザを新規登録する
	SignUp(user *model.User) (*model.User, error)
}