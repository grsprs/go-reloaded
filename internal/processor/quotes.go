// Copyright (c) 2025 Spiros Nikoloudakis
// Licensed under MIT License - see LICENSE file for details

package processor

import "regexp"

// formatQuotes cleans spacing around quotes
func formatQuotes(text string) string {
	// Handle pattern: ' content ' -> 'content'
	re := regexp.MustCompile(`'\s+([^']+?)\s+'`)
	return re.ReplaceAllString(text, "'$1'")
}