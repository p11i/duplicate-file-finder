package finder

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/p11i/duplicate-file-finder/pkg/hash"
)

// FindDuplicates traverses the specified root directory and identifies duplicate files.
func FindDuplicates(rootDir string, hashMethod string) error {
	hashToPath := make(map[string]string)
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Hash the file
		fileHash, err := hash.File(path, hashMethod)
		if err != nil {
			return err
		}

		// Check for duplicates
		if existingPath, exists := hashToPath[fileHash]; exists {
			fmt.Printf("Duplicate found: %s and %s\n", existingPath, path)
		} else {
			hashToPath[fileHash] = path
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
