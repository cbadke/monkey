package lexer

import "monkey/token"

type Lexer struct {
  input        string
  peekPosition int
  lineCount    int
  filename     string
}

func NewLexer(input string, filename string) *Lexer {
  l := &Lexer{input: input, filename: filename, lineCount: 1}
  return l
}

func (l *Lexer) readChar() (byte, int, int) {
  var ch byte
  position := l.peekPosition

  if l.peekPosition >= len(l.input) {
    ch = 0
  } else {
    ch = l.input[l.peekPosition]
    l.peekPosition += 1
  }

  return ch, position, l.lineCount
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
    firstChar, position, line := l.readChar()
    if l.peekChar() == '=' {
      secondChar, _, _ := l.readChar()
      literal := string(firstChar) + string(secondChar)
      return token.Token{Type: token.EQ, Literal: literal, Filename: l.filename, Line: line, Character: position}
    } else {
      return newToken(token.ASSIGN, firstChar, position, line)
    }
  case '+':
    char, line, position := readChar()
    return newToken(token.PLUS, char, p)
  case '(':
    return newToken(token.LPAREN, l.readChar())
  case ')':
    return newToken(token.RPAREN, l.readChar())
  case '{':
    return newToken(token.LBRACE, l.readChar())
  case '}':
    return newToken(token.RBRACE, l.readChar())
  case '\'':
    return newToken(token.SINGLE_QUOTE, l.readChar())
  case '"':
    return newToken(token.QUOTE, l.readChar())
  case ',':
    return newToken(token.COMMA, l.readChar())
  case ';':
    return newToken(token.SEMICOLON, l.readChar())
  case '-':
    return newToken(token.MINUS, l.readChar())
  case '!':
    firstChar := l.readChar()

    if l.peekChar() == '=' {
      secondChar := l.readChar()
      literal := string(firstChar) + string(secondChar)
      return token.Token{Type: token.NOT_EQ, Literal: literal}
    } else {
      return newToken(token.BANG, firstChar)
    }
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

func newToken(tokenType token.TokenType, ch byte, pos int, line int) token.Token {
  return token.Token{Type: tokenType, Literal: string(ch), Filename: l.filename, Line: line, Character: pos}
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
