// Copyright (c) 2025 Spiros Nikoloudakis
// Licensed under MIT License - see LICENSE file for details

package tests

import (
	"go-reloaded/internal/processor"
	"go-reloaded/internal/validator"
	"strings"
	"testing"
)

func TestValidationProtection(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  bool
		errType  string
	}{
		{
			name:    "Normal input",
			input:   "hello world (up)",
			wantErr: false,
		},
		{
			name:    "Unclosed parenthesis",
			input:   "hello (up world",
			wantErr: true,
			errType: "UNCLOSED_BRACKET",
		},
		{
			name:    "Unclosed single quote",
			input:   "hello 'world",
			wantErr: true,
			errType: "UNCLOSED_QUOTE",
		},
		{
			name:    "Unclosed double quote",
			input:   "hello \"world",
			wantErr: true,
			errType: "UNCLOSED_QUOTE",
		},
		{
			name:    "Unmatched closing parenthesis",
			input:   "hello world)",
			wantErr: true,
			errType: "UNMATCHED_BRACKET",
		},
		{
			name:    "Very long line",
			input:   strings.Repeat("a", 100001),
			wantErr: true,
			errType: "LINE_TOO_LONG",
		},
		{
			name:    "Properly closed quotes",
			input:   "hello 'world' and \"test\"",
			wantErr: false,
		},
		{
			name:    "Properly closed parentheses",
			input:   "hello (up) world (low)",
			wantErr: false,
		},
		{
			name:    "Excessive transformations",
			input:   strings.Repeat("test (up) ", 1001),
			wantErr: true,
			errType: "EXCESSIVE_TRANSFORMATIONS",
		},
		{
			name:    "Null byte injection",
			input:   "hello\x00world",
			wantErr: true,
			errType: "MALICIOUS_PATTERN",
		},
		{
			name:    "Excessive character repetition",
			input:   strings.Repeat(" ", 10001),
			wantErr: true,
			errType: "EXCESSIVE_REPETITION",
		},
		{
			name:    "Non-keyboard character",
			input:   "hello 中文 world", // Chinese characters
			wantErr: true,
			errType: "NON_KEYBOARD_CHARACTER",
		},
		{
			name:    "Emoji characters",
			input:   "hello 😀 world", // Emoji
			wantErr: true,
			errType: "NON_KEYBOARD_CHARACTER",
		},
		{
			name:    "Valid keyboard characters",
			input:   "Hello! @#$%^&*()_+-={}[]|\\:;<>?,./ 123",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateInput(tt.input)
			
			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				
				if validationErr, ok := err.(validator.ValidationError); ok {
					if validationErr.Type != tt.errType {
						t.Errorf("Expected error type %s, got %s", tt.errType, validationErr.Type)
					}
				} else {
					t.Errorf("Expected ValidationError, got %T", err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
			}
		})
	}
}

func TestProcessorWithValidation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  bool
	}{
		{
			name:    "Valid input processes normally",
			input:   "hello (up)",
			wantErr: false,
		},
		{
			name:    "Invalid input returns error",
			input:   "hello (up world",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := processor.ProcessText(tt.input)
			
			if tt.wantErr {
				if !strings.HasPrefix(result, "ERROR:") {
					t.Errorf("Expected error output, got: %s", result)
				}
			} else {
				if strings.HasPrefix(result, "ERROR:") {
					t.Errorf("Expected normal output, got error: %s", result)
				}
			}
		})
	}
}

func TestSafeBoundsChecking(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		start    int
		end      int
		expected string
	}{
		{
			name:     "Normal slice",
			str:      "hello",
			start:    1,
			end:      4,
			expected: "ell",
		},
		{
			name:     "Start out of bounds (negative)",
			str:      "hello",
			start:    -1,
			end:      3,
			expected: "hel",
		},
		{
			name:     "End out of bounds",
			str:      "hello",
			start:    2,
			end:      10,
			expected: "llo",
		},
		{
			name:     "Start >= end",
			str:      "hello",
			start:    3,
			end:      2,
			expected: "",
		},
		{
			name:     "Empty string",
			str:      "",
			start:    0,
			end:      1,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validator.SafeSlice(tt.str, tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("SafeSlice(%q, %d, %d) = %q, want %q", 
					tt.str, tt.start, tt.end, result, tt.expected)
			}
		})
	}
}

func TestSafeIndexing(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		index    int
		wantRune rune
		wantOk   bool
	}{
		{
			name:     "Valid index",
			str:      "hello",
			index:    1,
			wantRune: 'e',
			wantOk:   true,
		},
		{
			name:     "Negative index",
			str:      "hello",
			index:    -1,
			wantRune: 0,
			wantOk:   false,
		},
		{
			name:     "Index out of bounds",
			str:      "hello",
			index:    10,
			wantRune: 0,
			wantOk:   false,
		},
		{
			name:     "Empty string",
			str:      "",
			index:    0,
			wantRune: 0,
			wantOk:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRune, gotOk := validator.SafeIndex(tt.str, tt.index)
			if gotRune != tt.wantRune || gotOk != tt.wantOk {
				t.Errorf("SafeIndex(%q, %d) = (%c, %v), want (%c, %v)", 
					tt.str, tt.index, gotRune, gotOk, tt.wantRune, tt.wantOk)
			}
		})
	}
}