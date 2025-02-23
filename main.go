package main

import (
	"fmt"
	"net/http"
	"todo-api/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	database.GetDB()
	app := echo.New()

	// BodyDumpミドルウェアを追加
	app.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Printf("Request Body: %s\n", string(reqBody))
		fmt.Printf("Response Body: %s\n", string(resBody))
	}))

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, world!!")
	})

	app.Logger.Fatal(app.Start(":3000"))
}