package controller

import (
	"fmt"
	"net/http"
	model "server/model"

	"github.com/labstack/echo/v4"
)

func MakeMatch(c echo.Context) error {
	model := model.Match{}
	c.Bind(&model)
	fmt.Println(model)
	return c.JSON(http.StatusOK, model)
}

func EditMatch(c echo.Context) error {

	return c.JSON(http.StatusOK, "edit match")
}

func SearchMatch(c echo.Context) error {
	return c.JSON(http.StatusOK, "search match")
}

func FetchAllMatch(c echo.Context) error {
	// return c.JSON(http.StatusOK, {matches: [1, 2, 3]})
	m1 := model.Match{Id: 111, Team1_id: 1, Team2_id: 2, Start_time: "3"}
	m2 := model.Match{Id: 222, Team1_id: 12, Team2_id: 23, Start_time: "333"}
	m3 := model.Match{Id: 333, Team1_id: 13, Team2_id: 222, Start_time: "3121"}
	return c.JSON(http.StatusOK, [3]model.Match{m1, m2, m3})
}

func DeleteMatch(c echo.Context) error {
	return c.JSON(http.StatusOK, "deleted match")
}
