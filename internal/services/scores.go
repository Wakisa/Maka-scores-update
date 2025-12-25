package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"www.github.com/Wakisa/maka/internal/config"
	"www.github.com/Wakisa/maka/internal/schema"
)

type ScoresService interface {
	FetchLiveScores(competition string) ([]schema.ScoreResponse, error)
	FetchFinishedScores(competition string) ([]schema.ScoreResponse, error)
}

type scoresServiceImpl struct {
	footballCfg config.FootballConfiguration
}

func NewScoresService(cfg config.FootballConfiguration) ScoresService {
	return &scoresServiceImpl{
		footballCfg: cfg,
	}
}

func (s *scoresServiceImpl) FetchLiveScores(competition string) ([]schema.ScoreResponse, error) {
	url := fmt.Sprintf("%s/competitions/%s/matches?status=LIVE", s.footballCfg.BaseURL, competition)
	return s.fetchMatches(url)
}

func (s *scoresServiceImpl) FetchFinishedScores(competition string) ([]schema.ScoreResponse, error) {
	url := fmt.Sprintf("%s/competitions/%s/matches?status=FINISHED", s.footballCfg.BaseURL, competition)
	return s.fetchMatches(url)
}

func (s *scoresServiceImpl) fetchMatches(url string) ([]schema.ScoreResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Auth-Token", s.footballCfg.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var raw schema.FootballMatchesResponse
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	var results []schema.ScoreResponse
	for _, m := range raw.Matches {
		results = append(results, schema.ScoreResponse{
			HomeTeam:  m.HomeTeam.Name,
			AwayTeam:  m.AwayTeam.Name,
			HomeScore: m.Score.FullTime.Home,
			AwayScore: m.Score.FullTime.Away,
			Status:    m.Status,
			MatchDate: m.UtcDate,
		})
	}
	return results, nil
}
