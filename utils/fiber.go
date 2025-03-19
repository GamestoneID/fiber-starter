package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Utility untuk mendapatkan header opsional sebagai pointer string
func GetOptionalHeader(c *fiber.Ctx, key string) *string {
	value := c.Get(key)
	if value == "" {
		return nil
	}
	return &value
}

// Utility untuk mendapatkan query parameter boolean opsional
func GetOptionalBoolQuery(c *fiber.Ctx, key string) *bool {
	value := c.Query(key)
	if value == "" {
		return nil
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return nil
	}

	return &boolValue
}

// Utility untuk mendapatkan query parameter boolean opsional
func GetOptionalIntQuery(c *fiber.Ctx, key string) *uint {
	value := c.Query(key)
	if value == "" {
		return nil
	}

	uintID, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return nil
	}

	convertedID := uint(uintID) // Konversi dari uint64 ke uint
	return &convertedID
}

// Utility untuk mendapatkan query parameter boolean opsional
func GetOptionalStringQuery(c *fiber.Ctx, key string) *string {
	value := c.Query(key)
	if value == "" {
		return nil
	}

	return &value
}
