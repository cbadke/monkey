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

  l.skipWhitespace()

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
  case '-':
    return newToken(token.MINUS, l.readChar())
  case '!':
    return newToken(token.BANG, l.readChar())
  case '*':
    return newToken(token.ASTERISK, l.readChar())
  case '/':
    return newToken(token.SLASH, l.readChar())
  case '<':
    return newToken(token.LT, l.readChar())
  case '>':
    return newToken(token.GT, l.readChar())
  case 0:
    return token.Token{Type: token.EOF, Literal: ""}
  default:
    if isLetter(l.peekChar()) {
      ident := l.readIdentifier()
      t := token.LookupIdent(ident)
      return token.Token{Type: t, Literal: ident}
    } else if isDigit(l.peekChar()) {
      return token.Token{Type: token.INT, Literal: l.readNumber()}
    }else {
      return newToken(token.ILLEGAL, l.readChar())
    }
  }
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
  return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) skipWhitespace() {
  var ch = l.peekChar()

  for ch == ' ' ||  ch == '\t' || ch == '\n' || ch == '\r' {
    l.readChar()
    ch = l.peekChar()
  }
}


func (l *Lexer) readCharGroup(fn (func(byte) bool)) string {
  position := l.peekPosition
  for fn(l.peekChar()) {
    l.readChar()
  }

  return l.input[position:l.peekPosition]
}

func (l *Lexer) readIdentifier() string {
  return l.readCharGroup(isLetter)
}

func (l *Lexer) readNumber() string {
  return l.readCharGroup(isDigit)
}

func isLetter(ch byte) bool {
  return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
  return '0' <= ch && ch <= '9'
}
