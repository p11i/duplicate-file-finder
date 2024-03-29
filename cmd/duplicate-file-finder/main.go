package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/p11i/duplicate-file-finder/internal/finder"
)

func main() {
	// Parse command-line arguments
	rootDir := flag.String("dir", ".", "Root directory to search for duplicate files")
	hashMethod := flag.String("hash", "md5", "Hashing method (md5 or sha256)")

	flag.Parse()

	// Check if the specified directory exists
	stat, err := os.Stat(*rootDir)
	if os.IsNotExist(err) {
		fmt.Printf("Error: Directory '%s' does not exist\n", *rootDir)
		os.Exit(1)
	} else if err != nil {
		fmt.Printf("Error checking directory '%s': %v\n", *rootDir, err)
		os.Exit(1)
	}

	// Check if the specified path is a directory
	if !stat.IsDir() {
		fmt.Printf("Error: '%s' is not a valid directory\n", *rootDir)
		os.Exit(1)
	}
	// Check if the specified hash method is supported
	if *hashMethod != "md5" && *hashMethod != "sha256" {
		fmt.Println("Error: Unsupported hashing method. Choose md5 or sha256.")
		os.Exit(1)
	}

	err = finder.FindDuplicates(*rootDir, *hashMethod)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
