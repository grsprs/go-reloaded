// Copyright (c) 2025 Spiros Nikoloudakis
// Licensed under MIT License - see LICENSE file for details

package processor

import "strings"

// formatPunctuation ensures punctuation spacing consistency
func formatPunctuation(text string) string {
	replacer := strings.NewReplacer(
		" ,", ",",
		" .", ".",
		" !", "!",
		" ?", "?",
		" ;", ";",
		" :", ":",
		" ...", "...",
		"!!", "!!",
	)
	result := replacer.Replace(text)
	result = strings.ReplaceAll(result, "  ", " ")
	return strings.TrimSpace(result)
}
