package pvglob

// Parsed represents a parsed pattern that can be used to match strings.
type Parsed []token

func parse(input string) Parsed {
	tokens := newLexer(input).lex()
	tokens = optimizeConsequtiveMatchers(tokens)

	return tokens
}

// LiteralNode represents a single literal in the pattern.
type LiteralNode struct {
	Value      string
	FirstToken bool // True if this is the first token in the pattern.
	LastToken  bool // True if this is the last token in the pattern.
}

// Literals returns all literals in the pattern.
func (pp Parsed) Literals() []LiteralNode {
	literals := make([]LiteralNode, 0, len(pp))
	for i, token := range pp {
		if token.name != LITERAL {
			continue
		}

		literals = append(literals, LiteralNode{
			Value:      token.value,
			FirstToken: i == 0,
			LastToken:  i == len(pp)-2, // Last token is EOF.
		})
	}

	return literals
}

// TrigramsCount returns the number of trigrams of non-wildcard characters in the pattern.
func (pp Parsed) TrigramsCount() int {
	var trigramsCount int
	for _, token := range pp {
		if token.name != LITERAL {
			continue
		}

		// Not a trigram.
		if len(token.value) < 3 {
			continue
		}

		// Count trigrams.
		trigramsCount += len(token.value) - 2
	}

	return trigramsCount
}

// Match returns true if the pattern matches the value.
func (pp Parsed) Match(value string) bool {
	if len(pp) == 0 {
		return true
	}

	if pp[len(pp)-1].name == EOF {
		pp = pp[:len(pp)-1]
	}

	if len(pp) == 1 && pp[0].name == WILDCARD {
		return true
	}

	if len(pp) == 1 && pp[0].name == LITERAL {
		return pp[0].value == value
	}

	if pp[0].name == QUESTIONMARK {
		if len(value) == 0 {
			return false
		}

		return pp[1:].Match(value[1:])
	}

	if pp[0].name == LITERAL {
		if len(pp[0].value) > len(value) {
			return false
		}

		if pp[0].value != value[:len(pp[0].value)] {
			return false
		}

		return pp[1:].Match(value[len(pp[0].value):])
	}

	if pp[0].name == WILDCARD {
		for i := 0; i <= len(value); i++ {
			if pp[1:].Match(value[i:]) {
				return true
			}
		}
	}

	return false
}

func optimizeConsequtiveMatchers(tokens []token) []token {
	var optimizedTokens []token
	for i := 0; i < len(tokens); i++ {
		// If we have a `**` or `*?`  combination, we optimize it to a single `*`.
		if i+1 < len(tokens) && tokens[i].name == WILDCARD && (tokens[i+1].name == WILDCARD || tokens[i+1].name == QUESTIONMARK) {
			optimizedTokens = append(optimizedTokens, token{WILDCARD, "", tokens[i].pos})
			i++
		} else {
			optimizedTokens = append(optimizedTokens, tokens[i])
		}
	}

	return optimizedTokens
}
