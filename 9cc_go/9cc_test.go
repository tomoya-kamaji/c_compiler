package main

import (
	"fmt"
	"testing"
)

func TestExpr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "3+4*10",
			input:    "3+4*10",
			expected: 43,
		},
		{
			name:     "20*  10 - 9 * 3",
			input:    "20*  10 - 9 * 3",
			expected: 173,
		},
		{
			name:     "2*(10-3)*3",
			input:    "2*(10-3)*3",
			expected: 42,
		},
	}
	for _, td := range tests {
		td := td
		t.Run(fmt.Sprintf("tokenize: %s", td.name), func(t *testing.T) {
			tokens = tokenize(td.input)
			result := expr()
			if result != td.expected {
				t.Errorf("parse(%s) = %d, want %d", td.input, result, td.expected)
			}
			tokens = nil
		})
	}
}
