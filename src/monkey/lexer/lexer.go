package lexer

import (
	"monkey/token"
)

// Lexer should also have a version that takes an ioReader and a filename
// for interpreting source code files instead of REPL commands. It should also
// be modified to support unicode (see readChar)
type Lexer struct {
	input        string
	position     int  // current char
	readPosition int  // next char
	ch           byte // current char under examination
}

// New is a factory for creating new Lexer instance per input
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar currently only supports ascii. To support UTF-8 we
// will need to change lexer's char to a rune and change the
// way chars are read since they can now be more than one byte wide
// (and will have variable byte length).
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken reads and returns the next token in the input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
