package model

type Match struct {
	Id         string   `json:"id"`
	Team1_id   string   `json:"team1"`
	Team2_id   string   `json:"team2"`
	Start_time string   `json:"startTime"`
	Players1   []string `json:"players1"`
	Players2   []string `json:"players2"`
}
