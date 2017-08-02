package token

type TokenType string

type Token struct {
  Type    TokenType
  Literal string
}

const (
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

var keywords = map[string]TokenType{
  "fn": FUNCTION,
  "let": LET,
}

func LookupIdent(ident string) TokenType {
  if typ, ok := keywords[ident]; ok {
    return typ
  }
  return IDENT
}
