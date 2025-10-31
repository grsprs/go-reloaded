package processor

import (
	"testing"
)

func TestArticleCorrection(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Vowel cases
		{"a apple", "an apple"},
		{"a orange", "an orange"},
		{"a elephant", "an elephant"},
		{"a igloo", "an igloo"},
		{"a umbrella", "an umbrella"},

		// H cases
		{"a hour", "an hour"},
		{"a honest", "an honest"},
		{"a honor", "an honor"},

		// Should NOT change
		{"a university", "a university"},
		{"a unicorn", "a unicorn"},
		{"a house", "a house"},
		{"a hotel", "a hotel"},
		{"a happy", "a happy"},

		// Multiple occurrences
		{"a apple and a hour", "an apple and an hour"},
	}

	for _, tc := range tests {
		result := correctArticles(tc.input)
		if result != tc.expected {
			t.Errorf("correctArticles(%q) = %q, want %q", tc.input, result, tc.expected)
		}
	}
}
