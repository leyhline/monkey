package parser

import "leyhline.net/monkey/ast"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() ast.Token {
	var tok ast.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = ast.Token{Type: ast.EQ, Literal: "=="}
		} else {
			tok = newToken(ast.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(ast.SEMICOLON, l.ch)
	case '(':
		tok = newToken(ast.LPAREN, l.ch)
	case ')':
		tok = newToken(ast.RPAREN, l.ch)
	case ',':
		tok = newToken(ast.COMMA, l.ch)
	case '+':
		tok = newToken(ast.PLUS, l.ch)
	case '-':
		tok = newToken(ast.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = ast.Token{Type: ast.NOT_EQ, Literal: "!="}
		} else {
			tok = newToken(ast.BANG, l.ch)
		}
	case '/':
		tok = newToken(ast.SLASH, l.ch)
	case '*':
		tok = newToken(ast.ASTERISK, l.ch)
	case '<':
		tok = newToken(ast.LT, l.ch)
	case '>':
		tok = newToken(ast.GT, l.ch)
	case '{':
		tok = newToken(ast.LBRACE, l.ch)
	case '}':
		tok = newToken(ast.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = ast.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = ast.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = ast.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(ast.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenType ast.TokenType, ch byte) ast.Token {
	return ast.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
