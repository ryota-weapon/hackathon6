package model

type Match struct {
	Id         int    `json:"id"`
	Team1_id   int    `json:"team1"`
	Team2_id   int    `json:"team2"`
	Start_time string `json:"startTime"`
	// Players
}
