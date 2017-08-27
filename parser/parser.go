package parser

import (
  "fmt"
  "monkey/ast"
  "monkey/lexer"
  "monkey/token"
)

type Parser struct {
  l *lexer.Lexer

  errors []string

  curToken  token.Token
  peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
  p := &Parser{
    l: l,
    errors: []string{},
  }

  p.nextToken()
  p.nextToken()

  return p
}

func (p *Parser) Errors() []string {
  return p.errors
}

func (p *Parser) ParseProgram() *ast.Program {
  program := &ast.Program{}
  program.Statements = []ast.Statement{}

  for !p.curTokenIs(token.EOF) {
    stmt := p.parseStatement()
    if stmt != nil {
      program.Statements = append(program.Statements, stmt)
    }
    p.nextToken()
  }
  return program
}

func (p *Parser) nextToken () {
  p.curToken = p.peekToken
  p.peekToken = p.l.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
  switch p.curToken.Type {
  case token.LET:
    return p.parseLetStatement()
  default:
    return nil
  }
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
  stmt := &ast.LetStatement{Token: p.curToken}

  if !p.assertAndSkipToken(token.IDENT) {
    return nil
  }

  stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

  if !p.assertAndSkipToken(token.ASSIGN) {
    return nil
  }

  //TODO: We're skipping the expressions until we encounter a semicolon
  for !p.curTokenIs(token.SEMICOLON) {
    p.nextToken()
  }

  return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
  return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
  return p.peekToken.Type == t
}

func (p *Parser) assertAndSkipToken(t token.TokenType) bool {
  if p.peekTokenIs(t) {
    p.nextToken()
    return true
  } else {
    p.addError(t)
    return false
  }
}

func (p *Parser) addError (t token.TokenType) {
  msg := fmt.Sprintf("expected next token to be %s, got %s instead. [%s] line %d, char %d", t, p.peekToken.Type, p.peekToken.Filename, p.peekToken.Line, p.peekToken.Character)
  p.errors = append(p.errors, msg)
}