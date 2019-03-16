package token

// TokenType is an enum of strings.
// Should consider changing this to an int or byte for better performance
type TokenType string

type SourceLocation struct {
	filePath string
	line     int
	col      int
}

type Token struct {
	SourceLocation
	Type    TokenType
	Literal string
}

const (
	// Boundaries
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
