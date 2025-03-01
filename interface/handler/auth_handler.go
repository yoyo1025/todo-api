package handler

import (
	"fmt"
	"net/http"
	"strings"
	"todo-api/interface/dto"

	"github.com/labstack/echo/v4"
)

// ハンドラ実装（interface/handler 配下など）
func HandleGitHubAuth() echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is required")
		}

		// Bearerを除去
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
		}
		accessToken := parts[1]

		fmt.Print("accessToken: ")
		fmt.Println(accessToken)

		// GitHubのユーザ情報を取得してみる
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
		if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create request")
		}
		req.Header.Set("Authorization", "Bearer "+accessToken)
		resp, err := client.Do(req)
		if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to call GitHub API")
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid GitHub token")
		}

		var user dto.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to get json")
		}

		fmt.Println(user.Name)
		fmt.Println(user.Email)

			// 4. ユーザーIDをレスポンス
			return c.JSON(http.StatusOK, dto.GitHubAuthResponse{UserID: 1})
	}
}