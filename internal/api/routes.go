package api

import (
	"github.com/gofiber/fiber/v2"
	apiv1 "www.github.com/Wakisa/maka/internal/api/v1"
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
