// Copyright (c) 2025 Spiros Nikoloudakis
// Licensed under MIT License - see LICENSE file for details

package processor

import (
	"regexp"
	"strings"
)

// correctArticles adjusts 'a' and 'an' before appropriate words
func correctArticles(text string) string {
	re := regexp.MustCompile(`\b(a|an)\s+(\w+)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		if len(parts) != 2 {
			return match
		}
		next := strings.ToLower(parts[1])
		
		if shouldUseAn(next) {
			return "an " + parts[1]
		}
		return "a " + parts[1]
	})
}

// shouldUseAn determines if "an" should be used based on phonetic rules
func shouldUseAn(word string) bool {
	if len(word) == 0 {
		return false
	}
	
	first := word[0]
	// Vowel sounds
	if strings.ContainsRune("aeiou", rune(first)) {
		// Consonant sound exceptions for 'u'
		if first == 'u' && len(word) > 1 {
			second := word[1]
			// Only 'un-', 'us-', 'ur-' when they sound like 'yun', 'yus', 'yur'
			if second == 'n' && (strings.HasPrefix(word, "uni") || strings.HasPrefix(word, "unk")) {
				return false
			}
			if second == 's' && strings.HasPrefix(word, "us") {
				return false
			}
		}
		if strings.HasPrefix(word, "on") || strings.HasPrefix(word, "eu") {
			return false
		}
		return true
	}
	
	// Silent H exceptions
	return first == 'h' && len(word) > 1 && strings.ContainsRune("oe", rune(word[1]))
}
