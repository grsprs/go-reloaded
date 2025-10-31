package processor

import "testing"

func TestApplyNumberConversions(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Hexadecimal Conversion",
			input:    "Add 1A (hex) and 10 (hex)",
			expected: "Add 26 and 16",
		},
		{
			name:     "Binary Conversion",
			input:    "Add 1010 (bin) and 10 (bin)",
			expected: "Add 10 and 2",
		},
		{
			name:     "Invalid Hexadecimal",
			input:    "Convert GH (hex) should remain",
			expected: "Convert GH (hex) should remain",
		},
		{
			name:     "Invalid Binary",
			input:    "Convert 102 (bin) should remain",
			expected: "Convert 102 (bin) should remain",
		},
		{
			name:     "Mixed Input",
			input:    "Mix 1F (hex) and 101 (bin) and GH (hex)",
			expected: "Mix 31 and 5 and GH (hex)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := applyNumberConversions(tt.input)
			if output != tt.expected {
				t.Errorf("FAILED [%s]\nExpected: %q\nGot:      %q", tt.name, tt.expected, output)
			}
		})
	}
}
