package lexer

import (
  "testing"

  "monkey/token"
)

type ExpectedTokens struct {
    expectedType    token.TokenType
    expectedLiteral string
}

func TestNextToken_CanReadBasicSymbols(t *testing.T) {
  input := `=+(){},;-!*/<>`

  expectations := []ExpectedTokens {
    {token.ASSIGN, "="},
    {token.PLUS, "+"},
    {token.LPAREN, "("},
    {token.RPAREN, ")"},
    {token.LBRACE, "{"},
    {token.RBRACE, "}"},
    {token.COMMA, ","},
    {token.SEMICOLON, ";"},
    {token.MINUS, "-"},
    {token.BANG, "!"},
    {token.ASTERISK, "*"},
    {token.SLASH, "/"},
    {token.LT, "<"},
    {token.GT, ">"},
    {token.EOF, ""},
  }

  testTokens(t, input, expectations)
}

func TestNextToken_CanReadKeywords(t *testing.T) {
  input := `let fn true false if else return`

  expectations := []ExpectedTokens {
    {token.LET, "let"},
    {token.FUNCTION, "fn"},
    {token.TRUE, "true"},
    {token.FALSE, "false"},
    {token.IF, "if"},
    {token.ELSE, "else"},
    {token.RETURN, "return"},
  }

  testTokens(t, input, expectations)
}

func TestNextToken_CanReadMultiline(t *testing.T) {
  input := `{

  }`
  expectations := []ExpectedTokens {
    {token.LBRACE, "{"},
    {token.RBRACE, "}"},
  }

  testTokens(t, input, expectations)
}

func TestNextToken_CanReadIdentifiers(t *testing.T) {
  input := `let bob`

  expectations := []ExpectedTokens {
    {token.LET, "let"},
    {token.IDENT, "bob"},
  }

  testTokens(t, input, expectations)
}

func TestNextToken_CanReadNumbers(t *testing.T) {
  input := `123`

  expectations := []ExpectedTokens {
    {token.INT, "123"},
  }

  testTokens(t, input, expectations)
}

func testTokens(t *testing.T, input string, tests []ExpectedTokens) {
  l := NewLexer(input)

  for i, tt := range tests {
    tok := l.NextToken()

    if tok.Literal != tt.expectedLiteral {
      t.Fatalf("tests[%d] - literal wrong. expected %q, got %q", i, tt.expectedLiteral, tok.Literal)
    }

    if tok.Type != tt.expectedType {
      t.Fatalf("tests[%d] - tokentype wrong. expected %q, got %q", i, tt.expectedType, tok.Type)
    }
  }
}
