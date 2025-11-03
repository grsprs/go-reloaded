// Copyright (c) 2025 Spiros Nikoloudakis
// Licensed under MIT License - see LICENSE file for details

package processor

import (
	"regexp"
	"strings"
)

// Global variable to track punctuation corrections
var punctuationCorrections []string

func addPunctuationCorrection(correction string) {
	punctuationCorrections = append(punctuationCorrections, correction)
}

func getAndClearPunctuationCorrections() []string {
	corrections := punctuationCorrections
	punctuationCorrections = nil
	return corrections
}

// formatPunctuation ensures punctuation spacing consistency for all punctuation marks
func formatPunctuation(text string) string {
	original := text
	result := text
	
	// Remove spaces before basic punctuation: . , ! ? ; :
	re := regexp.MustCompile(`\s+([,.!?;:])`)
	result = re.ReplaceAllString(result, "$1")
	
	// Remove spaces before extended punctuation
	re = regexp.MustCompile(`\s+([\-–—_~\*\+\=\|\\\/%@#\$&])`)
	result = re.ReplaceAllString(result, "$1")
	
	// Handle ellipsis and multiple dots
	result = strings.ReplaceAll(result, " ...", "...")
	result = strings.ReplaceAll(result, " …", "…")
	
	// Handle multiple punctuation groups
	re = regexp.MustCompile(`\s+([!]{2,})`)
	result = re.ReplaceAllString(result, "$1")
	re = regexp.MustCompile(`\s+([?]{2,})`)
	result = re.ReplaceAllString(result, "$1")
	re = regexp.MustCompile(`\s+([!?]+)`)
	result = re.ReplaceAllString(result, "$1")
	
	// Clean up multiple spaces
	result = regexp.MustCompile(`\s+`).ReplaceAllString(result, " ")
	
	result = strings.TrimSpace(result)
	
	// Track if any punctuation corrections were made
	if original != result {
		addPunctuationCorrection("Applied punctuation formatting (spacing and grouping)")
	}
	
	return result
}
