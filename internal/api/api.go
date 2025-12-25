package api

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/Wakisa/maka/internal/config"
	"www.github.com/Wakisa/maka/internal/services"
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
		ScoresService: services.NewScoresService(cfg.Football),
	}
}

func (api *API) registerRoutes() {
	api.routes()
}

func (api *API) Start(address string) error {
	api.registerRoutes()
	return api.app.Listen(address)
}
