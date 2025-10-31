package tests

import (
	"go-reloaded/internal/processor"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGoldenTests(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test 1 - Number System Conversions",
			input:    "Add 1A (hex) and 1010 (bin) to get the result.",
			expected: "Add 26 and 10 to get the result.",
		},
		{
			name:     "Test 2 - Basic Case Transformations",
			input:    "make this IMPORTANT (low) and that small (up)",
			expected: "make this important and that SMALL",
		},
		{
			name:     "Test 3 - Multi-Word Case Operations",
			input:    "transform these three words (cap, 3) properly",
			expected: "Transform These Three words properly",
		},
		{
			name:     "Test 4 - Punctuation Formatting",
			input:    "Hello , world ... how are you today ?",
			expected: "Hello, world... how are you today?",
		},
		{
			name:     "Test 5 - Quote Cleaning",
			input:    "He said ' I am ready ' and left ' immediately '",
			expected: "He said 'I am ready' and left 'immediately'",
		},
		{
			name:     "Test 6 - Article Correction",
			input:    "I saw a apple and a hour later a university.",
			expected: "I saw an apple and an hour later a university.",
		},
		{
			name:     "Test 7 - Invalid Number Handling",
			input:    "Convert GH (hex) and 23 (bin) should remain.",
			expected: "Convert GH (hex) and 23 (bin) should remain.",
		},
		{
			name:     "Test 8 - Excessive Word Count",
			input:    "only two words here (up, 5) total",
			expected: "ONLY TWO words here total",
		},
		{
			name:     "Test 9 - Complex Punctuation",
			input:    "Wait ... really ! ? , ; : all together !!",
			expected: "Wait... really!?,;: all together!!",
		},
		{
			name:     "Test 10 - Nested Transformations",
			input:    "1F (hex) becomes uppercase number",
			expected: "31 becomes uppercase number",
		},
		{
			name:     "Test 11 - Mixed Transformations",
			input:    "' QUOTED TEXT (low) ' needs formatting , correctly !",
			expected: "'quoted text' needs formatting, correctly!",
		},
		{
			name:     "Test 12 - Comprehensive Example",
			input:    "Start with 1F (hex) and apply (up, 2) transformations. Format ' this quote ' and handle a hour correctly. Don't forget punctuation , like this ... should work !",
			expected: "Start with 31 and APPLY TRANSFORMATIONS. Format 'this quote' and handle an hour correctly. Don't forget punctuation, like this... should work!",
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

func TestGoldenFileBased(t *testing.T) {
	inputDir := "input_samples"
	expectedDir := "expected_outputs"
	
	// Test the UTF-8 files we created
	testFiles := []string{"test1.txt", "test2.txt", "test3.txt"}
	
	for _, filename := range testFiles {
		t.Run(filename, func(t *testing.T) {
			inputPath := filepath.Join(inputDir, filename)
			expectedPath := filepath.Join(expectedDir, filename)
			
			// Read input
			inputData, err := os.ReadFile(inputPath)
			if err != nil {
				t.Fatalf("Failed to read input file %s: %v", inputPath, err)
			}
			
			// Read expected output
			expectedData, err := os.ReadFile(expectedPath)
			if err != nil {
				t.Fatalf("Failed to read expected file %s: %v", expectedPath, err)
			}
			
			// Process input
			result := processor.ProcessText(string(inputData))
			expected := strings.TrimSpace(string(expectedData))
			
			if result != expected {
				t.Errorf("\nFile: %s\nExpected: %q\nGot:      %q", filename, expected, result)
			}
		})
	}
}