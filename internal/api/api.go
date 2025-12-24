package api

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/Wakisa/Maka-scores-update/internal/config"
)

type API struct {
	app           *fiber.App
	Config        *config.Config
	ScoresService services.ScoresService
}

func NewAPI(cfg *config.Config) *API {
	return &API{
		app:           fiber.New(),
		Config:        cfg,
		ScoresService: services.NewScoresService(cfg),
	}
}

func (api *API) registerRoutes() {
	api.middleware()
	api.routes()
}

func (api *API) Start(address string) error {
	api.registerRoutes()
	return api.app.Listen(address)
}
