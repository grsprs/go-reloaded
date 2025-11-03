// Copyright (c) 2025 Spiros Nikoloudakis
// Licensed under MIT License - see LICENSE file for details

package processor

import "regexp"

// Global variable to track quote corrections
var quoteCorrections []string

func addQuoteCorrection(correction string) {
	quoteCorrections = append(quoteCorrections, correction)
}

func getAndClearQuoteCorrections() []string {
	corrections := quoteCorrections
	quoteCorrections = nil
	return corrections
}

// formatQuotes cleans spacing around quotes and all paired characters
func formatQuotes(text string) string {
	original := text
	result := text
	
	// Process single quotes (original behavior)
	re := regexp.MustCompile(`'\s+([^']+?)\s+'`)
	result = re.ReplaceAllString(result, "'$1'")
	
	// Process double quotes
	re = regexp.MustCompile(`"\s+([^"]+?)\s+"`)
	result = re.ReplaceAllString(result, `"$1"`)
	
	// Process parentheses
	re = regexp.MustCompile(`\(\s+([^)]+?)\s+\)`)
	result = re.ReplaceAllString(result, "($1)")
	
	// Process square brackets
	re = regexp.MustCompile(`\[\s+([^\]]+?)\s+\]`)
	result = re.ReplaceAllString(result, "[$1]")
	
	// Process curly braces
	re = regexp.MustCompile(`\{\s+([^}]+?)\s+\}`)
	result = re.ReplaceAllString(result, "{$1}")
	
	// Process angle brackets
	re = regexp.MustCompile(`<\s+([^>]+?)\s+>`)
	result = re.ReplaceAllString(result, "<$1>")
	
	// Track if any quote corrections were made
	if original != result {
		addQuoteCorrection("Applied quote and bracket formatting (spacing normalization)")
	}
	
	return result
}