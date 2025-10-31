// Copyright (c) 2025 Spiros Nikoloudakis
// Licensed under MIT License - see LICENSE file for details

package processor

import (
	"regexp"
	"strconv"
)

// applyNumberConversions dynamically converts (hex) and (bin) notations to decimal
func applyNumberConversions(text string) string {
	// Regex for hexadecimal numbers, e.g. 1F (hex)
	hexRe := regexp.MustCompile(`\b([0-9A-Fa-f]+)\s*\(hex\)`)

	// Regex for binary numbers, e.g. 1010 (bin)
	binRe := regexp.MustCompile(`\b([01]+)\s*\(bin\)`)

	// Replace all (hex)
	text = hexRe.ReplaceAllStringFunc(text, func(match string) string {
		sub := hexRe.FindStringSubmatch(match)
		if len(sub) < 2 {
			return match
		}
		val, err := strconv.ParseInt(sub[1], 16, 64)
		if err != nil {
			return match
		}
		return strconv.FormatInt(val, 10)
	})

	// Replace all (bin)
	text = binRe.ReplaceAllStringFunc(text, func(match string) string {
		sub := binRe.FindStringSubmatch(match)
		if len(sub) < 2 {
			return match
		}
		val, err := strconv.ParseInt(sub[1], 2, 64)
		if err != nil {
			return match
		}
		return strconv.FormatInt(val, 10)
	})

	return text
}
