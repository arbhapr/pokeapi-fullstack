// model/response.go
package model

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// PaginatedResponse represents the structure of the API response with pagination
type PaginatedResponse struct {
	Data       interface{}    `json:"data"` // Data can be of any type
	Pagination PaginationInfo `json:"pagination"`
}

// PaginationInfo contains pagination details
type PaginationInfo struct {
	NextPage  string `json:"nextPage"`
	PrevPage  string `json:"prevPage"`
	TotalData int    `json:"totalData"`
}

// Pagination holds the parameters for paginating data
type Pagination struct {
	Limit  int // Number of items per page
	Offset int // The starting point for fetching items
}

// ExtractPagination extracts pagination parameters from the request query
func ExtractPagination(c *fiber.Ctx, defaultLimit int) Pagination {
	limit, _ := strconv.Atoi(c.Query("limit", strconv.Itoa(defaultLimit)))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	return Pagination{
		Limit:  limit,
		Offset: offset,
	}
}
