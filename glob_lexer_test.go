package pvglob

import (
	"reflect"
	"testing"
)

func TestGlobLexer(t *testing.T) {
	for _, tc := range []struct {
		pattern string
		tokens  []token
	}{
		{
			pattern: "abcef",
			tokens: []token{
				{LITERAL, "abcef", 0},
				{EOF, "", 5},
			},
		},
		{
			pattern: "abcef*",
			tokens: []token{
				{LITERAL, "abcef", 0},
				{WILDCARD, "", 5},
				{EOF, "", 6},
			},
		},
		{
			pattern: "*abcef",
			tokens: []token{
				{WILDCARD, "", 0},
				{LITERAL, "abcef", 1},
				{EOF, "", 6},
			},
		},
		{
			pattern: "*abcef*",
			tokens: []token{
				{WILDCARD, "", 0},
				{LITERAL, "abcef", 1},
				{WILDCARD, "", 6},
				{EOF, "", 7},
			},
		},
		{
			pattern: "ab*cdefg*hi",
			tokens: []token{
				{LITERAL, "ab", 0},
				{WILDCARD, "", 2},
				{LITERAL, "cdefg", 3},
				{WILDCARD, "", 8},
				{LITERAL, "hi", 9},
				{EOF, "", 11},
			},
		},
		{
			pattern: "ab?cdefg",
			tokens: []token{
				{LITERAL, "ab", 0},
				{QUESTIONMARK, "", 2},
				{LITERAL, "cdefg", 3},
				{EOF, "", 8},
			},
		},
		{
			pattern: "ab?cdefg*",
			tokens: []token{
				{LITERAL, "ab", 0},
				{QUESTIONMARK, "", 2},
				{LITERAL, "cdefg", 3},
				{WILDCARD, "", 8},
				{EOF, "", 9},
			},
		},
		{
			pattern: "ab\\*cde",
			tokens: []token{
				{LITERAL, "ab*cde", 0},
				{EOF, "", 7},
			},
		},
		{
			pattern: "ab\\?cde",
			tokens: []token{
				{LITERAL, "ab?cde", 0},
				{EOF, "", 7},
			},
		},
		{
			pattern: "ab\\?cde*q",
			tokens: []token{
				{LITERAL, "ab?cde", 0},
				{WILDCARD, "", 7},
				{LITERAL, "q", 8},
				{EOF, "", 9},
			},
		},
		{
			pattern: `a\\b`,
			tokens: []token{
				{LITERAL, `a\b`, 0},
				{EOF, "", 4},
			},
		},
		{
			pattern: `a\\\*b`,
			tokens: []token{
				{LITERAL, `a\*b`, 0},
				{EOF, "", 6},
			},
		},
		{
			pattern: `a\\\\*b`,
			tokens: []token{
				{LITERAL, `a\\`, 0},
				{WILDCARD, "", 5},
				{LITERAL, "b", 6},
				{EOF, "", 7},
			},
		},
		{
			pattern: "john?*",
			tokens: []token{
				{LITERAL, "john", 0},
				{QUESTIONMARK, "", 4},
				{WILDCARD, "", 5},
				{EOF, "", 6},
			},
		},
		{
			pattern: "john?*@gmail.com",
			tokens: []token{
				{LITERAL, "john", 0},
				{QUESTIONMARK, "", 4},
				{WILDCARD, "", 5},
				{LITERAL, "@gmail.com", 6},
				{EOF, "", 16},
			},
		},
		{
			pattern: "john*?",
			tokens: []token{
				{LITERAL, "john", 0},
				{WILDCARD, "", 4},
				{QUESTIONMARK, "", 5},
				{EOF, "", 6},
			},
		},
		{
			pattern: "john*?@gmail.com",
			tokens: []token{
				{LITERAL, "john", 0},
				{WILDCARD, "", 4},
				{QUESTIONMARK, "", 5},
				{LITERAL, "@gmail.com", 6},
				{EOF, "", 16},
			},
		},
		{
			pattern: "",
			tokens: []token{
				{EOF, "", 0},
			},
		},
		{
			pattern: "*",
			tokens: []token{
				{WILDCARD, "", 0},
				{EOF, "", 1},
			},
		},
		{
			pattern: "\\",
			tokens: []token{
				{LITERAL, `\`, 0},
				{EOF, "", 1},
			},
		},
		{
			pattern: "\\*",
			tokens: []token{
				{LITERAL, `*`, 0},
				{EOF, "", 2},
			},
		},
	} {
		t.Run(tc.pattern, func(t *testing.T) {
			lexer := newLexer(tc.pattern)
			tokens := lexer.lex()
			if !reflect.DeepEqual(tc.tokens, tokens) {
				t.Errorf("expected: %v, got: %v", tc.tokens, tokens)
			}
		})
	}
}
