package lexer

import (
	"monkey/token"
)

// TokenExpectation is a simple data structure
// for checking test output
type TokenExpectation struct {
	expectedType    token.TokenType
	expectedLiteral string
}

var input0 = `=+(){},;`
var input1 = `
let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);
`

var expected0 = []TokenExpectation{
	{token.ASSIGN, "="},
	{token.PLUS, "+"},
	{token.LPAREN, "("},
	{token.RPAREN, ")"},
	{token.LBRACE, "{"},
	{token.RBRACE, "}"},
	{token.COMMA, ","},
	{token.SEMICOLON, ";"},
	{token.EOF, ""},
}

var expected1 = []TokenExpectation{
	{token.LET, "let"},
	{token.IDENT, "five"},
	{token.ASSIGN, "="},
	{token.INT, "5"},
	{token.SEMICOLON, ";"},
	{token.LET, "let"},
	{token.IDENT, "ten"},
	{token.ASSIGN, "="},
	{token.INT, "10"},
	{token.SEMICOLON, ";"},
	{token.LET, "let"},
	{token.IDENT, "add"},
	{token.ASSIGN, "="},
	{token.FUNCTION, "fn"},
	{token.LPAREN, "("},
	{token.IDENT, "x"},
	{token.COMMA, ","},
	{token.IDENT, "y"},
	{token.RPAREN, ")"},
	{token.LBRACE, "{"},
	{token.IDENT, "x"},
	{token.PLUS, "+"},
	{token.IDENT, "y"},
	{token.SEMICOLON, ";"},
	{token.RBRACE, "}"},
	{token.SEMICOLON, ";"},
	{token.LET, "let"},
	{token.IDENT, "result"},
	{token.ASSIGN, "="},
	{token.IDENT, "add"},
	{token.LPAREN, "("},
	{token.IDENT, "five"},
	{token.COMMA, ","},
	{token.IDENT, "ten"},
	{token.RPAREN, ")"},
	{token.SEMICOLON, ";"},
	{token.EOF, ""},
}
