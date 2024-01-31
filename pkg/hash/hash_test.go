package hash

import (
	"os"
	"testing"
)

func TestFileHashing(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "testfile.txt")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.WriteString("test content")
	tempFile.Close()

	// Test MD5 hashing
	md5Hash, err := File(tempFile.Name(), "md5")
	if err != nil {
		t.Errorf("MD5 hashing error: %v", err)
	}
	expectedMD5Hash := "9473fdd0d880a43c21b7778d34872157"
	if md5Hash != expectedMD5Hash {
		t.Errorf("MD5 hashing result incorrect. Expected: %s, Got: %s", expectedMD5Hash, md5Hash)
	}

	// Test SHA-256 hashing
	sha256Hash, err := File(tempFile.Name(), "sha256")
	if err != nil {
		t.Errorf("SHA-256 hashing error: %v", err)
	}
	expectedSHA256Hash := "6ae8a75555209fd6c44157c0aed8016e763ff435a19cf186f76863140143ff72"
	if sha256Hash != expectedSHA256Hash {
		t.Errorf("SHA-256 hashing result incorrect. Expected: %s, Got: %s", expectedSHA256Hash, sha256Hash)
	}
}
