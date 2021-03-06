package token

type TokenType string

type Token struct {
  Type      TokenType
  Literal   string
  Filename  string
  Line      int
  Character int
}

const (
  ILLEGAL = "ILLEGAL"
  EOF     = "EOF"

  // Identifiers + literals
  IDENT = "IDENT"
  INT   = "INT"

  // Operators
  ASSIGN   = "="
  PLUS     = "+"
  MINUS    = "-"
  BANG     = "!"
  ASTERISK = "*"
  SLASH    = "/"

  LT = "<"
  GT = ">"

  EQ     = "=="
  NOT_EQ = "!="

  // Delimiters
  QUOTE        = "\""
  SINGLE_QUOTE = "'"

  COMMA     = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  // Keywords
  FUNCTION = "FUNCTION"
  LET      = "LET"
  TRUE     = "TRUE"
  FALSE    = "FALSE"
  IF       = "IF"
  ELSE     = "ELSE"
  RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
  "fn":     FUNCTION,
  "let":    LET,
  "true":   TRUE,
  "false":  FALSE,
  "if":     IF,
  "else":   ELSE,
  "return": RETURN,
}

func LookupIdent(ident string) TokenType {
  if typ, ok := keywords[ident]; ok {
    return typ
  }
  return IDENT
}
