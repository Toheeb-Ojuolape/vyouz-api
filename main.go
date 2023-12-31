package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Toheeb-Ojuolape/shopafrique-api/initializers"
	"github.com/Toheeb-Ojuolape/shopafrique-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	app := fiber.New()

	// change this in production
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Vyouz's API is LIVE 🥳",
		})
	})

	app.Get("/api/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Everything's ok, stop worrying",
		})
	})

	auth := app.Group("/api/auth")
	routes.AuthRoutes(auth)

	user := app.Group("/api/user")
	routes.UserRoutes(user)

	wallet := app.Group("/api/wallet")
	routes.WalletRoutes(wallet)

	campaigns := app.Group("/api/campaigns")
	routes.CampaignRoutes(campaigns)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
