package processor

import (
	"fmt"
	"go-reloaded/internal/fileio"
	"strings"
)

// ProcessFile is the main entry point for text processing
func ProcessFile(inputPath, outputPath string) error {
	content, err := fileio.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read input file: %v", err)
	}

	processedText := ProcessText(string(content))

	err = fileio.WriteFile(outputPath, []byte(processedText))
	if err != nil {
		return fmt.Errorf("failed to write output file: %v", err)
	}

	return nil
}

// ProcessText applies all transformations to the input text
func ProcessText(text string) string {
	result := text

	// 1️⃣ Numeric conversions
	result = applyNumberConversions(result)

	// 2️⃣ Case transformations (now handles quoted text)
	result = applyCaseTransformations(result)

	// 3️⃣ Article corrections (a → an)
	result = correctArticles(result)

	// 4️⃣ Quote formatting (spacing)
	result = formatQuotes(result)

	// 5️⃣ Punctuation formatting (final cleanup)
	result = formatPunctuation(result)

	return strings.TrimSpace(result)
}
