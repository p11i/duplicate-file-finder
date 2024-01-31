package main

import (
	"fmt"
	"os"

	"github.com/p11i/duplicate-file-finder/internal/finder"
)

func main() {
	rootDir := "/path/to/directory"
	err := finder.FindDuplicates(rootDir)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
