package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"poke-go/helper"
	"poke-go/model"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// GetPokemonList fetches and returns the Pokémon list
func GetPokemonList(c *fiber.Ctx, cfg model.Config) error {
	limitPerPage, _ := strconv.Atoi(c.Params("limit"))
	if limitPerPage == 0 {
		limitPerPage = cfg.LimitPerPage
	}
	pagination := model.ExtractPagination(c, limitPerPage)
	paginatedResponse, err := fetchPokemonList(cfg, pagination)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to fetch Pokémon list",
			"message": err.Error(), // Include detailed error message
		})
	}

	return c.JSON(paginatedResponse)
}

func fetchPokemonList(cfg model.Config, pagination model.Pagination) (model.PaginatedResponse, error) {
	url := fmt.Sprintf("%s/api/v2/pokemon?limit=%d&offset=%d", cfg.SourceURL, pagination.Limit, pagination.Offset)

	resp, err := http.Get(url)
	if err != nil {
		return model.PaginatedResponse{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.PaginatedResponse{}, fmt.Errorf("request failed with status: %s", resp.Status)
	}

	var pokemonResponse model.PokemonResponse
	if err := json.NewDecoder(resp.Body).Decode(&pokemonResponse); err != nil {
		return model.PaginatedResponse{}, fmt.Errorf("failed to decode JSON: %w", err)
	}

	var pokemons []model.Pokemon
	for _, p := range pokemonResponse.Results {
		id := extractPokemonID(p.URL)
		if id == "" {
			continue
		}

		// Getting photo of each Pokemon
		detailUrl := fmt.Sprintf("%s/api/v2/pokemon/%s", cfg.SourceURL, id)
		respDetail, err := http.Get(detailUrl)
		if err != nil {
			return model.PaginatedResponse{}, fmt.Errorf("failed to send request: %w", err)
		}
		defer respDetail.Body.Close()

		var pokemonDetailAPI model.PokemonDetailAPI
		if err := json.NewDecoder(respDetail.Body).Decode(&pokemonDetailAPI); err != nil {
			return model.PaginatedResponse{}, fmt.Errorf("failed to decode JSON: %w", err)
		}
		// Getting photo of each Pokemon

		pokemons = append(pokemons, model.Pokemon{
			ID:     id,
			URL:    fmt.Sprintf("%s/pokemon/%s", cfg.BaseURL, id),
			Name:   helper.Ucwords(p.Name),
			Sprite: pokemonDetailAPI.Sprites.FrontDefault,
		})
	}

	// Calculate pagination info
	totalData := pokemonResponse.Count // Assuming the API response includes a count of total items
	nextPage := ""
	prevPage := ""

	if pagination.Offset+pagination.Limit < totalData {
		nextPage = fmt.Sprintf("%s/pokemon?limit=%d&offset=%d", cfg.BaseURL, pagination.Limit, pagination.Offset+pagination.Limit)
	}

	if pagination.Offset > 0 {
		prevPage = fmt.Sprintf("%s/pokemon?limit=%d&offset=%d", cfg.BaseURL, pagination.Limit, max(0, pagination.Offset-pagination.Limit))
	}

	paginationInfo := model.PaginationInfo{
		NextPage:  nextPage,
		PrevPage:  prevPage,
		TotalData: totalData,
	}

	return model.PaginatedResponse{
		Data:       pokemons,
		Pagination: paginationInfo,
	}, nil
}

// GetPokemonDetail fetches detailed information of a Pokémon by ID or name
func GetPokemonDetail(c *fiber.Ctx, cfg model.Config) error {
	idPokemon := c.Params("idPokemon")
	detail, err := fetchPokemonDetail(cfg.SourceURL, idPokemon)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to fetch Pokémon details",
			"message": err.Error(), // Include detailed error message
		})
	}
	return c.JSON(detail)
}

func fetchPokemonDetail(sourceURL, idPokemon string) (*model.PokemonDetail, error) {
	url := fmt.Sprintf("%s/api/v2/pokemon/%s", sourceURL, idPokemon)

	var pokemonDetailAPI model.PokemonDetailAPI
	if err := helper.FetchData(url, &pokemonDetailAPI); err != nil {
		return nil, fmt.Errorf("failed to fetch Pokémon detail: %w", err)
	}

	var pokemonDetail model.PokemonDetail
	pokemonDetail.ID, _ = strconv.Atoi(idPokemon)
	pokemonDetail.Name = helper.Ucwords(pokemonDetailAPI.Name)
	pokemonDetail.Sprite = pokemonDetailAPI.Sprites.FrontDefault
	pokemonDetail.Types = helper.RemapTypes(pokemonDetailAPI.Types)
	pokemonDetail.Moves = helper.RemapMoves(pokemonDetailAPI.Moves)
	return &pokemonDetail, nil
}

// extractPokemonID extracts Pokémon ID from the URL
func extractPokemonID(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) > 0 {
		return parts[len(parts)-2]
	}
	return ""
}
