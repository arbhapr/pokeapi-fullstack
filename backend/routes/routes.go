package routes

import (
	"poke-go/controller" // Adjust the import path based on your project structure

	"poke-go/model"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(app *fiber.App, cfg model.Config) {
	// Route Group of V1
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/api/v1/pokemon", fiber.StatusMovedPermanently)
	})
	api := app.Group("/api/v1")
	api.Get("/pokemon", func(c *fiber.Ctx) error {
		return controller.GetPokemonList(c, cfg)
	})
	api.Get("/pokemon/:idPokemon", func(c *fiber.Ctx) error {
		return controller.GetPokemonDetail(c, cfg)
	})
	api.Post("/pokemon/:idPokemon", func(c *fiber.Ctx) error {
		return controller.CatchPokemon(c, cfg)
	})
	api.Get("/my-pokemon", func(c *fiber.Ctx) error {
		return controller.MyPokemonList(c, cfg)
	})
	api.Patch("/my-pokemon/:idPokemon", func(c *fiber.Ctx) error {
		return controller.RenamePokemon(c, cfg)
	})
	api.Delete("/my-pokemon/:idPokemon", func(c *fiber.Ctx) error {
		return controller.ReleasePokemon(c, cfg)
	})
}
