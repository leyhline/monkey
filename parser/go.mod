module leyhline.net/monkey/parser

go 1.17

require (
	leyhline.net/monkey/ast v0.0.0
	leyhline.net/monkey/lexer v0.0.0
	leyhline.net/monkey/token v0.0.0
)

replace leyhline.net/monkey/token => ../token

replace leyhline.net/monkey/lexer => ../lexer

replace leyhline.net/monkey/ast => ../ast
