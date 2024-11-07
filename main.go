package main

import (
	"acme-manager/logger"
)

import (
	"acme-manager/config"
	"acme-manager/database"
	"acme-manager/graphql"
	_ "acme-manager/secret"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func main() {
	database.Connect()
	database.Migration()
	database.Seed()
	defer database.Close()

	cfg := config.Get()
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	graphql.MapGraphQLRoutes(app)
	logger.Infof("Server listening on port %d", cfg.Server.Port)
	err := app.Listen(":" + strconv.Itoa(int(cfg.Server.Port)))
	if err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
