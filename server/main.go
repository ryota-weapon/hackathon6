package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/", handler)

	Addr := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(Addr))
}

func handler(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello world")
}
