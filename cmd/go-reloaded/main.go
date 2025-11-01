package main

import (
	"fmt"
	"go-reloaded/internal/processor"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Go Reloaded - Text Transformation Tool")
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Println("  CLI:  go run . <input-file> <output-file>")
		fmt.Println("  Web:  go run ./cmd/go-reloaded-web")
		os.Exit(1)
	}

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input-file> <output-file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	err := processor.ProcessFile(inputFile, outputFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully processed %s → %s\n", inputFile, outputFile)
}
