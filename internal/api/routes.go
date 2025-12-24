package api

import (
	apiv1 "github.com/Wakisa/maka-scores-update/internal/api/v1"
	"github.com/gofiber/fiber/v2"
)

func (api *API) routes() {
	router := api.app

	// Health check
	router.Get("/healthz", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Versioned routes
	v1 := router.Group("/v1")
	{
		v1.Get("/scores/live/:competition", apiv1.GetLiveScores(api.ScoresService))
		v1.Get("/scores/finished/:competition", apiv1.GetFinishedScores(api.ScoresService))
	}
}
