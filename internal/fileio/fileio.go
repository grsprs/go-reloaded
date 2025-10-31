// Copyright (c) 2025 Spiros Nikoloudakis
// Licensed under MIT License - see LICENSE file for details

package fileio

import (
	"fmt"
	"os"
)

// ReadFile reads the entire content of a file
func ReadFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read file %s: %v", path, err)
	}
	return content, nil
}

// WriteFile writes data to a file, creating it if it doesn't exist
func WriteFile(path string, data []byte) error {
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("could not write file %s: %v", path, err)
	}
	return nil
}
