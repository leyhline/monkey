module leyhline.net/monkey

go 1.17

require leyhline.net/monkey/repl v0.0.0

require (
	leyhline.net/monkey/ast v0.0.0 // indirect
	leyhline.net/monkey/evaluator v0.0.0 // indirect
	leyhline.net/monkey/lexer v0.0.0 // indirect
	leyhline.net/monkey/object v0.0.0 // indirect
	leyhline.net/monkey/parser v0.0.0 // indirect
	leyhline.net/monkey/token v0.0.0 // indirect
)

replace leyhline.net/monkey/repl => ./repl

replace leyhline.net/monkey/token => ./token

replace leyhline.net/monkey/lexer => ./lexer

replace leyhline.net/monkey/parser => ./parser

replace leyhline.net/monkey/ast => ./ast

replace leyhline.net/monkey/evaluator => ./evaluator

replace leyhline.net/monkey/object => ./object
