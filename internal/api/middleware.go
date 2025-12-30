package api

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func (api *API) middleware() {
	app := api.app

	// CORS - allow React frontend to call backend
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Panic recovery with stack trace logging
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			if api.logger != nil {
				api.logger.Error("failed to process a request", "errorStackTrace", fmt.Sprintf("panic: %v\n%s\n", e, debug.Stack()))
			} else {
				_, _ = os.Stderr.WriteString(fmt.Sprintf("panic: %v\n%s\n", e, debug.Stack()))
			}
		},
	}))

	// Request logging
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

}
