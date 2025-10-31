// Copyright (c) 2025 Spiros Nikoloudakis
// Licensed under MIT License - see LICENSE file for details

package processor

import "strings"

// formatQuotes cleans spacing around quotes
func formatQuotes(text string) string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return text
	}

	var result []string
	inQuote := false
	var quoteBuffer []string

	for _, word := range words {
		if word == "'" {
			if !inQuote {
				inQuote = true
				quoteBuffer = []string{}
			} else {
				if len(quoteBuffer) > 0 {
					result = append(result, "'"+strings.Join(quoteBuffer, " ")+"'")
				} else {
					result = append(result, "''")
				}
				inQuote = false
				quoteBuffer = nil
			}
		} else if strings.HasPrefix(word, "'") && !inQuote {
			inQuote = true
			remaining := strings.TrimPrefix(word, "'")
			if remaining != "" {
				quoteBuffer = append(quoteBuffer, remaining)
			}
		} else if strings.HasSuffix(word, "'") && inQuote {
			remaining := strings.TrimSuffix(word, "'")
			if remaining != "" {
				quoteBuffer = append(quoteBuffer, remaining)
			}
			result = append(result, "'"+strings.Join(quoteBuffer, " ")+"'")
			inQuote = false
			quoteBuffer = nil
		} else if inQuote {
			quoteBuffer = append(quoteBuffer, word)
		} else {
			result = append(result, word)
		}
	}

	if inQuote && len(quoteBuffer) > 0 {
		result = append(result, "'"+strings.Join(quoteBuffer, " "))
	}

	return strings.Join(result, " ")
}
