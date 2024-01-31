package hash

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
)

// File computes the MD5 hash of the specified file.
func File(filePath string, method string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var hasher hash.Hash

	switch method {
	case "md5":
		hasher = md5.New()
	case "sha256":
		hasher = sha256.New()
	default:
		return "", fmt.Errorf("unsupported hashing method: %s", method)
	}

	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
