package v1

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/Wakisa/maka/internal/services"
)

func GetLiveScores(service services.ScoresService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		comp := c.Params("competition")
		data, err := service.FetchLiveScores(comp)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"competition": comp,
			"count":       len(data),
			"matches":     data,
		})
	}
}

func GetFinishedScores(service services.ScoresService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		comp := c.Params("competition")
		data, err := service.FetchFinishedScores(comp)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"competition": comp,
			"count":       len(data),
			"matches":     data,
		})
	}
}
