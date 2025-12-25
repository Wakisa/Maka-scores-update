package schema

type ScoreResponse struct {
	HomeTeam  string `json:"home_team"`
	AwayTeam  string `json:"away_team"`
	HomeScore int    `json:"home_score"`
	AwayScore int    `json:"away_score"`
	Status    string `json:"status"`
	MatchDate string `json:"match_date"`
}

type FootballMatchesResponse struct {
	Matches []Match `json:"matches"`
}

type Team struct {
	Name string `json:"name"`
}

type FullTimeScore struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type Score struct {
	FullTime FullTimeScore `json:"fullTime"`
}

type Match struct {
	HomeTeam Team   `json:"homeTeam"`
	AwayTeam Team   `json:"awayTeam"`
	Score    Score  `json:"score"`
	Status   string `json:"status"`
	UtcDate  string `json:"utcDate"`
}
