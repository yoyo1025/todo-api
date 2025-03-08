package usecase

import (
	"net/http"
	"todo-api/domain/model"
	commandRepo "todo-api/domain/repository"

	"github.com/labstack/echo/v4"
)

type IUserUsecase interface {
	// 登録されたメールアドレスに基づいてユーザ情報を取得する
	// 未登録のユーザであった場合、model.user.name が空になる
	FindByEmail(c echo.Context, email string) (*model.User, error)
	// ユーザを新規登録する
	SingUp(c echo.Context, name, email string) (*model.User, error)
}

type UserUsecase struct {
	userRepository commandRepo.IUserRepository
}

func NewUserUsecase(userRepository commandRepo.IUserRepository) IUserUsecase {
	return &UserUsecase {
		userRepository: userRepository,
	}
}

func (uu *UserUsecase) FindByEmail(c echo.Context, email string) (*model.User, error) {
	user, err := uu.userRepository.FindByEmail(email)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "ユーザの取得に失敗")
	}
	return user, nil
}

func (uu *UserUsecase) SingUp(c echo.Context, name, email string) (*model.User, error) {
	user := model.NewUser(name, email)
	newUser, err := uu.userRepository.SignUp(user)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "ユーザの取得に失敗")
	}
	return newUser, nil
}