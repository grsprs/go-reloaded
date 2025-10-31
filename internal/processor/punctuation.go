// Copyright (c) 2025 Spiros Nikoloudakis
// Licensed under MIT License - see LICENSE file for details

package processor

import (
	"regexp"
	"strings"
)

// formatPunctuation ensures punctuation spacing consistency
func formatPunctuation(text string) string {
	// Remove spaces before punctuation marks
	re := regexp.MustCompile(`\s+([,.!?;:])`)
	result := re.ReplaceAllString(text, "$1")
	
	// Handle ellipsis
	result = strings.ReplaceAll(result, " ...", "...")
	
	// Clean up multiple spaces
	result = regexp.MustCompile(`\s+`).ReplaceAllString(result, " ")
	
	return strings.TrimSpace(result)
}
