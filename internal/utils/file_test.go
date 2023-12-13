package utils

import (
	"strings"
	"testing"
)

func TestGenerateFileName(t *testing.T) {
	originalFileName := "testfile.txt"

	generatedName := GenerateFileName(originalFileName)

	// Check if the generated filename contains the original filename
	if !strings.Contains(generatedName, originalFileName) {
		t.Errorf("Generated filename does not contain the original filename. Got: %s", generatedName)
	}

	// Check if a UUID has been generated as a prefix
	parts := strings.Split(generatedName, "_")
	if len(parts) != 2 {
		t.Errorf("Generated filename format is not correct. Expected UUID_originalFilename. Got: %s", generatedName)
	} else {
		// Additional checks can be performed here, such as checking the format of the UUID
		// However, validating the exact UUID is not necessary as it's generated randomly
	}
}
