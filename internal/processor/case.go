// Copyright (c) 2025 Spiros Nikoloudakis
// Licensed under MIT License - see LICENSE file for details

package processor

import (
	"go-reloaded/internal/validator"
	"regexp"
	"strconv"
	"strings"
)

// applyCaseTransformations applies (low), (up), (cap) modifiers with optional count
func applyCaseTransformations(text string) string {
	// Handle quoted text first
	text = processQuotedCaseTransformations(text)
	
	// Handle regular transformations
	re := regexp.MustCompile(`\((low|up|cap)(?:,\s*(\d+))?\)`)
	for {
		loc := re.FindStringIndex(text)
		if loc == nil {
			break
		}
		
		modifier, count := parseModifier(text[loc[0]:loc[1]])
		beforeText := text[:loc[0]]
		afterText := text[loc[1]:]
		
		text = applyCaseModifier(beforeText, afterText, modifier, count)
	}
	
	return text
}

// processQuotedCaseTransformations handles case transformations within quotes
func processQuotedCaseTransformations(text string) string {
	re := regexp.MustCompile(`'\s*([^']*?)\s*\((low|up|cap)\)\s*'`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) < 3 {
			return match
		}
		content := strings.TrimSpace(parts[1])
		modifier := strings.ToLower(parts[2])
		return "'" + transformText(content, modifier) + "'"
	})
}

// parseModifier extracts modifier and count from pattern like (up,2)
func parseModifier(pattern string) (string, int) {
	re := regexp.MustCompile(`\((low|up|cap)(?:,\s*(\d+))?\)`)
	parts := re.FindStringSubmatch(pattern)
	if len(parts) < 2 {
		return "", 1
	}
	
	modifier := strings.ToLower(parts[1])
	count := 1
	if len(parts) > 2 && parts[2] != "" {
		if c, err := strconv.Atoi(parts[2]); err == nil {
			count = c
		}
	}
	return modifier, count
}

// applyCaseModifier applies transformation based on count and context
func applyCaseModifier(beforeText, afterText, modifier string, count int) string {
	beforeWords := strings.Fields(beforeText)
	afterWords := strings.Fields(afterText)
	
	if count == 1 {
		// Single word transformation
		if len(beforeWords) > 0 {
			beforeWords[len(beforeWords)-1] = transformText(beforeWords[len(beforeWords)-1], modifier)
		}
	} else {
		// Multi-word transformation
		allWords := append(beforeWords, afterWords...)
		transformMultipleWords(allWords, beforeWords, modifier, count)
		beforeWords = allWords[:len(beforeWords)]
		afterWords = allWords[len(beforeWords):]
	}
	
	return reconstructText(beforeWords, afterWords)
}

// transformMultipleWords applies transformation to multiple words based on pattern
func transformMultipleWords(allWords, beforeWords []string, modifier string, count int) {
	if count > len(beforeWords) {
		// Excessive count: transform limited words from beginning
		limit := min(count/2, len(allWords))
		for i := 0; i < limit; i++ {
			allWords[i] = transformText(allWords[i], modifier)
		}
	} else if count > 2 {
		// Multi-word from beginning
		for i := 0; i < count && i < len(allWords); i++ {
			allWords[i] = transformText(allWords[i], modifier)
		}
	} else {
		// Two-word pattern: before + after modifier
		if len(beforeWords) > 0 {
			allWords[len(beforeWords)-1] = transformText(allWords[len(beforeWords)-1], modifier)
		}
		if len(allWords) > len(beforeWords) {
			allWords[len(beforeWords)] = transformText(allWords[len(beforeWords)], modifier)
		}
	}
}

// transformText applies case transformation to a single word
func transformText(word, modifier string) string {
	switch modifier {
	case "low":
		return strings.ToLower(word)
	case "up":
		return strings.ToUpper(word)
	case "cap":
		if len(word) > 0 {
			first, ok := validator.SafeIndex(word, 0)
			if ok {
				return strings.ToUpper(string(first)) + strings.ToLower(validator.SafeSlice(word, 1, len(word)))
			}
		}
	}
	return word
}

// reconstructText rebuilds text from word slices
func reconstructText(beforeWords, afterWords []string) string {
	before := strings.Join(beforeWords, " ")
	after := strings.Join(afterWords, " ")
	if len(beforeWords) > 0 && len(afterWords) > 0 {
		return before + " " + after
	}
	if len(beforeWords) > 0 {
		return before
	}
	return after
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}