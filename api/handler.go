package api

import (
	"fmt"
	"search-eng/loader"
	"search-eng/search"

	"github.com/gofiber/fiber/v2"
)

type SearchResponse struct {
	Results    []search.Record `json:"results"`
	TotalHits  int             `json:"totalHits"`
	SearchTime string          `json:"searchTime"`
}

type SuccessResponse struct {
	Message      string `json:"message"`
	TotalRecords int    `json:"totalRecords"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// DataHandler handles POST /upload
func DataHandler(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to get uploaded file",
		})
	}

	// Save file temporarily
	tempFilePath := "./temp_uploaded.parquet"
	if err := c.SaveFile(fileHeader, tempFilePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save uploaded file",
		})
	}

	// Load the new data into memory
	records := loader.LoadParquetFile(tempFilePath)
	search.LoadData(records)

	fmt.Printf("addd loaded successfully with respected path")

	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Message:      "File uploaded and data loaded successfully",
		TotalRecords: len(records),
	})
}

// SearchHandler handles GET /search
func SearchHandler(c *fiber.Ctx) error {
	query := c.Query("query")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing 'query' parameter",
		})
	}

	results, elapsed := search.Search(query)

	response := SearchResponse{
		Results:    results,
		TotalHits:  len(results),
		SearchTime: elapsed.String(),
	}

	return c.JSON(response)
}
