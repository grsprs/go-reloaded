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
	
	return processTextCore(text)
}

// ProcessTextUnsafe processes text without validation (for intentional errors)
func ProcessTextUnsafe(text string) string {
	return processTextCore(text)
}

// ProcessTextWithInfo processes text and includes transformation info for web UI
func ProcessTextWithInfo(text string) string {
	// Validate input for security and correctness
	if err := validator.ValidateInput(text); err != nil {
		return "ERROR: " + err.Error()
	}
	
	return processTextInternal(text)
}

// ProcessTextUnsafeWithInfo processes text without validation and includes info
func ProcessTextUnsafeWithInfo(text string) string {
	return processTextInternal(text)
}

// processTextCore performs core text processing without info messages
func processTextCore(text string) string {
	result := text

	// 1️⃣ Numeric conversions
	result = applyNumberConversions(result)

	// 2️⃣ Case transformations
	result = applyCaseTransformations(result)

	// 3️⃣ Article corrections
	result = correctArticles(result)

	// 4️⃣ Quote formatting
	result = formatQuotes(result)

	// 5️⃣ Punctuation formatting
	result = formatPunctuation(result)

	// Clear tracking data without using it
	getAndClearArticleCorrections()
	getAndClearCaseCorrections()
	getAndClearPunctuationCorrections()
	getAndClearQuoteCorrections()

	return strings.TrimSpace(result)
}

// processTextInternal performs text processing with info messages
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

	// Check for corrections and append info
	articleCorrections := getAndClearArticleCorrections()
	caseCorrections := getAndClearCaseCorrections()
	punctuationCorrections := getAndClearPunctuationCorrections()
	quoteCorrections := getAndClearQuoteCorrections()
	
	if len(articleCorrections) > 0 || len(caseCorrections) > 0 || len(punctuationCorrections) > 0 || len(quoteCorrections) > 0 {
		result += "\n\nINFO: Transformations applied:\n"
		
		for _, correction := range articleCorrections {
			result += fmt.Sprintf("• Article: '%s' → '%s'\n", correction.Original, correction.Corrected)
		}
		
		for _, correction := range caseCorrections {
			result += fmt.Sprintf("• Case: %s\n", correction)
		}
		
		for _, correction := range punctuationCorrections {
			result += fmt.Sprintf("• Punctuation: %s\n", correction)
		}
		
		for _, correction := range quoteCorrections {
			result += fmt.Sprintf("• Quotes: %s\n", correction)
		}
	}

	return strings.TrimSpace(result)
}
