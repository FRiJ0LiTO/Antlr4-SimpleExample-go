package listener

import (
	"ExampleAntlr/parser"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
)

type calcListener struct {
	*parser.BaseCalcListener
	stack []int
}

// push adds an integer value to the stack.
// It appends the given integer `i` to the `stack` slice.
//
// Parameters:
//
//	i - The integer value to be pushed onto the stack.
func (l *calcListener) push(i int) {
	l.stack = append(l.stack, i)
}

// pop removes and returns the top integer value from the stack.
// It retrieves the last element of the `stack` slice, removes it, and returns it.
// If the stack is empty, it panics.
//
// Returns:
//
//	int - The top integer value from the stack.
//
// Panics:
//
//	If the stack is empty, a panic is triggered with the message "stack is empty unable to pop".
func (l *calcListener) pop() int {
	if len(l.stack) < 1 {
		panic("Stack is empty unable to pop")
	}
	result := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]
	return result
}

// ExitMulDiv is triggered when exiting a multiplication or division operation in the parse tree.
// It pops the top two values from the stack, performs the operation (multiplication or division),
// and pushes the result back onto the stack.
//
// Parameters:
//
//	c - The context of the MulDiv rule, which provides access to the operator and operands.
//
// Panics:
//
//	If an unexpected operator is encountered, a panic is triggered with a message indicating the invalid operator.
func (l *calcListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.pop(), l.pop()
	switch c.GetOp().GetTokenType() {
	case parser.CalcLexerMUL:
		l.push(left * right)
	case parser.CalcLexerDIV:
		l.push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

// ExitAddSub is triggered when exiting an addition or subtraction operation in the parse tree.
// It pops the top two values from the stack, performs the operation (addition or subtraction),
// and pushes the result back onto the stack.
//
// Parameters:
//
//	c - The context of the AddSub rule, which provides access to the operator and operands.
//
// Panics:
//
//	If an unexpected operator is encountered, a panic is triggered with a message indicating the invalid operator.
func (l *calcListener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.pop(), l.pop()
	switch c.GetOp().GetTokenType() {
	case parser.CalcLexerADD:
		l.push(left + right)
	case parser.CalcLexerSUB:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

// ExitNumber is triggered when exiting a number node in the parse tree.
// It converts the text of the number node to an integer and pushes it onto the stack.
//
// Parameters:
//
//	c - The context of the Number rule, which provides access to the text of the number.
//
// Panics:
//
//	If the text cannot be converted to an integer, a panic is triggered with the error message.
func (l *calcListener) ExitNumber(c *parser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}
	l.push(i)
}

// Calc evaluates a mathematical expression provided as a string and returns the result as an integer.
// It uses ANTLR to parse the input string, constructs a parse tree, and processes it using a custom listener.
//
// Parameters:
//
//	input - A string containing the mathematical expression to be evaluated.
//
// Returns:
//
//	int - The result of evaluating the mathematical expression.
func Calc(input string) int {
	// Set up the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener calcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start_())

	return listener.pop()
}
