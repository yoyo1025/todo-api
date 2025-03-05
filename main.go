package main

import (
	"fmt"
	"net/http"
	"todo-api/database"
	"todo-api/infrastructure/persistence"
	"todo-api/presentation/handler"
	"todo-api/usecase"
	commandtask "todo-api/usecase/task/command-task"
	querytask "todo-api/usecase/task/query-task"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	db := database.GetDB()
	app := echo.New()

	// BodyDumpミドルウェアを追加
	app.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Printf("Request Body: %s\n", string(reqBody))
		fmt.Printf("Response Body: %s\n", string(resBody))
	}))

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins:     []string{
			"http://localhost:3000",
			// "https://todo-front-ochre.vercel.app",
		},
    AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
    AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
    AllowCredentials: true,
}))

	userPersistence := persistence.NewUserPersistence(db)
	userUsecase := usecase.NewUserUsecase(userPersistence)
	userHandler := handler.NewUserHandler(userUsecase)

	app.POST("/api/auth/github", userHandler.HandleLogin)

	taskCommandPersistence := persistence.NewTaskCommandPersistence(db)
	taskQueryPersistence := persistence.NewTaskQueryPersistence(db)
	taskCommandUsecase := commandtask.NewTaskCommandUsecase(taskCommandPersistence)
	taskQueryUsecase := querytask.NewTaskQueryUsecase(taskQueryPersistence)

	taskHandler := handler.NewTaskHandler(taskCommandUsecase, taskQueryUsecase)

	app.GET("/task/:userId", taskHandler.HandleGetAllTasks)
	app.POST("/task/:userId", taskHandler.HandleCreateTask)
	app.PUT("/task/:userId/:taskId", taskHandler.HandleUpdateTask)

	app.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world!")
	})

	app.Logger.Fatal(app.Start(":3000"))
}