// Copyright (c) 2025 Spiros Nikoloudakis
// Licensed under MIT License - see LICENSE file for details

package validator

import (
	"fmt"
	"strings"
)

const (
	MaxInputSize = 10 * 1024 * 1024 // 10MB limit
	MaxLineLength = 100000          // 100K chars per line
	MaxNestingDepth = 50            // Max nested parentheses
	MaxTransformations = 1000       // Max transformations per input
)

type ValidationError struct {
	Type     string
	Position int
	Message  string
	Context  string
}

func (e ValidationError) Error() string {
	if e.Position > 0 {
		return fmt.Sprintf("%s\n\nFound at position %d in your text.\nPlease check and fix this issue before continuing.", e.Message, e.Position)
	}
	return e.Message
}

// ValidateInput performs comprehensive input validation
func ValidateInput(input string) error {
	// Check buffer overflow protection
	if len(input) > MaxInputSize {
		return ValidationError{
			Type:    "BUFFER_OVERFLOW",
			Message: fmt.Sprintf("Input size %d exceeds maximum allowed %d bytes", len(input), MaxInputSize),
		}
	}

	// Check line length limits
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if len(line) > MaxLineLength {
			return ValidationError{
				Type:     "LINE_TOO_LONG",
				Position: i + 1,
				Message:  fmt.Sprintf("Line %d length %d exceeds maximum %d characters", i+1, len(line), MaxLineLength),
				Context:  truncateString(line, 50),
			}
		}
	}

	// Check for unclosed brackets and quotes
	if err := validateBrackets(input); err != nil {
		return err
	}

	// Check for excessive transformations (DoS protection)
	if err := validateTransformationCount(input); err != nil {
		return err
	}

	// Check for malicious patterns
	if err := validateMaliciousPatterns(input); err != nil {
		return err
	}

	// Check for non-text content
	if err := validateTextContent(input); err != nil {
		return err
	}

	return nil
}

// validateBrackets checks for unclosed parentheses and quotes
func validateBrackets(input string) error {
	var parenStack []int
	var singleQuotes []int
	var doubleQuotes []int
	
	runes := []rune(input)
	
	for i, r := range runes {
		switch r {
		case '(':
			parenStack = append(parenStack, i)
		case ')':
			if len(parenStack) == 0 {
				return ValidationError{
					Type:     "UNMATCHED_BRACKET",
					Position: i,
					Message:  "You have a closing parenthesis ()) without a matching opening parenthesis. Please add the opening parenthesis or remove the extra closing one.",
					
				}
			}
			parenStack = parenStack[:len(parenStack)-1]
		case '\'':
			// Skip contractions (apostrophes between letters)
			if isContraction(runes, i) {
				continue
			}
			if len(singleQuotes) > 0 && singleQuotes[len(singleQuotes)-1] != -1 {
				singleQuotes[len(singleQuotes)-1] = -1 // Mark as closed
			} else {
				singleQuotes = append(singleQuotes, i)
			}
		case '"':
			if len(doubleQuotes) > 0 && doubleQuotes[len(doubleQuotes)-1] != -1 {
				doubleQuotes[len(doubleQuotes)-1] = -1 // Mark as closed
			} else {
				doubleQuotes = append(doubleQuotes, i)
			}
		}
	}

	// Check nesting depth
	if len(parenStack) > MaxNestingDepth {
		return ValidationError{
			Type:    "EXCESSIVE_NESTING",
			Message: fmt.Sprintf("Parentheses nesting depth %d exceeds maximum %d", len(parenStack), MaxNestingDepth),
		}
	}

	// Check for unclosed parentheses
	if len(parenStack) > 0 {
		pos := parenStack[0]
		return ValidationError{
			Type:     "UNCLOSED_BRACKET",
			Position: pos,
			Message:  "You have an unclosed opening parenthesis (() in your text. Please add the closing parenthesis or remove it if not needed.",
			
		}
	}

	// Check for unclosed single quotes
	for _, pos := range singleQuotes {
		if pos != -1 {
			return ValidationError{
				Type:     "UNCLOSED_QUOTE",
				Position: pos,
				Message:  "You have an unclosed single quote (') in your text. Please add the closing quote or remove it if not needed.",
			}
		}
	}

	// Check for unclosed double quotes
	for _, pos := range doubleQuotes {
		if pos != -1 {
			return ValidationError{
				Type:     "UNCLOSED_QUOTE",
				Position: pos,
				Message:  "You have an unclosed double quote (\") in your text. Please add the closing quote or remove it if not needed.",
			}
		}
	}

	return nil
}

// getContext returns surrounding text for error highlighting
func getContext(runes []rune, pos int) string {
	start := max(0, pos-20)
	end := min(len(runes), pos+21)
	
	context := string(runes[start:end])
	
	// Highlight the problem character
	relativePos := pos - start
	if relativePos >= 0 && relativePos < len(context) {
		contextRunes := []rune(context)
		highlighted := string(contextRunes[:relativePos]) + ">>>" + string(contextRunes[relativePos]) + "<<<" + string(contextRunes[relativePos+1:])
		return highlighted
	}
	
	return context
}

// truncateString safely truncates a string to specified length
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

// SafeSlice performs bounds-checked slice operations
func SafeSlice(s string, start, end int) string {
	if start < 0 {
		start = 0
	}
	if end > len(s) {
		end = len(s)
	}
	if start >= end {
		return ""
	}
	return s[start:end]
}

// SafeIndex performs bounds-checked string indexing
func SafeIndex(s string, index int) (rune, bool) {
	runes := []rune(s)
	if index < 0 || index >= len(runes) {
		return 0, false
	}
	return runes[index], true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// isContraction checks if apostrophe is part of a contraction like "don't"
func isContraction(runes []rune, pos int) bool {
	if pos <= 0 || pos >= len(runes)-1 {
		return false
	}
	// Check if surrounded by letters
	return isLetter(runes[pos-1]) && isLetter(runes[pos+1])
}

// isLetter checks if rune is a letter
func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

// isKeyboardCharacter checks if rune is a standard keyboard character
func isKeyboardCharacter(r rune) bool {
	// Allow basic whitespace
	if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
		return true
	}
	
	// Allow printable ASCII (standard keyboard characters)
	if r >= 32 && r <= 126 {
		return true
	}
	
	// Allow common extended ASCII for international keyboards
	if r >= 160 && r <= 255 {
		return true
	}
	
	// Allow common Unicode characters that appear on keyboards
	commonUnicode := []rune{
		'€', '£', '¥', '©', '®', '™', '°', '±', '²', '³',
		'¼', '½', '¾', '×', '÷', 'α', 'β', 'γ', 'δ', 'π',
	}
	
	for _, char := range commonUnicode {
		if r == char {
			return true
		}
	}
	
	return false
}

// validateTransformationCount prevents DoS attacks via excessive transformations
func validateTransformationCount(input string) error {
	count := strings.Count(input, "(hex)") + strings.Count(input, "(bin)") +
		strings.Count(input, "(up)") + strings.Count(input, "(low)") + strings.Count(input, "(cap)")
	
	if count > MaxTransformations {
		return ValidationError{
			Type:    "EXCESSIVE_TRANSFORMATIONS",
			Message: fmt.Sprintf("Input contains %d transformations, maximum allowed is %d", count, MaxTransformations),
		}
	}
	return nil
}

// validateMaliciousPatterns checks for potentially malicious input patterns
func validateMaliciousPatterns(input string) error {
	// Check for null bytes
	if strings.Contains(input, "\x00") {
		return ValidationError{
			Type:    "MALICIOUS_PATTERN",
			Message: "Input contains null bytes which are not allowed",
		}
	}
	
	// Check for excessive repeated characters (potential DoS)
	for _, char := range []string{"(", ")", "'", "\"", " "} {
		if strings.Count(input, char) > 10000 {
			return ValidationError{
				Type:    "EXCESSIVE_REPETITION",
				Message: fmt.Sprintf("Excessive repetition of character '%s' detected", char),
			}
		}
	}
	
	return nil
}

// validateTextContent ensures input contains only valid keyboard characters
func validateTextContent(input string) error {
	// Check for keyboard character compliance
	for i, r := range input {
		if !isKeyboardCharacter(r) {
			return ValidationError{
				Type:     "NON_KEYBOARD_CHARACTER",
				Position: i,
				Message:  fmt.Sprintf("Non-keyboard character detected: '%c' (U+%04X). Please use only standard keyboard characters.", r, r),
				Context:  getContext([]rune(input), i),
			}
		}
	}
	
	// Check for binary content indicators
	nonTextCount := 0
	for _, r := range input {
		// Allow printable ASCII, common Unicode, and whitespace
		if r < 32 && r != '\t' && r != '\n' && r != '\r' {
			nonTextCount++
		}
		// Check for common binary file signatures
		if r > 127 && r < 160 {
			nonTextCount++
		}
	}
	
	// If more than 5% non-text characters, likely binary
	if len(input) > 0 && float64(nonTextCount)/float64(len(input)) > 0.05 {
		return ValidationError{
			Type:    "NON_TEXT_CONTENT",
			Message: "Input appears to contain binary or non-text content. Please paste only plain text.",
		}
	}
	
	// Check for common binary file headers
	binaryHeaders := []string{
		"\x89PNG", "\xFF\xD8\xFF", "GIF8", "\x00\x00\x00\x20ftypmp4",
		"PK\x03\x04", "\x50\x4B", "\x7FELF", "MZ",
	}
	
	for _, header := range binaryHeaders {
		if strings.HasPrefix(input, header) {
			return ValidationError{
				Type:    "BINARY_FILE",
				Message: "Binary file detected. This tool only processes plain text. Please paste text content instead.",
			}
		}
	}
	
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}