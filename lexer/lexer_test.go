package lexer

import (
  "testing"

  "monkey/token"
)

type ExpectedTokens struct {
    expectedType    token.TokenType
    expectedLiteral string
}

func TestNextToken(t *testing.T) {
  input := `=+(){},;`

  expectations := []ExpectedTokens {
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

  testTokens(t, input, expectations)
}

func testTokens(t *testing.T, input string, tests []ExpectedTokens) {
  l := NewLexer(input)

  for i, tt := range tests {
    tok := l.NextToken()

    if tok.Type != tt.expectedType {
      t.Fatalf("tests[%d] - tokentype wrong. expected %q, got %q", i, tt.expectedType, tok.Type)
    }

    if tok.Literal != tt.expectedLiteral {
      t.Fatalf("tests[%d] - literal wrong. expected %q, got %q", i, tt.expectedLiteral, tok.Literal)
    }
  }
}
