package command

import (
	"net/http"
	"todo-api/domain/model"
	commandRepo "todo-api/domain/repository"

	"github.com/labstack/echo/v4"
)

type ICommandUserUsecase interface {
	// ユーザを新規登録する
	SingUp(c echo.Context, name, email string) error
}

type CommandUserUsecase struct {
	commandUserRepository commandRepo.ICommandUserRepository
}

func NewCommandUserUsecase(commandUserRepository commandRepo.ICommandUserRepository) ICommandUserUsecase {
	return &CommandUserUsecase {
		commandUserRepository: commandUserRepository,
	}
}

func (uu *CommandUserUsecase) SingUp(c echo.Context, name, email string) error {
	user := model.NewUser(name, email)
	if err := uu.commandUserRepository.SignUp(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "ユーザの取得に失敗")
	}
	return nil
}