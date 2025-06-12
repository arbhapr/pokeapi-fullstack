package helper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// FetchData performs an HTTP GET request and decodes the JSON response into the provided result interface.
// It returns an error if the request fails or the JSON decoding fails.
func FetchData(url string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	return nil
}

// RemapTypes transforms the nested types structure of a Pokémon into a flat structure.
// It extracts and returns the slot number and type name for easier access.
func RemapTypes(types []struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}) []struct {
	Name string `json:"name"`
} {
	flatTypes := make([]struct {
		Name string `json:"name"`
	}, len(types))

	for i, t := range types {
		flatTypes[i] = struct {
			Name string `json:"name"`
		}{
			Name: t.Type.Name,
		}
	}
	return flatTypes
}

// RemapMoves transforms the nested moves structure of a Pokémon into a flat structure.
// It extracts and returns the move name for easier access.
func RemapMoves(moves []struct {
	Move struct {
		Name string `json:"name"`
	} `json:"move"`
}) []struct {
	Name string `json:"name"`
} {
	flatMoves := make([]struct {
		Name string `json:"name"`
	}, len(moves))

	for i, m := range moves {
		flatMoves[i] = struct {
			Name string `json:"name"`
		}{
			Name: m.Move.Name,
		}
	}
	return flatMoves
}

// ParseUUIDParam parses a UUID from the request parameters in a Fiber context.
// It returns an error response if the UUID is invalid.
func ParseUUIDParam(c *fiber.Ctx, param string) (uuid.UUID, error) {
	id := c.Params(param)
	idParsed, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "Invalid UUID format",
			"message": fmt.Sprintf("Failed to parse UUID: %v", err),
			"success": false,
		})
	}
	return idParsed, nil
}
