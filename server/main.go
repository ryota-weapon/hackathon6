package main

import (
	"net/http"
	// "os"

	controller "server/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/", handler)
	e.POST("/", controller.SignUpUser)
	e.GET("/user/:id", controller.FindUser)

	e.POST("/makeMatch", controller.MakeMatch)
	e.GET("/matches", controller.FetchAllMatch)
	e.GET("/match/:id", controller.FindMatch)

	// Addr := ":" + os.Getenv("PORT")
	Addr := ":1323"
	e.Logger.Fatal(e.Start(Addr))
}

func handler(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello world")
}
