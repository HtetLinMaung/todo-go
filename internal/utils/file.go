package utils

import (
	"github.com/google/uuid"
)

func GenerateFileName(originalFileName string) string {
	// Generate a UUID
	id := uuid.New()

	// Concatenate the new file name
	newFileName := id.String() + "_" + originalFileName

	return newFileName
}
