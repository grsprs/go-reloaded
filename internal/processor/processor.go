package processor

import (
	"fmt"
	"go-reloaded/internal/fileio"
	"go-reloaded/internal/validator"
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
	// Validate input for security and correctness
	if err := validator.ValidateInput(text); err != nil {
		return "ERROR: " + err.Error()
	}
	
	return processTextInternal(text)
}

// ProcessTextUnsafe processes text without validation (for intentional errors)
func ProcessTextUnsafe(text string) string {
	return processTextInternal(text)
}

// processTextInternal performs the actual text processing
func processTextInternal(text string) string {
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
