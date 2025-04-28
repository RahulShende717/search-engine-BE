package main

import (
	"search-eng/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // allow all origins, or you can specify "http://localhost:3000"
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	app.Post("/upload", api.DataHandler)
	app.Get("/search", api.SearchHandler)

	app.Listen(":3000")
}
