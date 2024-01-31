package hash

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// File computes the MD5 hash of the specified file.
func File(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := md5.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
