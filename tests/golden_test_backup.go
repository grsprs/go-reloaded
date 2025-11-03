// BACKUP of original golden tests - can be restored if needed
package tests

import (
	"go-reloaded/internal/processor"
	"testing"
)

func TestGoldenTestsBackup(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test 3 - Multi-Word Case Operations (OLD BEHAVIOR)",
			input:    "transform these three words (cap, 3) properly",
			expected: "Transform These Three words properly", // OLD: from beginning
		},
		{
			name:     "Test 8 - Excessive Word Count (OLD BEHAVIOR)",
			input:    "only two words here (up, 5) total",
			expected: "ONLY TWO words here total", // OLD: limited from beginning
		},
		{
			name:     "Test 12 - Comprehensive Example (OLD BEHAVIOR)",
			input:    "Start with 1F (hex) and apply (up, 2) transformations. Format ' this quote ' and handle a hour correctly. Don't forget punctuation , like this ... should work !",
			expected: "Start with 31 and APPLY TRANSFORMATIONS. Format 'this quote' and handle an hour correctly. Don't forget punctuation, like this... should work!", // OLD: from beginning
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := processor.ProcessText(tc.input)
			if result != tc.expected {
				t.Errorf("\nInput:    %q\nExpected: %q\nGot:      %q", tc.input, tc.expected, result)
			}
		})
	}
}