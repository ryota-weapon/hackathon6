package controller

import (
	"fmt"
	"net/http"
	infra "server/infrastructure"
	model "server/model"

	"github.com/labstack/echo/v4"
)

func SignUpUser(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	fmt.Println(user.Auth_id)
	fmt.Println(user.Name)

	err := infra.RegisterUser(user)
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, "created")
}
