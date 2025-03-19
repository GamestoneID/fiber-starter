package utils

import (
	"strconv"

	"github.com/gosimple/slug"
)

// Convert string ID to uint safely
func ParseID(id string) (uint, error) {
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(uintID), nil
}

// GenerateSlug generates a slug from a text
func GenerateSlug(text string) string {
	return slug.Make(text)
}
