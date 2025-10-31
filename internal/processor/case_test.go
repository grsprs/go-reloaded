package processor

import (
	"testing"
)

func TestCaseConversions(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Single word transformations
		{"hello (up)", "HELLO"},
		{"WORLD (low)", "world"},
		{"test (cap)", "Test"},

		// Multiple word transformations
		{"one two (up, 2)", "ONE TWO"},
		{"ONE TWO (low, 2)", "one two"},
		{"first second (cap, 2)", "First Second"},

		// Edge cases
		{"only (up, 5)", "ONLY"},
		{"TWO words (low, 3)", "two words"},
	}

	for _, tc := range tests {
		result := applyCaseTransformations(tc.input)
		if result != tc.expected {
			t.Errorf("applyCaseTransformations(%q) = %q, want %q", tc.input, result, tc.expected)
		}
	}
}
