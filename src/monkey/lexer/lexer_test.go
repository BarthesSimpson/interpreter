package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextTokenBasicSourceCode(t *testing.T) {
	runTest(t, New(input0), expected0)
}

func TestNextTokenRealisticSourceCode(t *testing.T) {
	runTest(t, New(input1), expected1)
}

func runTest(t *testing.T, lexer *Lexer, tests []TokenExpectation) {
	for i, tt := range tests {
		tok := lexer.NextToken()
		checkToken(i, tok, tt, t)
	}
}

func checkToken(testIndex int, tok token.Token, tt TokenExpectation, t *testing.T) {
	if tok.Type != tt.expectedType {
		t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
			testIndex, tt.expectedType, tok.Type)
	}

	if tok.Literal != tt.expectedLiteral {
		t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
			testIndex, tt.expectedLiteral, tok.Literal)
	}
}
