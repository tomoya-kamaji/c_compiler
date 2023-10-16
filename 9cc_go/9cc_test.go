package main

import (
	"fmt"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Token
	}{
		{
			name:  "3+4*10",
			input: "3+4*10",
			expected: []Token{
				{TOKEN_INT, "3"},
				{TOKEN_PLUS, "+"},
				{TOKEN_INT, "4"},
				{TOKEN_MUL, "*"},
				{TOKEN_INT, "10"},
				{TOKEN_EOF, ""},
			},
		},
	}
	for _, td := range tests {
		td := td
		t.Run(fmt.Sprintf("tokenize: %s", td.name), func(t *testing.T) {
			result := tokenize(td.input)
			if !slicesEqual(result, td.expected) {
				t.Errorf("tokenize(%s) = nil", td.input)
			}

		})
	}
}

func TestParse(t *testing.T) {
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
			name:     "20*  10 * 3",
			input:    "20*  10 - 9 * 3",
			expected: 173,
		},
	}
	for _, td := range tests {
		td := td
		t.Run(fmt.Sprintf("tokenize: %s", td.name), func(t *testing.T) {
			tokens := tokenize(td.input)
			parser := Parser{tokens, 0}
			result := parser.parseExpr()
			if result != td.expected {
				t.Errorf("parse(%s) = %d, want %d", td.input, result, td.expected)
			}

		})
	}
}

func slicesEqual(a, b []Token) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
