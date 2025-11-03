// Copyright (c) 2025 Spiros Nikoloudakis
// Licensed under MIT License - see LICENSE file for details

package processor

import (
	"regexp"
	"strings"
)

// ArticleCorrection represents a correction made
type ArticleCorrection struct {
	Original string
	Corrected string
	Position int
}

// correctArticles adjusts 'a' and 'an' before appropriate words
func correctArticles(text string) string {
	re := regexp.MustCompile(`\b(a|an)\s+(\w+)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		if len(parts) != 2 {
			return match
		}
		next := strings.ToLower(parts[1])
		
		originalArticle := parts[0]
		correctArticle := "a"
		if shouldUseAn(next) {
			correctArticle = "an"
		}
		
		// Track correction if changed
		if strings.ToLower(originalArticle) != correctArticle {
			addArticleCorrection(originalArticle + " " + parts[1], correctArticle + " " + parts[1])
		}
		
		return correctArticle + " " + parts[1]
	})
}

// Global variable to track corrections
var articleCorrections []ArticleCorrection

func addArticleCorrection(original, corrected string) {
	articleCorrections = append(articleCorrections, ArticleCorrection{
		Original: original,
		Corrected: corrected,
	})
}

func getAndClearArticleCorrections() []ArticleCorrection {
	corrections := articleCorrections
	articleCorrections = nil
	return corrections
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
