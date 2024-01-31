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
	flag.Parse()

	// Check if the specified directory exists
	if _, err := os.Stat(*rootDir); os.IsNotExist(err) {
		fmt.Printf("Error: Directory '%s' does not exist\n", *rootDir)
		os.Exit(1)
	}

	err := finder.FindDuplicates(*rootDir)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
