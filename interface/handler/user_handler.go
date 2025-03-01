package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"todo-api/interface/dto"
	"todo-api/usecase"

	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	HandleLogin(c echo.Context) error
}

type UserHandler struct {
	userUsecase usecase.IUserUsecase
}

func NewUserHandler(userUsecase usecase.IUserUsecase) IUserHandler  {
	return &UserHandler {
		userUsecase: userUsecase,
	}
}

// リクエストヘッダーに含まれるGitHubアクセストークンを取得する
func getGitHubToken(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
				return "", echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is required")
		}

		// Bearerを除去
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
				return "", echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
		}
		accessToken := parts[1]
		return accessToken, nil
}

// 取得したGitHubアクセストークンが認証済みか問い合わせ、トークンの持ち主のGitHubユーザ名を取得する
func getGitHubUserName(token string) (string, error) {
	client := &http.Client{}
		req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
		if err != nil {
				return "", echo.NewHTTPError(http.StatusInternalServerError, "Failed to create request")
		}
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := client.Do(req)
		if err != nil {
				return "", echo.NewHTTPError(http.StatusInternalServerError, "Failed to call GitHub API")
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
				return "", fmt.Errorf("failed to get userName: %s", resp.Status)
		}

		var user dto.GitHubUser
		if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
			return "", err
		}

		return user.Login, nil
}

// 取得したGitHubアクセストークンが認証済みか問い合わせ、トークンの持ち主のメールアドレスを取得する
func getGitHubEmail(token string) (string, error) {
	client := &http.Client{}
		req, err := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
		if err != nil {
				return "", echo.NewHTTPError(http.StatusInternalServerError, "Failed to create request")
		}
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := client.Do(req)
		if err != nil {
				return "", echo.NewHTTPError(http.StatusInternalServerError, "Failed to call GitHub API")
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
				return "", fmt.Errorf("failed to get emails: %s", resp.Status)
		}

		var emails []dto.GitHubEmail
		if err := json.NewDecoder(resp.Body).Decode(&emails); err != nil {
			return "", err
		}

		return emails[0].Email, nil
}

func (uh *UserHandler) HandleLogin(c echo.Context) error {
	accessToken, err := getGitHubToken(c)
	if err != nil {
		return err
	}

	username, err := getGitHubUserName(accessToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	email, err := getGitHubEmail(accessToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	user, err := uh.userUsecase.FindByEmail(c, email)
	if err != nil {
		return err
	}

	if user.GetEmail() == "" {
		newUser, err := uh.userUsecase.SingUp(c, username, email)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, dto.ToUserResponse(newUser))
	}
	return c.JSON(http.StatusOK, dto.ToUserResponse(user))
}