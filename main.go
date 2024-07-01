package main

import (
	"fmt"

	"github.com/Kraeutersalz/rts-games-ranker/data"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong!")
	})

	// Get a List of all Players
	app.Get("/players", func(c *fiber.Ctx) error {
		return c.SendString("List of Players")
	})

	// Calculate total matches for a mock player
	app.Get("/totalMatches", func(c *fiber.Ctx) error {
		// Mock player
		player := data.Player{
			Name:   "John Doe",
			Wins:   10,
			Losses: 5,
			Rating: 1500.0,
		}

		total := data.TotalMatches(player)
		return c.SendString(fmt.Sprintf("Total Matches: %d", total))
	})

	app.Listen(":3000")
}
