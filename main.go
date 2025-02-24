package main

import (
	"fmt"
	"todo-api/database"
	"todo-api/infrastructure/persistence"
	"todo-api/interface/handler"
	"todo-api/usecase"

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

	taskPersistence := persistence.NewTaskPersistence(db)
	taskUsecase := usecase.NewTaskUsecase(taskPersistence)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	app.GET("/task/:userId", taskHandler.HandleGetAllTasks)
	app.POST("/task/:userId", taskHandler.HandleCreateTask)
	app.PUT("/task/:userId/:taskId", taskHandler.HandleUpdateTask)

	app.Logger.Fatal(app.Start(":3000"))
}