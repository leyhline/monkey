module leyhline.net/monkey/repl

go 1.17

require (
	leyhline.net/monkey/lexer v0.0.0
	leyhline.net/monkey/parser v0.0.0
)

replace leyhline.net/monkey/parser => ../parser

replace leyhline.net/monkey/lexer => ../lexer
