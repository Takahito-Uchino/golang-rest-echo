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

func GetTodos(c echo.Context) error {
	todos := []model.Todo{}
	model.DB.Find(&todos)
	return c.JSON(http.StatusOK, todos)
}

func GetTodo(c echo.Context) error {
	todo := model.Todo{}
	if err := c.Bind(&todo); err != nil {
		return err
	}
	model.DB.Take(&todo)
	return c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c echo.Context) error {
	todo := model.Todo{}
	if err := c.Bind(&todo); err != nil {
		return err
	}
	model.DB.Save(&todo)
	return c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c echo.Context) error {
	todo := model.Todo{}
	if err := c.Bind(&todo); err != nil {
		return err
	}
	model.DB.Save(&todo)
	return c.JSON(http.StatusOK, todo)
}
