package controller

import (
	"fmt"
	"poke-go/helper"
	"poke-go/model"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var validate = validator.New()

// MyPokemonList returns the list of caught Pokémon
func MyPokemonList(c *fiber.Ctx, cfg model.Config) error {
	caughtPokemonList := []model.CaughtPokemon{}
	if model.MyPokemonList != nil {
		caughtPokemonList = model.MyPokemonList
	}

	return c.JSON(fiber.Map{
		"my_pokemons": caughtPokemonList,
	})
}

// RenamePokemon renames a caught Pokémon by UUID
func RenamePokemon(c *fiber.Ctx, cfg model.Config) error {
	id, err := helper.ParseUUIDParam(c, "idPokemon")
	if err != nil {
		return err
	}

	var renameReq model.NicknameRequest
	if err := c.BodyParser(&renameReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": err.Error(),
			"success": false,
		})
	}

	if err := validate.Struct(&renameReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"message": "Nickname is required",
			"success": false,
		})
	}

	if !updatePokemonNickname(id, renameReq.Nickname) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "Pokémon not found",
			"message": "No Pokémon found with the provided UUID",
			"success": false,
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Successfully renamed Pokémon to: %s", renameReq.Nickname),
		"success": true,
	})
}

// CatchPokemon attempts to catch a Pokémon and adds it to the list
func CatchPokemon(c *fiber.Ctx, cfg model.Config) error {
	idPokemon := c.Params("idPokemon")
	detail, err := fetchPokemonDetail(cfg.SourceURL, idPokemon)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to fetch Pokémon details",
			"message": err.Error(),
			"success": false,
		})
	}

	if !helper.RandomCatchSuccess() {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Failed to catch the Pokémon.",
		})
	}

	caughtPokemon := model.CaughtPokemon{
		ID:      uuid.New(),
		Pokemon: *detail,
		Nickname: fmt.Sprintf("%s-%d",
			helper.GenerateRandomNickname(detail.Name),
			helper.GenerateFibonacci(len(model.MyPokemonList)+1)),
	}

	model.MyPokemonList = append(model.MyPokemonList, caughtPokemon)

	return c.JSON(fiber.Map{
		"success": true,
		"pokemon": detail,
		"message": fmt.Sprintf("Successfully caught Pokémon: %s", detail.Name),
	})
}

// ReleasePokemon removes a Pokémon from the caught list by UUID
func ReleasePokemon(c *fiber.Ctx, cfg model.Config) error {
	id, err := helper.ParseUUIDParam(c, "idPokemon")
	if err != nil {
		return err
	}

	if !helper.RandomReleaseSuccess() {
		return c.JSON(fiber.Map{
			"message": "Failed to release the Pokémon.",
			"success": false,
		})
	}

	for i, pokemon := range model.MyPokemonList {
		if pokemon.ID == id {
			model.MyPokemonList = append(model.MyPokemonList[:i], model.MyPokemonList[i+1:]...)
			return c.JSON(fiber.Map{
				"success": true,
				"message": "Successfully released Pokémon",
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error":   "Pokémon not found",
		"message": "No Pokémon found with the provided UUID",
		"success": false,
	})
}

// updatePokemonNickname updates the nickname of a Pokémon by UUID
func updatePokemonNickname(id uuid.UUID, newNickname string) bool {
	for i, caughtPokemon := range model.MyPokemonList {
		if caughtPokemon.ID == id {
			model.MyPokemonList[i].Nickname = newNickname
			return true
		}
	}
	return false
}
