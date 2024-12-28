package pvglob

const (
	charWildcard     = '*'
	charQuestionMark = '?'
	charEscape       = '\\'
)

type tokenType int

const (
	LITERAL tokenType = iota
	WILDCARD
	QUESTIONMARK
	EOF
)

type lexer struct {
	buf string
	pos int
}

type token struct {
	name  tokenType
	value string
	pos   int
}

func newLexer(input string) *lexer {
	return &lexer{input, 0}
}

func (l *lexer) lex() []token {
	tokens := []token{}
	for {
		tok := l.nextToken()
		tokens = append(tokens, tok)
		if tok.name == EOF {
			break
		}
	}

	return tokens
}

func (l *lexer) nextToken() token {
	lpos := l.pos

	if lpos >= len(l.buf) {
		return token{EOF, "", l.pos}
	}

	switch l.buf[l.pos] {
	case charWildcard:
		l.pos++
		return token{WILDCARD, "", lpos}
	case charQuestionMark:
		l.pos++
		return token{QUESTIONMARK, "", lpos}
	default:
		return l.readLiteral()
	}
}

func (l *lexer) readLiteral() token {
	start := l.pos

	// Advance the position until we find a wildcard or a question mark.
	l.scanLiteral()

	// Remove escape characters from the literal.
	buf := l.buf[start:l.pos]
	for i := 0; i+1 < len(buf); i++ {
		if buf[i] == charEscape && (buf[i+1] == charWildcard || buf[i+1] == charQuestionMark || buf[i+1] == charEscape) {
			buf = buf[:i] + buf[i+1:]
		}
	}

	return token{LITERAL, buf, start}
}

func (l *lexer) scanLiteral() {
	var escaped bool

	for l.pos < len(l.buf) {
		char := l.buf[l.pos]

		switch char {

		// If the character is an escape character, mark the next character as escaped
		// unless it is already escaped, in that case we go `\\` so we don't escape the next character.
		case charEscape:
			escaped = !escaped

			// If the character is a wildcard or a question mark and it is not escaped, we are done scanning the literal.
		case charWildcard, charQuestionMark:
			if !escaped {
				return
			}

			// If it is escaped, we continue scanning the literal.
			escaped = false
		}

		l.pos++
	}
}
