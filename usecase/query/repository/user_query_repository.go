package queryRepo

import (
	"todo-api/infrastructure/response"
)


type IQueryUserRepository interface {
	// 登録されたメールアドレスに基づいてユーザ情報を取得する
	// 未登録のユーザであった場合、model.user.name が空になる
	FindByEmail(email string) (*response.User, error)
}