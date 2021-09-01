package parser

import (
	"leyhline.net/monkey/ast"
	"testing"
)

func TestNextToken1(t *testing.T) {
	input := `=+(){},;-!/*<>[]`
	tests := []struct {
		expectedType    ast.TokenType
		expectedLiteral string
	}{
		{ast.ASSIGN, "="},
		{ast.PLUS, "+"},
		{ast.LPAREN, "("},
		{ast.RPAREN, ")"},
		{ast.LBRACE, "{"},
		{ast.RBRACE, "}"},
		{ast.COMMA, ","},
		{ast.SEMICOLON, ";"},
		{ast.MINUS, "-"},
		{ast.BANG, "!"},
		{ast.SLASH, "/"},
		{ast.ASTERISK, "*"},
		{ast.LT, "<"},
		{ast.GT, ">"},
		{ast.LBRACKET, "["},
		{ast.RBRACKET, "]"},
		{ast.EOF, ""},
	}
	l := NewLexer(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken2(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
	x + y;
};
let result = add(five, ten);
"foobar"
"foo bar"`
	tests := []struct {
		expectedType    ast.TokenType
		expectedLiteral string
	}{
		{ast.LET, "let"},
		{ast.IDENT, "five"},
		{ast.ASSIGN, "="},
		{ast.INT, "5"},
		{ast.SEMICOLON, ";"},
		{ast.LET, "let"},
		{ast.IDENT, "ten"},
		{ast.ASSIGN, "="},
		{ast.INT, "10"},
		{ast.SEMICOLON, ";"},
		{ast.LET, "let"},
		{ast.IDENT, "add"},
		{ast.ASSIGN, "="},
		{ast.FUNCTION, "fn"},
		{ast.LPAREN, "("},
		{ast.IDENT, "x"},
		{ast.COMMA, ","},
		{ast.IDENT, "y"},
		{ast.RPAREN, ")"},
		{ast.LBRACE, "{"},
		{ast.IDENT, "x"},
		{ast.PLUS, "+"},
		{ast.IDENT, "y"},
		{ast.SEMICOLON, ";"},
		{ast.RBRACE, "}"},
		{ast.SEMICOLON, ";"},
		{ast.LET, "let"},
		{ast.IDENT, "result"},
		{ast.ASSIGN, "="},
		{ast.IDENT, "add"},
		{ast.LPAREN, "("},
		{ast.IDENT, "five"},
		{ast.COMMA, ","},
		{ast.IDENT, "ten"},
		{ast.RPAREN, ")"},
		{ast.SEMICOLON, ";"},
		{ast.STRING, "foobar"},
		{ast.STRING, "foo bar"},
		{ast.EOF, ""},
	}
	l := NewLexer(input)
	for i, tt := range tests {
		tok := l.NextToken()
		t.Log(tok.Literal)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken3(t *testing.T) {
	input := `if (5 < 10) {
	return true;
} else {
	return false;
}
10 == 10;
10 != 9;`
	tests := []struct {
		expectedType    ast.TokenType
		expectedLiteral string
	}{
		{ast.IF, "if"},
		{ast.LPAREN, "("},
		{ast.INT, "5"},
		{ast.LT, "<"},
		{ast.INT, "10"},
		{ast.RPAREN, ")"},
		{ast.LBRACE, "{"},
		{ast.RETURN, "return"},
		{ast.TRUE, "true"},
		{ast.SEMICOLON, ";"},
		{ast.RBRACE, "}"},
		{ast.ELSE, "else"},
		{ast.LBRACE, "{"},
		{ast.RETURN, "return"},
		{ast.FALSE, "false"},
		{ast.SEMICOLON, ";"},
		{ast.RBRACE, "}"},
		{ast.INT, "10"},
		{ast.EQ, "=="},
		{ast.INT, "10"},
		{ast.SEMICOLON, ";"},
		{ast.INT, "10"},
		{ast.NOT_EQ, "!="},
		{ast.INT, "9"},
		{ast.SEMICOLON, ";"},
		{ast.EOF, ""},
	}
	l := NewLexer(input)
	for i, tt := range tests {
		tok := l.NextToken()
		t.Log(tok.Literal)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
