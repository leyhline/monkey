module leyhline.net/monkey/repl

go 1.17

require (
	leyhline.net/monkey/evaluator v0.0.0
	leyhline.net/monkey/parser v0.0.0
)

require leyhline.net/monkey/ast v0.0.0 // indirect

replace leyhline.net/monkey/parser => ../parser

replace leyhline.net/monkey/evaluator => ../evaluator

replace leyhline.net/monkey/ast => ../ast
