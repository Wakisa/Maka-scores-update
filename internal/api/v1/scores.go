package v1

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"www.github.com/Wakisa/maka/internal/schema"
	"www.github.com/Wakisa/maka/internal/services"
)

func GetUpcomingScores(service services.ScoresService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		comp := c.Params("competition")
		page, _ := strconv.Atoi(c.Query("page", "0"))
		limit, _ := strconv.Atoi(c.Query("limit", "10"))

		data, err := service.FetchUpcomingScores(comp)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// TOD DO DELETE DEBUG LOG
		fmt.Printf("COMP: %s | PAGE: %d | LIMIT: %d | COUNT: %d | RETURNED: %d\n",
			comp, page, limit, len(data), len(paginate(data, page, limit)))

		return c.JSON(fiber.Map{
			"competition": comp,
			"count":       len(data),
			"matches":     paginate(data, page, limit),
		})
	}
}

func GetLiveScores(service services.ScoresService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		comp := c.Params("competition")
		page, _ := strconv.Atoi(c.Query("page", "0"))
		limit, _ := strconv.Atoi(c.Query("limit", "10"))

		data, err := service.FetchLiveScores(comp)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"competition": comp,
			"count":       len(data),
			"matches":     paginate(data, page, limit),
		})
	}
}

func GetFinishedScores(service services.ScoresService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		comp := c.Params("competition")
		page, _ := strconv.Atoi(c.Query("page", "0"))
		limit, _ := strconv.Atoi(c.Query("limit", "10"))

		data, err := service.FetchFinishedScores(comp)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"competition": comp,
			"count":       len(data),
			"matches":     paginate(data, page, limit),
		})
	}
}

func paginate(data []schema.ScoreResponse, page, limit int) []schema.ScoreResponse {
	if limit <= 0 {
		limit = 10 // current default size
	}
	if page < 0 {
		page = 0
	}

	start := page * limit
	if start >= len(data) {
		return []schema.ScoreResponse{}
	}

	end := start + limit
	if end > len(data) {
		end = len(data)
	}
	return data[start:end]
}
