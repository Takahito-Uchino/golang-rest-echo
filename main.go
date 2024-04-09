package main

import (
	"net/http"

	"github.com/Takahito-Uchino/golang-rest-echo/model"
	"github.com/labstack/echo/v4"
)

func connect(c echo.Context) error {
	db, _ := model.DB.DB()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		return c.String(http.StatusInternalServerError, "database connection failed")
	} else {
		return c.String(http.StatusOK, "database connect")
	}
}

func main() {
	e := echo.New()
	e.GET("/", connect)
	e.Logger.Fatal(e.Start(":8080"))
}
