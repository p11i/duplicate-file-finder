package finder

import (
	"os"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("Error creating temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create two duplicate files in the temporary directory
	duplicateFile1, err := os.CreateTemp(tempDir, "duplicate1.txt")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}

	defer os.Remove(duplicateFile1.Name())
	duplicateFile2, err := os.CreateTemp(tempDir, "duplicate2.txt")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}

	defer os.Remove(duplicateFile2.Name())
	// Test finding duplicates
	err = FindDuplicates(tempDir, "md5")
	if err != nil {
		t.Errorf("Error finding duplicates: %v", err)
	}
}
