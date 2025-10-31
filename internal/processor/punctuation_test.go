package processor

import (
	"testing"
)

func TestPunctuationFormatting(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Single punctuation
		{"Hello , world", "Hello, world"},
		{"Wait ! really", "Wait! really"},
		{"How are you ?", "How are you?"},
		{"Test : example", "Test: example"},
		{"This ; that", "This; that"},

		// Punctuation groups
		{"Wait ... really", "Wait... really"},
		{"Really !?", "Really!?"},
		{"Amazing !!", "Amazing!!"},

		// Complex cases
		{"Hello , world ... how are you ?", "Hello, world... how are you?"},
		{"Wait ... really ! ? , ; : all together", "Wait... really!?,;: all together"},
	}

	for _, tc := range tests {
		result := formatPunctuation(tc.input)
		if result != tc.expected {
			t.Errorf("formatPunctuation(%q) = %q, want %q", tc.input, result, tc.expected)
		}
	}
}
