package model

type Evaluation struct {
	Id       string `json:"id"`
	Value    int    `json:"value"`
	UserId   string `json:"userId"`
	MatchId  string `json:"matchId"`
	PlayerId string `json:"playerId"`
}
