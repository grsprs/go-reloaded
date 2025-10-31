// Copyright (c) 2025 Spiros Nikoloudakis
// Licensed under MIT License - see LICENSE file for details

package processor

import (
	"regexp"
	"strconv"
)

// applyNumberConversions converts (hex) and (bin) notations to decimal
func applyNumberConversions(text string) string {
	conversions := map[string]int{
		"hex": 16,
		"bin": 2,
	}
	
	for format, base := range conversions {
		pattern := `\b([0-9A-Fa-f]+)\s*\(` + format + `\)`
		if format == "bin" {
			pattern = `\b([01]+)\s*\(bin\)`
		}
		
		re := regexp.MustCompile(pattern)
		text = re.ReplaceAllStringFunc(text, func(match string) string {
			parts := re.FindStringSubmatch(match)
			if len(parts) < 2 {
				return match
			}
			if val, err := strconv.ParseInt(parts[1], base, 64); err == nil {
				return strconv.FormatInt(val, 10)
			}
			return match
		})
	}
	
	return text
}
