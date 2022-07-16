package controller

import (
	"net/http"
	"server/infrastructure"
	model "server/model"

	"github.com/labstack/echo/v4"
)

func MakeMatch(c echo.Context) error {
	match := model.Match{}
	c.Bind(&match)
	infrastructure.SaveMatch(match)

	return c.JSON(http.StatusOK, "saved")
}

func EditMatch(c echo.Context) error {

	return c.JSON(http.StatusOK, "edit match")
}

func SearchMatch(c echo.Context) error {
	return c.JSON(http.StatusOK, "search match")
}

func FindMatch(c echo.Context) error {
	// id := c.Param("id")

	// m := model.Match{Id: 11, Team1_id: 1, Team2_id: 2, Start_time: "3"}
	// res := []map[string]interface{}
	// res["id"] = m.Id
	// res["team1"] = convert(m.Team1_id)
	// res["team2"] = convert(m.Team2_id)
	// res["startTime"] = m.Start_time

	// client, ctx := infrastructure.ConnectToFirestore()
	// match, err := client.Collection("matches").Doc(id).Get(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil
	// }

	res := map[string]interface{}{"id": 11, "team1": "日本", "team2": "ブラジル", "startTime": "2022-07-20"}

	return c.JSON(http.StatusOK, res)
}

func FetchAllMatch(c echo.Context) error {
	// m1 := model.Match{Id: 111, Team1_id: 1, Team2_id: 2, Start_time: "3"}
	// m2 := model.Match{Id: 222, Team1_id: 12, Team2_id: 23, Start_time: "333"}
	// m3 := model.Match{Id: 333, Team1_id: 13, Team2_id: 222, Start_time: "3121"}
	// return c.JSON(http.StatusOK, [3]model.Match{m1, m2, m3})

	matches := infrastructure.GetMatches()
	return c.JSON(http.StatusOK, matches)

}

func DeleteMatch(c echo.Context) error {
	return c.JSON(http.StatusOK, "deleted match")
}
