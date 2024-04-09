package main

import (
	"github.com/Takahito-Uchino/golang-rest-echo/controller"
	"github.com/Takahito-Uchino/golang-rest-echo/model"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db, _ := model.DB.DB()
	defer db.Close()

	e.POST("/todos", controller.CreateTodo)
	e.Logger.Fatal(e.Start(":8080"))
}
