package lexer

import "monkey/token"

type Lexer struct {
  input        string
  peekPosition int
}

func NewLexer(input string) *Lexer {
  l := &Lexer{input: input}
  return l
}

func (l *Lexer) readChar() byte {
  var ch byte

  if l.peekPosition >= len(l.input) {
    ch = 0
  } else {
    ch = l.input[l.peekPosition]
  }
  l.peekPosition += 1
  return ch
}

func (l *Lexer) peekChar() byte {
  var ch byte

  if l.peekPosition >= len(l.input) {
    ch = 0
  } else {
    ch = l.input[l.peekPosition]
  }
  return ch
}

func (l *Lexer) NextToken() token.Token {
  switch l.peekChar() {
  case '=':
    return newToken(token.ASSIGN, l.readChar())
  case '+':
    return newToken(token.PLUS, l.readChar())
  case '(':
    return newToken(token.LPAREN, l.readChar())
  case ')':
    return newToken(token.RPAREN, l.readChar())
  case '{':
    return newToken(token.LBRACE, l.readChar())
  case '}':
    return newToken(token.RBRACE, l.readChar())
  case ',':
    return newToken(token.COMMA, l.readChar())
  case ';':
    return newToken(token.SEMICOLON, l.readChar())
  case 0:
    return token.Token{Type: token.EOF, Literal: ""}
  default:
    return newToken(token.ILLEGAL, 0)
  }
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
  return token.Token{Type: tokenType, Literal: string(ch)}
}
