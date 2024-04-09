package controller

import (
	"net/http"

	"github.com/Takahito-Uchino/golang-rest-echo/model"
	"github.com/labstack/echo/v4"
)

func CreateTodo(c echo.Context) error {
	todo := model.Todo{}
	if err := c.Bind(&todo); err != nil {
		return err
	}
	model.DB.Create(&todo)
	return c.JSON(http.StatusCreated, todo)
}
