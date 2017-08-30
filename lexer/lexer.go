package lexer

import "monkey/token"

type Lexer struct {
  input        string
  peekPosition int
  lineCount    int
  charCount    int
  filename     string
}

func NewLexer(input string, filename string) *Lexer {
  l := &Lexer{input: input, filename: filename, lineCount: 1, charCount: 0}
  return l
}

func (l *Lexer) readChar() (byte, int, int) {
  var ch byte

  if l.peekPosition >= len(l.input) {
    ch = 0
  } else {
    ch = l.input[l.peekPosition]
    l.peekPosition += 1
    l.charCount += 1
  }

  return ch, l.charCount, l.lineCount
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
      return newComplexToken(token.EQ, literal, position, line, l.filename)
    } else {
      return newToken(token.ASSIGN, firstChar, position, line, l.filename)
    }
  case '+':
    char, position, line := l.readChar()
    return newToken(token.PLUS, char, position, line, l.filename)
  case '(':
    char, position, line := l.readChar()
    return newToken(token.LPAREN, char, position, line, l.filename)
  case ')':
    char, position, line := l.readChar()
    return newToken(token.RPAREN, char, position, line, l.filename)
  case '{':
    char, position, line := l.readChar()
    return newToken(token.LBRACE, char, position, line, l.filename)
  case '}':
    char, position, line := l.readChar()
    return newToken(token.RBRACE, char, position, line, l.filename)
  case '\'':
    char, position, line := l.readChar()
    return newToken(token.SINGLE_QUOTE, char, position, line, l.filename)
  case '"':
    char, position, line := l.readChar()
    return newToken(token.QUOTE, char, position, line, l.filename)
  case ',':
    char, position, line := l.readChar()
    return newToken(token.COMMA, char, position, line, l.filename)
  case ';':
    char, position, line := l.readChar()
    return newToken(token.SEMICOLON, char, position, line, l.filename)
  case '-':
    char, position, line := l.readChar()
    return newToken(token.MINUS, char, position, line, l.filename)
  case '!':
    firstChar, position, line := l.readChar()

    if l.peekChar() == '=' {
      secondChar, _, _ := l.readChar()
      literal := string(firstChar) + string(secondChar)
      return newComplexToken(token.NOT_EQ, literal, position, line, l.filename)
    } else {
      return newToken(token.BANG, firstChar, position, line, l.filename)
    }
  case '*':
    char, position, line := l.readChar()
    return newToken(token.ASTERISK, char, position, line, l.filename)
  case '/':
    char, position, line := l.readChar()
    return newToken(token.SLASH, char, position, line, l.filename)
  case '<':
    char, position, line := l.readChar()
    return newToken(token.LT, char, position, line, l.filename)
  case '>':
    char, position, line := l.readChar()
    return newToken(token.GT, char, position, line, l.filename)
  case 0:
    return newComplexToken(token.EOF, "", l.charCount, l.lineCount, l.filename)
  default:
    if isLetter(l.peekChar()) {
      ident, position, line := l.readIdentifier()
      t := token.LookupIdent(ident)
      return newComplexToken(t, ident, position, line, l.filename)
    } else if isDigit(l.peekChar()) {
      number, position, line := l.readNumber()
      return newComplexToken(token.INT, number, position, line, l.filename)
    }else {
      char, position, line := l.readChar()
      return newToken(token.ILLEGAL, char, position, line, l.filename)
    }
  }
}

func newToken(tokenType token.TokenType, ch byte, pos int, line int, fileName string) token.Token {
  return token.Token{Type: tokenType, Literal: string(ch), Filename: fileName, Line: line, Character: pos}
}

func newComplexToken(tokenType token.TokenType, literal string, pos int, line int, fileName string) token.Token {
  return token.Token{Type: tokenType, Literal: literal, Filename: fileName, Line: line, Character: pos}
}

func (l *Lexer) skipWhitespace() {
  var ch = l.peekChar()

  for ch == ' ' ||  ch == '\t' || ch == '\n' || ch == '\r' {
    if ch == '\n' {
      l.lineCount += 1
      l.charCount = 0;
    }

    l.readChar()
    ch = l.peekChar()
  }
}


func (l *Lexer) readCharGroup(fn (func(byte) bool)) (string, int, int) {
  charPos := l.charCount
  position:= l.peekPosition
  for fn(l.peekChar()) {
    l.readChar()
  }

  return l.input[position:l.peekPosition], charPos, l.lineCount
}

//allow numbers in identifiers (if not first character)
func (l *Lexer) readIdentifier() (string, int, int) {
  return l.readCharGroup(isLetter)
}

func (l *Lexer) readNumber() (string, int, int) {
  return l.readCharGroup(isDigit)
}

func isLetter(ch byte) bool {
  return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
  return '0' <= ch && ch <= '9'
}
