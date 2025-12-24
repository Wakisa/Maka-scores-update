package services

import "www.github.com/Wakisa/Maka-scores-update/internal/config"

type ScoresService struct {
	client *client.FootballClient
}

func NewScoresService(cfg *config.Config) ScoresService {
	return ScoresService{
		client: clients.NewFootballClient(cfg.Football.APIKey),
	}
}

func (s ScoresService) FetchLiveScores(competition string) (map[string]interface{}, error) {
	return s.client.GetLiveMatches(competition)
}

func (s ScoresService) FetchFinishedScores(competition string) (map[string]interface{}, error) {
	return s.client.GetFinishedMatches(competition)
}
