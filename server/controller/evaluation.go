package controller

import (
	// "net/http"
	// "server/infrastructure"

	"fmt"
	"log"
	"net/http"
	"server/infrastructure"
	"server/model"

	"github.com/labstack/echo/v4"
	"google.golang.org/api/iterator"
)

func AggregatedEvaluations(c echo.Context) error {
	e := model.Evaluation{}
	c.Bind(&e)

	playerIds := []string{"1", "2"}

	res := map[string]int{}
	// client, ctx := infrastructure.ConnectToFirestore()
	// evaluations := client.Collection("evaluations").Where("matchId", "==", e.MatchId).Documents(ctx)
	for i := 0; i < len(playerIds); i++ {
		res[playerIds[i]] = 2
	}

	return c.JSON(http.StatusOK, res)

	//map["player_id"]evaluation of an user
}

func EvaluatePlayer(c echo.Context) error {
	evaluation := model.Evaluation{}
	c.Bind(&evaluation)
	fmt.Println(evaluation)

	client, ctx := infrastructure.ConnectToFirestore()
	evaluationData := client.Collection("evaluations").Where("userId", "==", evaluation.UserId).Where("playerId", "==", evaluation.PlayerId).Where("matchId", "==", evaluation.MatchId).Documents(ctx)
	evaluateExist := false

	for {
		doc, err := evaluationData.Next()
		if err == iterator.Done {
			fmt.Println("~~~~~end~~~~~~")
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		evaluateExist = true
		doc.Ref.Set(ctx, map[string]interface{}{
			"value":    evaluation.Value,
			"userId":   evaluation.UserId,
			"playerId": evaluation.PlayerId,
			"matchId":  evaluation.MatchId,
		})
		fmt.Println("------updated-----")
	}

	if !evaluateExist {
		client.Collection("evaluations").Add(ctx, map[string]interface{}{
			"value":    evaluation.Value,
			"userId":   evaluation.UserId,
			"playerId": evaluation.PlayerId,
			"matchId":  evaluation.MatchId,
		})
		fmt.Println("-----added------")
	}

	return c.JSON(http.StatusOK, "evaluated")
}
