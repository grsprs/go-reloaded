// Copyright (c) 2025 Spiros Nikoloudakis
// Licensed under MIT License - see LICENSE file for details

package processor

import (
	"regexp"
	"strings"
)

// correctArticles adjusts 'a' and 'an' before appropriate words
func correctArticles(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		word := strings.ToLower(words[i])
		next := strings.ToLower(words[i+1])

		if word == "a" || word == "an" {
			// Rules for when to use "an"
			useAn := false

			// If starts with vowel sound
			if matched, _ := regexp.MatchString(`^[aeiou]`, next); matched {
				useAn = true
			}

			// Exceptions (words that start with vowel but sound like consonant)
			exceptionsA := []string{"university", "unicorn", "european", "one", "user", "unit"}
			for _, ex := range exceptionsA {
				if strings.HasPrefix(next, ex) {
					useAn = false
					break
				}
			}

			// Words with silent 'h' that require 'an'
			exceptionsAn := []string{"hour", "honest", "honor", "heir"}
			for _, ex := range exceptionsAn {
				if strings.HasPrefix(next, ex) {
					useAn = true
					break
				}
			}

			if useAn {
				words[i] = "an"
			} else {
				words[i] = "a"
			}
		}
	}

	return strings.Join(words, " ")
}
