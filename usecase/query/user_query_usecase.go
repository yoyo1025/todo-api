package query

import (
	"net/http"
	"todo-api/infrastructure/response"
	queryRepo "todo-api/usecase/query/repository"

	"github.com/labstack/echo/v4"
)



type IQueryUserUsecase interface {
	// 登録されたメールアドレスに基づいてユーザ情報を取得する
	// 未登録のユーザであった場合、model.user.name が空になる
	FindByEmail(c echo.Context, email string) (*response.User, error)
}

type QueryUserUsecase struct {
	queryUserRepository queryRepo.IQueryUserRepository
}

func NewQueryUserUsecase(queryUserRepository queryRepo.IQueryUserRepository) IQueryUserUsecase {
	return &QueryUserUsecase {
		queryUserRepository: queryUserRepository,
	}
}

func (uu *QueryUserUsecase) FindByEmail(c echo.Context, email string) (*response.User, error) {
	user, err := uu.queryUserRepository.FindByEmail(email)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "ユーザの取得に失敗")
	}
	return user, nil
}