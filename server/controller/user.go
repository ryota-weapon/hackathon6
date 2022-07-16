package controller

import (
	"fmt"
	"net/http"
	"server/infrastructure"
	infra "server/infrastructure"
	model "server/model"

	"github.com/labstack/echo/v4"
)

func SignUpUser(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := infra.RegisterUser(user)
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, "created")
}

func FindUser(c echo.Context) error {
	user_id := c.Param("id")
	fmt.Println(user_id)

	client, ctx := infrastructure.ConnectToFirestore()

	user, err := client.Collection("users").Doc(user_id).Get(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, user.Data())
}
