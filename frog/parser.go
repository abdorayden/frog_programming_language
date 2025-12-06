// Copyright (C) by abdenour souane
// you have a right to modify it upgrade it or do whatever you want
// but u have to keep my name on it

// all keywords functions types are structures that implements those 3 interfaces

package frog

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node // like extends Node in java
	statementNode()
}

type Expression interface {
	Node // like extends Node in java
	expressionNode()
}

// implement Node interface
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	out.WriteString("FRG_Begin\n")
	for _, s := range p.Statements {
		out.WriteString(s.String() + "\n")
	}
	out.WriteString("FRG_End\n")
	return out.String()
}

type DeclarationStatement struct {
	Token       Token // The type token (FRG_Int, FRG_Real, FRG_Strg)
	IsArray     bool  // FRG_Int[]
	Identifiers []*Identifier
}

func (ds *DeclarationStatement) statementNode() {
	// mark that Declarations implements Statement interface
	// because DeclarationStatement are statement u know what im saying :)
}
func (ds *DeclarationStatement) TokenLiteral() string {
	return ds.Token.Literal
}
func (ds *DeclarationStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ds.TokenLiteral())
	if ds.IsArray {
		out.WriteString("[]")
	}
	out.WriteString(" ")
	for i, ident := range ds.Identifiers {
		out.WriteString(ident.String())
		if i < len(ds.Identifiers)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(" #")
	return out.String()
}

// helper structure
type Parameter struct {
	Type Token // FRG_Int, FRG_Real, FRG_Strg
	Name *Identifier
}

type FunctionDeclarationStatement struct {
	Token      Token
	Name       *Identifier
	ReturnType Token // FRG_Int, FRG_Real, FRG_Strg
	Parameters []*Parameter
	Body       *BlockStatement
}

func (fds *FunctionDeclarationStatement) statementNode() {
	// same thing FunctionDeclaration are statement
	// read at line 72 :)
}
func (fds *FunctionDeclarationStatement) TokenLiteral() string {
	return fds.Token.Literal
}
func (fds *FunctionDeclarationStatement) String() string {
	var out bytes.Buffer
	out.WriteString(fds.TokenLiteral())
	out.WriteString(" ")
	out.WriteString(fds.Name.String())
	out.WriteString("(")
	for i, param := range fds.Parameters {
		out.WriteString(param.Type.Literal)
		out.WriteString(" ")
		out.WriteString(param.Name.String())
		if i < len(fds.Parameters)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(") : ")
	out.WriteString(fds.ReturnType.Literal)
	out.WriteString("\n")
	out.WriteString(fds.Body.String())
	return out.String()
}

type AssignmentStatement struct {
	Token Token
	Left  Expression // ID name as expression
	Value Expression // expression value
}

func (as *AssignmentStatement) statementNode() {
	// read at line 72 :)
}
func (as *AssignmentStatement) TokenLiteral() string { return as.Token.Literal }
func (as *AssignmentStatement) String() string {
	var out bytes.Buffer
	out.WriteString(as.Left.String())
	out.WriteString(" := ")
	if as.Value != nil {
		out.WriteString(as.Value.String())
	}
	out.WriteString(" #")
	return out.String()
}

type PrintStatement struct {
	Token       Token        // FRG_Print
	Expressions []Expression // 1+1*2 ...
}

func (ps *PrintStatement) statementNode() {
	// read at line 72 :)
}

func (ps *PrintStatement) TokenLiteral() string { return ps.Token.Literal }
func (ps *PrintStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ps.TokenLiteral() + " ")
	for i, expr := range ps.Expressions {
		out.WriteString(expr.String())
		if i < len(ps.Expressions)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(" #")
	return out.String()
}

type InputStatement struct {
	Token       Token
	Expressions []Expression
}

func (ps *InputStatement) statementNode() {
	// read at line 72 :)
}
func (ps *InputStatement) TokenLiteral() string { return ps.Token.Literal }
func (ps *InputStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ps.TokenLiteral() + " ")
	for i, expr := range ps.Expressions {
		out.WriteString(expr.String())
		if i < len(ps.Expressions)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(" #")
	return out.String()
}

// why is not array of statements because If has no body block
// so the If/Else body will be handled by Begin/End block
type IfStatement struct {
	Token       Token      // If
	Condition   Expression // condition 1+1 == 2
	Consequence Statement  // If Statement
	Alternative Statement  // Else Statement
}

func (is *IfStatement) statementNode() {
	// read at line 72 :)
}
func (is *IfStatement) TokenLiteral() string { return is.Token.Literal }
func (is *IfStatement) String() string {
	var out bytes.Buffer
	out.WriteString("If ")
	out.WriteString(is.Condition.String())
	out.WriteString(" ")
	out.WriteString(is.Consequence.String())
	if is.Alternative != nil {
		out.WriteString(" Else ")
		out.WriteString(is.Alternative.String())
	}
	return out.String()
}

type RepeatStatement struct {
	Token     Token       // Repeat
	Condition Expression  // condition expression
	Body      []Statement // the body array of statements
}

func (rs *RepeatStatement) statementNode() {
	// read at line 72 :)
}
func (rs *RepeatStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *RepeatStatement) String() string {
	var out bytes.Buffer
	out.WriteString("Repeat\n")
	for _, s := range rs.Body {
		out.WriteString(s.String() + "\n")
	}
	out.WriteString("until ")
	out.WriteString(rs.Condition.String())
	return out.String()
}

type BlockStatement struct {
	Token      Token       // Begin/Else
	Statements []Statement // block statements
}

func (bs *BlockStatement) statementNode() {
	// read at line 72 :)
}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	out.WriteString("Begin\n")
	for _, s := range bs.Statements {
		out.WriteString(s.String() + "\n")
	}
	out.WriteString("End")
	return out.String()
}

type BreakStatement struct {
	Token Token // Break
}

func (bs *BreakStatement) statementNode() {
	// read at line 72 :)

}
func (bs *BreakStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BreakStatement) String() string       { return bs.TokenLiteral() + " #" }

type ContinueStatement struct {
	Token Token // Continue
}

func (cs *ContinueStatement) statementNode() {
	// read at line 72 :)
}
func (cs *ContinueStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *ContinueStatement) String() string       { return cs.TokenLiteral() + " #" }

type UseStatement struct {
	Token    Token          // The FRG_Use token
	Filename *StringLiteral // frog file name (frog code) | string litteral
}

func (us *UseStatement) statementNode() {
	// read at line 72 :)
}
func (us *UseStatement) TokenLiteral() string { return us.Token.Literal }
func (us *UseStatement) String() string {
	var out bytes.Buffer
	out.WriteString(us.TokenLiteral())
	out.WriteString(" ")
	out.WriteString(us.Filename.String())
	out.WriteString(" #")
	return out.String()
}

// ExpressionStatement is a statement that consists of a single expression.
type ExpressionStatement struct {
	Token      Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {
	// read at line 72 :)
}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type Identifier struct {
	Token Token  // the IDENT token
	Value string // <ident name>
}

func (i *Identifier) expressionNode() {
	// mark as expression node
	// Identifier struct are implements Expression interface
	// because ID is expression u know what im saying :)
}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.Value
}

// integerliteral represents an integer literal.
type IntegerLiteral struct {
	Token Token // Number
	Value int64 // [0-9]+
}

func (il *IntegerLiteral) expressionNode() {
	// mark as expression node
	// read line 362 :))
}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// RealLiteral represents a real number literal.
type RealLiteral struct {
	Token Token   // Real
	Value float64 // [0-9]+\.[0-9]+
}

func (rl *RealLiteral) expressionNode() {
	// mark as expression node
	// read line 362 :))
}
func (rl *RealLiteral) TokenLiteral() string { return rl.Token.Literal }
func (rl *RealLiteral) String() string       { return rl.Token.Literal }

// StringLiteral represents a string literal.
type StringLiteral struct {
	Token Token  // FRG_Strg
	Value string // the value
}

func (sl *StringLiteral) expressionNode() {
	// mark as expression node
	// read line 362 :))
}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return "\"" + sl.Value + "\"" }

// PrefixExpression represents a prefix operator expression (e.g., -5).
type PrefixExpression struct {
	Token    Token      // The prefix token, e.g. - (left one)
	Operator string     // the operator
	Right    Expression // the right expression
}

func (pe *PrefixExpression) expressionNode() {
	// mark as expression node
	// read line 362 :))
}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

// InfixExpression represents an infix operator expression (e.g., 5 + 5).
type InfixExpression struct {
	Token    Token      // The operator token, e.g. + (the middle)
	Left     Expression // the left expression
	Operator string     // the operator
	Right    Expression // the right operator
}

func (ie *InfixExpression) expressionNode() {
	// mark as expression node
	// read line 362 :))
}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}

type ArrayLiteral struct {
	/*
		{
			1+1,
			10,
			...
		}
	*/
	Token    Token        // the '{' token
	Elements []Expression // values can be an Expressions
}

func (al *ArrayLiteral) expressionNode() {
	// mark as expression node
	// read line 362 :))
}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("{")
	for i, el := range al.Elements {
		out.WriteString(el.String())
		if i < len(al.Elements)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString("}")
	return out.String()
}

type ArraySizeLiteral struct {
	/*
		FRG_Int[] xs #
		xs := [20 + 5] #
				^
				expression
	*/
	Token Token      // the '[' token
	Size  Expression // size can be Expression
}

func (asl *ArraySizeLiteral) expressionNode() {
	// mark as expression node
	// read line 362 :))
}
func (asl *ArraySizeLiteral) TokenLiteral() string { return asl.Token.Literal }
func (asl *ArraySizeLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("[")
	out.WriteString(asl.Size.String())
	out.WriteString("]")
	return out.String()
}

type IndexExpression struct {
	Token Token // The [ token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode() {
	// mark as expression node
	// read line 362 :))
}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	return out.String()
}

type CallExpression struct {
	Token     Token // The ( token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {
	// mark as expression node
	// read line 362 :))
}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	for i, arg := range ce.Arguments {
		out.WriteString(arg.String())
		if i < len(ce.Arguments)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(")")
	return out.String()
}

// =============================================================================
// parser : after implements all structures for each statement and expression interfaces
// =============================================================================

// NOTE: using (The Pratt parser algorithm)
// check reffrence : https://en.wikipedia.org/wiki/Operator-precedence_parser
// check reffrence : https://en.wikipedia.org/wiki/LL_parser - for diffrence between LL and OPP

// Example1:		Parse id + id * id
// 	Stack           | Remaining Input  |   Action
//  ----------------+------------------+-----------------------------
//  $               |  id + id * id $  |   Shift id
//  $ id            |   + id * id $    |   Shift +
//  $ id +          |    id * id $     |   Shift id
//  $ id + id       |     * id $       |   Now check: + ⋖ * (shift)
//  $ id + id *     |      id $        |   Shift id
//  $ id + id * id  |        $         |   Check: * ⋗ $ (reduce!)

// Example2: a * b + c / d - e
// Step 1:  [a * b] + c / d - e
//
//	 ↓
//	E1 + c / d - e
//
// Step 2:  E1 + [c / d] - e
//
//	       ↓
//	E1 + E2 - e
//
// Step 3:  [E1 + E2] - e
//
//	  ↓
//	E3 - e
//
// Step 4:  [E3 - e]
//
//	  ↓
//	E4 (final result)
//
// Step 4:  [E3 - e]
//
//	  ↓
//	E4 (final result)
//
//	      E4
//	      |
//	     (-)
//	     / \
//	   E3   e
//	   |
//	  (+)
//	  / \
//	E1   E2
//	|     |
//
// (*)   (/)
// / \   / \
// a b   c d
//
// start parsing from bottom to top

// Precedence constants define operator precedence levels for the Pratt parser.
// The Pratt parser uses these constants to correctly handle operator precedence
// when parsing expressions like "a + b * c" which should be parsed as "a + (b * c)".
//
// How it works:
//  1. parseExpression(precedence int) parses expressions starting with a minimum precedence level (LOWEST)
//  2. It continues parsing infix operators (like +, -, etc.) as long as the next
//     operator has higher precedence than the current precedence parameter
//  3. The condition `precedence < p.peekPrecedence()` determines whether to continue
//
// Why LOWEST is used everywhere:
// In statement contexts (assignments, print statements, if conditions, etc.), we use
// parseExpression(LOWEST) because we want to parse ALL possible operations in the
// expression, regardless of precedence level. The precedence constraint in these
// contexts comes from statement termination (e.g., semicolon marked by TokenHash),
// not from operator precedence relationships. The Pratt parser algorithm itself
// handles the precedence internally through the precedence comparison.
const (
	_ int = iota
	// the lowest level
	LOWEST
	EQUALS      // ==, !=
	LESSGREATER // >, <, >=, <=
	SUM         // +
	PRODUCT     // *
	INDEX       // []
	CALL        // function()
	PREFIX      // -X
)

// map of TokenType and OpLevels
var precedences = map[TokenType]int{
	TokenEqual:        EQUALS,
	TokenNotEqual:     EQUALS,
	TokenLessThan:     LESSGREATER,
	TokenGreaterThan:  LESSGREATER,
	TokenLessEqual:    LESSGREATER,
	TokenGreaterEqual: LESSGREATER,
	TokenPlus:         SUM,
	TokenMinus:        SUM,
	TokenAsterisk:     PRODUCT,
	TokenSlash:        PRODUCT,
	TokenModulo:       PRODUCT,
	TokenLBracket:     INDEX,
	TokenLParen:       CALL,
}

type (
	// similare to typedef Expression(*prefixParseFn)(void); in C
	prefixParseFn func() Expression // the left side
	// similare to typedef Expression(*infixParseFn)(Expression); in C
	infixParseFn func(Expression) Expression // the middle side
)

// parser structure contain the informations
type Parser struct {
	// pointer to lexer that passed from the constructor
	// used to walk througth the tokens
	lexer *Lexer

	// array of strings for error parsing happend
	errors []string

	// currentToken and peekToken two fields of Token struct type
	// have informations of the current token
	currentToken Token
	peekToken    Token

	// two maps with same key (TokenType) and diffrent values prefixParseFn and infixParseFn
	// which is function type that defined at line 574
	// for each token will have it's own function parser , that parse the token
	// those two maps are needed for parseExpression method

	// map of each TokenType and it's callback parsing (prefix/infix)
	prefixParseFns map[TokenType]prefixParseFn
	infixParseFns  map[TokenType]infixParseFn
}

// constructor
func NewParser(l *Lexer) *Parser {
	p := &Parser{
		lexer:  l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[TokenType]prefixParseFn)

	p.registerPrefix(TokenIdentifier, p.parseIdentifier)
	p.registerPrefix(TokenNumber, p.parseNumberLiteral)
	p.registerPrefix(TokenString, p.parseStringLiteral)
	p.registerPrefix(TokenTrue, p.parseBooleanLiteral)
	p.registerPrefix(TokenFalse, p.parseBooleanLiteral)
	p.registerPrefix(TokenMinus, p.parsePrefixExpression)
	p.registerPrefix(TokenLBrace, p.parseArrayLiteral)
	p.registerPrefix(TokenLBracket, p.parseArraySizeLiteral)

	p.infixParseFns = make(map[TokenType]infixParseFn)

	// NOTE: rayden was here
	p.registerInfix(TokenPlus, p.parseInfixExpression)
	p.registerInfix(TokenMinus, p.parseInfixExpression)
	p.registerInfix(TokenAsterisk, p.parseInfixExpression)
	p.registerInfix(TokenSlash, p.parseInfixExpression)
	p.registerInfix(TokenModulo, p.parseInfixExpression)
	p.registerInfix(TokenEqual, p.parseInfixExpression)
	p.registerInfix(TokenNotEqual, p.parseInfixExpression)
	p.registerInfix(TokenLessThan, p.parseInfixExpression)
	p.registerInfix(TokenGreaterThan, p.parseInfixExpression)
	p.registerInfix(TokenLessEqual, p.parseInfixExpression)
	p.registerInfix(TokenGreaterEqual, p.parseInfixExpression)
	p.registerInfix(TokenLBracket, p.parseIndexExpression)
	p.registerInfix(TokenLParen, p.parseCallExpression)

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) registerPrefix(tokenType TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t TokenType) {
	msg := fmt.Sprintf(
		"ERROR: expected next token to be %s, got %s instead (line %d, col %d)",
		TokenToString(t), TokenToString(p.peekToken.Type), p.peekToken.Line, p.peekToken.Column)
	p.errors = append(p.errors, msg)
}

func (p *Parser) IsThereAnyErrors() bool {
	return len(p.errors) != 0
}

// the Program struct that defined above contains array of Statement interface
// this array can hold any struct implemented the interface
// because frog file are body main start with FRG_Begin ... FRG_End between this block contains statements
// so after we parse the programe we will get pointer Program struct with array of statements

// ParseProgram is the main entry point for the parser that parses an entire Frog program.
// It expects a program to begin with FRG_Begin and end with FRG_End tokens.
// This function implements a top-level parsing loop that continues until EOF or FRG_End is encountered.
func (p *Parser) ParseProgram() *Program {
	program := &Program{}
	program.Statements = []Statement{}

	// frog code must start with FRG_Begin
	if !p.expectCurrent(TokenFRGBegin) {
		p.errors = append(p.errors, "program must start with FRG_Begin")
		return program
	}

	// escape FRG_Begin
	p.nextToken()

	// end used to detect if FRG_End token are at the end of the frog file or not
	var end bool = false

	for !p.currentTokenIs(TokenEOF) {
		if p.currentTokenIs(TokenFRGEnd) {
			// parsed successfuly and get FRG_End Token
			end = true
			break
		}
		// call private Parser method that parse the statements
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	// checked :)
	if !end {
		p.errors = append(p.errors, "program must end with FRG_End")
	}

	return program
}

// parseStatement is the main statement dispatcher that routes parsing to the appropriate function
// based on the current token type. This function implements a recursive descent approach to
// statement parsing where each token type has its own dedicated parsing method.
func (p *Parser) parseStatement() Statement {
	switch p.currentToken.Type {
	case TokenFRGInt, TokenFRGReal, TokenFRGStrg:
		// DONE
		return p.parseDeclarationStatement() // parsing Token Declarations
	case TokenFRGPrint:
		// DONE
		return p.parsePrintStatement()
	case TokenIf:
		// TOP to BOTTOM recursive parsing
		// DONE
		return p.parseIfStatement()
	case TokenRepeat:
		// DONE
		return p.parseRepeatStatement()
	case TokenBegin:
		// DONE
		return p.parseBlockStatement()
	case TokenBreak:
		// DONE
		return p.parseBreakStatement()
	case TokenContinue:
		// DONE
		return p.parseContinueStatement()
	case TokenFRGUse:
		return p.parseUseStatementAndInclude()
	case TokenFRGFn:
		return p.parseFunctionDeclarationStatement()
	case TokenFRGInput:
		return p.parseInputStatement()
	case TokenIdentifier:
		// (2*x+1 / 2)
		expr := p.parseExpression(LOWEST)
		// parse identifier function (parseIdentifier) implement Expression interface
		// we used parseExpression and it will called this method
		// and return just definition of the token nothing special
		if p.peekTokenIs(TokenAssign) {
			stmt := &AssignmentStatement{Left: expr}
			p.nextToken() // kill :=
			stmt.Token = p.currentToken
			p.nextToken()
			stmt.Value = p.parseExpression(LOWEST)
			if !p.expectPeek(TokenHash) {
				return nil
			}
			return stmt
		} else {
			if ident, ok := expr.(*Identifier); ok {
				if ident.Value == "if" {
					msg := fmt.Sprintf("ERROR: syntax error, did you mean 'If'? (line %d, col %d)", ident.Token.Line, ident.Token.Column)
					p.errors = append(p.errors, msg)
					return nil
				} else if ident.Value == "else" {
					msg := fmt.Sprintf("ERROR: syntax error, did you mean 'Else'? (line %d, col %d)", ident.Token.Line, ident.Token.Column)
					p.errors = append(p.errors, msg)
					return nil
				} else if ident.Value == "repeat" {
					msg := fmt.Sprintf("ERROR: syntax error, did you mean 'Repeat'? (line %d, col %d)", ident.Token.Line, ident.Token.Column)
					p.errors = append(p.errors, msg)
					return nil
				} else if ident.Value == "until" {
					msg := fmt.Sprintf("ERROR: syntax error, did you mean 'Until'? (line %d, col %d)", ident.Token.Line, ident.Token.Column)
					p.errors = append(p.errors, msg)
					return nil
				} else if ident.Value == "begin" {
					msg := fmt.Sprintf("ERROR: syntax error, did you mean 'Begin'? (line %d, col %d)", ident.Token.Line, ident.Token.Column)
					p.errors = append(p.errors, msg)
					return nil
				} else if ident.Value == "end" {
					msg := fmt.Sprintf("ERROR: syntax error, did you mean 'End'? (line %d, col %d)", ident.Token.Line, ident.Token.Column)
					p.errors = append(p.errors, msg)
					return nil
				}
			}
			if !p.expectPeek(TokenHash) {
				return nil
			}
			return &ExpressionStatement{Expression: expr}
		}
	}
	msg := fmt.Sprintf("ERROR: Unexpected token '%s' at line %d, column %d. Cannot parse it as a statement.", p.currentToken.Literal, p.currentToken.Line, p.currentToken.Column)
	p.errors = append(p.errors, msg)
	return nil
}

// parseDeclarationStatement parses variable declaration statements like "int x, y #".
// This function handles both single and multiple variable declarations with optional array types.
// The statement must end with a hash (#) token.
func (p *Parser) parseDeclarationStatement() *DeclarationStatement {
	stmt := &DeclarationStatement{Token: p.currentToken}
	// why it's array of Identifiers ??
	// simple because we can declare multiple variables in one line
	// FRG_Int a,b,c,d,e,f#
	stmt.Identifiers = []*Identifier{} /// initialize to 0

	if p.peekTokenIs(TokenLBracket) {
		// check if we declare a variable or array
		p.nextToken()
		if !p.expectPeek(TokenRBracket) {
			return nil // error FRG_Int[<something else>
		}
		stmt.IsArray = true // it's an array :)
	}

	if !p.expectPeek(TokenIdentifier) { // expect identifier
		return nil
	}

	stmt.Identifiers = append(stmt.Identifiers, &Identifier{Token: p.currentToken, Value: p.currentToken.Literal})

	for p.peekTokenIs(TokenComma) {
		// start iterating all ids
		p.nextToken()
		if !p.expectPeek(TokenIdentifier) {
			return nil
		}
		stmt.Identifiers = append(stmt.Identifiers, &Identifier{Token: p.currentToken, Value: p.currentToken.Literal})
	}

	if !p.expectPeek(TokenHash) { // must end with HASH(#)
		return nil
	}

	return stmt
}

// parseFunctionDeclarationStatement parses function declarations in the form:
// "fn functionName(paramType paramName, ...) : returnType Begin ... End".
// This function handles parsing the function name, parameter list, return type,
// and function body block.
func (p *Parser) parseFunctionDeclarationStatement() *FunctionDeclarationStatement {
	stmt := &FunctionDeclarationStatement{Token: p.currentToken}

	if !p.expectPeek(TokenIdentifier) {
		return nil
	}
	stmt.Name = &Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectPeek(TokenLParen) {
		return nil
	}

	// Parse parameters
	stmt.Parameters = []*Parameter{}
	if !p.peekTokenIs(TokenRParen) {
		p.nextToken()
		for {
			param := &Parameter{}
			if !p.currentTokenIs(TokenFRGInt) && !p.currentTokenIs(TokenFRGReal) && !p.currentTokenIs(TokenFRGStrg) {
				p.errors = append(p.errors, fmt.Sprintf("expected parameter type, got %s", p.currentToken.Literal))
				return nil
			}
			param.Type = p.currentToken

			if !p.expectPeek(TokenIdentifier) {
				return nil
			}
			param.Name = &Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

			stmt.Parameters = append(stmt.Parameters, param)

			if p.peekTokenIs(TokenRParen) {
				break
			}
			if !p.expectPeek(TokenComma) {
				return nil
			}
			p.nextToken()
		}
	}

	if !p.expectPeek(TokenRParen) {
		return nil
	}

	if !p.expectPeek(TokenColon) {
		return nil
	}

	p.nextToken()
	if !p.currentTokenIs(TokenFRGInt) && !p.currentTokenIs(TokenFRGReal) && !p.currentTokenIs(TokenFRGStrg) {
		p.errors = append(p.errors, fmt.Sprintf("expected return type, got %s", p.currentToken.Literal))
		return nil
	}
	stmt.ReturnType = p.currentToken

	p.nextToken()
	stmt.Body = p.parseBlockStatement()

	return stmt
}

func (p *Parser) parseAssignmentStatement() *AssignmentStatement {
	stmt := &AssignmentStatement{}

	stmt.Left = p.parseExpression(LOWEST)

	p.nextToken()
	stmt.Token = p.currentToken

	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)

	if !p.expectPeek(TokenHash) {
		return nil
	}

	return stmt
}

func (p *Parser) parseInputStatement() *InputStatement {
	stmt := &InputStatement{Token: p.currentToken}
	stmt.Expressions = []Expression{}

	p.nextToken()
	expr := p.parseIdentifier()
	stmt.Expressions = append(stmt.Expressions, expr)
	for p.peekTokenIs(TokenComma) {
		p.nextToken()
		p.nextToken()
		expr := p.parseIdentifier()
		stmt.Expressions = append(stmt.Expressions, expr)
	}
	if !p.expectPeek(TokenHash) {
		return nil
	}
	return stmt
}

func (p *Parser) parsePrintStatement() *PrintStatement {
	stmt := &PrintStatement{Token: p.currentToken}
	stmt.Expressions = []Expression{}

	p.nextToken() // skip FRG_Print

	expr := p.parseExpression(LOWEST)
	stmt.Expressions = append(stmt.Expressions, expr)

	for p.peekTokenIs(TokenComma) {
		// why comma ? LOL
		// you can print multiple ids or litterals
		// FRG_Print x , y , z + t , 12*10+1 #
		p.nextToken()
		p.nextToken()
		expr := p.parseExpression(LOWEST)
		stmt.Expressions = append(stmt.Expressions, expr)
	}

	if !p.expectPeek(TokenHash) {
		// must end with hash
		return nil
	}

	return stmt
}

func (p *Parser) parseUseStatementAndInclude() Statement {
	useToken := p.currentToken // Store the FRG_Use token

	if !p.expectPeek(TokenString) {
		return nil
	}
	filename := p.currentToken.Literal

	if !p.expectPeek(TokenHash) {
		return nil
	}

	// Read the content of the included file
	content, err := os.ReadFile(filename)
	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("ERROR: could not read included file %s: %v", filename, err))
		return nil
	}

	// Create a new lexer and parser for the included file
	includedLexer := NewLexer(string(content))
	includedParser := NewParser(includedLexer)

	// Parse the included file's program
	includedProgram := includedParser.ParseProgram()

	if includedParser.IsThereAnyErrors() {
		for _, err := range includedParser.Errors() {
			p.errors = append(p.errors, fmt.Sprintf("ERROR in included file %s: %s", filename, err))
		}
		return nil
	}

	// Return a BlockStatement containing all statements from the included file
	// This effectively "inlines" the included file's statements into the current AST.
	block := &BlockStatement{Token: useToken} // Use the TokenFRGUse token for the block
	block.Statements = includedProgram.Statements
	return block
}

// parseIfStatement parses if statements in the form "If [condition] statement [Else statement]".
// The condition expression is enclosed in square brackets and may optionally have an else clause.
// This implements conditional execution with both if and if-else variants.
func (p *Parser) parseIfStatement() *IfStatement {
	stmt := &IfStatement{Token: p.currentToken}

	if !p.expectPeek(TokenLBracket) { // [
		return nil
	}

	p.nextToken() // kill [

	stmt.Condition = p.parseExpression(LOWEST) // keep parsing expr condition

	if !p.expectPeek(TokenRBracket) { // must have ]
		return nil
	}

	p.nextToken() // kill ]

	// recursively
	stmt.Consequence = p.parseStatement() // parse if statement
	if p.peekTokenIs(TokenElse) {
		p.nextToken() // kill else
		p.nextToken()
		stmt.Alternative = p.parseStatement() // parse else

	}
	return stmt
}

func (p *Parser) parseRepeatStatement() *RepeatStatement {
	stmt := &RepeatStatement{Token: p.currentToken}
	stmt.Body = []Statement{}

	p.nextToken() // kill Repeat Token

	for !p.currentTokenIs(TokenUntil) && !p.currentTokenIs(TokenEOF) {
		// keep parsing body
		s := p.parseStatement()
		if s != nil {
			stmt.Body = append(stmt.Body, s)
		}
		p.nextToken()
	}

	// check if current token is Until if not web reach to the eof and that's error
	if !p.currentTokenIs(TokenUntil) {
		p.peekError(TokenUntil)
		return nil
	}

	if !p.expectPeek(TokenLBracket) { // [
		return nil
	}
	p.nextToken()                              // kill [
	stmt.Condition = p.parseExpression(LOWEST) // parse expression
	if !p.expectPeek(TokenRBracket) {          // ]
		return nil
	}
	return stmt
}

func (p *Parser) parseBlockStatement() *BlockStatement {
	block := &BlockStatement{Token: p.currentToken}
	block.Statements = []Statement{}

	p.nextToken() // kill Begin Token

	for !p.currentTokenIs(TokenEnd) && !p.currentTokenIs(TokenEOF) {
		// keep parsing the body
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	if !p.currentTokenIs(TokenEnd) { // if it's not token end that means we reach to the eof and that's error
		p.errors = append(p.errors, fmt.Sprintf("unterminated block statement, expected End, got %s", TokenToString(p.currentToken.Type)))
	}

	return block
}

func (p *Parser) parseBreakStatement() *BreakStatement {
	stmt := &BreakStatement{Token: p.currentToken}
	if !p.expectPeek(TokenHash) { // we need to end line with HASH(#)
		return nil
	}
	return stmt
}

func (p *Parser) parseContinueStatement() *ContinueStatement {
	stmt := &ContinueStatement{Token: p.currentToken}
	if !p.expectPeek(TokenHash) { // must end with HASH(#)
		return nil
	}
	return stmt
}

// parseExpression with simpler precedence-based recursive descent parser
//
// This function implements a recursive descent parser with operator precedence.
// It handles expressions like: 5 + 3 * 2, (x + y) * z, arr[5], func(3, 4), etc.
//
// Operator precedence hierarchy (higher number = higher precedence):
// | Level | Operators      | Description          |
// |-------|----------------|----------------------|
// | 0     | LOWEST         | lowest precedence    |
// | 1     | ==, !=         | equality             |
// | 2     | >, <, >=, <=   | less/greater than    |
// | 3     | +, -           | sum                  |
// | 4     | *, /, %        | product              |
// | 5     | []             | index/array access   |
// | 6     | function()     | function/call        |
// | 7     | -X             | prefix operators     |
//
// How it works:
// 1. Parse the leftmost operand (prefix expression)
// 2. Continue parsing infix operators and their right operands based on precedence
//
// Example 1: Parsing "5 + 3 * 2" with initial precedence LOWEST (0)
// Initially: currentToken='5', peekToken='+', precedence=LOWEST(0)
//
// | Step | Action                                | leftExp            | Next Token | Notes                          |
// |------|---------------------------------------|--------------------|------------|--------------------------------|
// | 1    | Parse prefix '5'                      | 5                  | peek: '+'  | leftExp = IntegerLiteral(5)    |
// | 2    | Check infix: +, SUM(4) > LOWEST(0)    | 5                  | '+'        | Process infix expression       |
// | 3    | Parse right side of + with SUM(4)     | 5 + ?              | '3'        | Calls parseExpression(SUM)     |
// | 4    | Parse prefix '3'                      | 5 + 3              | peek: '*'  | On right side of +             |
// | 5    | Check infix: *, PROD(5) > SUM(4)      | 5 + 3              | '*'        | Process higher precedence *    |
// | 6    | Parse right side of * with PROD(5)    | 5 + (3 * ?)        | '2'        | Calls parseExpression(PROD)    |
// | 7    | Parse prefix '2'                      | 5 + (3 * 2)        | peek: #    | On right side of *             |
// | 8    | No more infix ops                     | (5 + (3 * 2))      | -          | Done                           |
//
// Final result: (5 + (3 * 2)) - correct precedence: multiplication before addition
//
// Example 2: Parsing "arr[0] + 5" with initial precedence LOWEST (0)
// Initially: currentToken='arr', peekToken='[', precedence=LOWEST(0)
//
// | Step | Action                                | leftExp            | Next Token | Notes                          |
// |------|---------------------------------------|--------------------|------------|--------------------------------|
// | 1    | Parse prefix 'arr'                    | arr                | peek: '['  | leftExp = Identifier(arr)      |
// | 2    | Check infix: [, INDEX(5) > LOWEST(0)  | arr                | '['        | Process indexing               |
// | 3    | Parse index [0], INDEX(5)             | arr[0]             | peek: '+'  | leftExp updated to IndexExp    |
// | 4    | Check infix: +, SUM(3) > LOWEST(0)    | arr[0]             | '+'        | Process addition               |
// | 5    | Parse right side of + with SUM(3)     | arr[0] + ?         | '5'        | Calls parseExpression(SUM)     |
// | 6    | Parse prefix '5'                      | arr[0] + 5         | peek: #    | On right side of +             |
// | 7    | No more infix ops                     | (arr[0] + 5)       | -          | Done                           |
//
// Final result: (arr[0] + 5) - correct precedence: indexing before addition
//
// Example 3: Parsing "func(10) * 2" with initial precedence LOWEST (0)
// Initially: currentToken='func', peekToken='(', precedence=LOWEST(0)
//
// | Step | Action                                | leftExp              | Next Token | Notes                        |
// |------|---------------------------------------|----------------------|------------|------------------------------|
// | 1    | Parse prefix 'func'                   | func                 | peek: '('  | leftExp = Identifier(func)   |
// | 2    | Check infix: (, CALL(6) > LOWEST(0)   | func                 | '('        | Process function call        |
// | 3    | Parse function call func(10)          | func(10)             | peek: '*'  | leftExp updated to CallExp   |
// | 4    | Check infix: *, PROD(4) > LOWEST(0)   | func(10)             | '*'        | Process multiplication       |
// | 5    | Parse right side of * with PROD(4)    | func(10) * ?         | '2'        | Calls parseExpression(PROD)  |
// | 6    | Parse prefix '2'                      | func(10) * 2         | peek: #    | On right side of *           |
// | 7    | No more infix ops                     | (func(10) * 2)       | -          | Done                         |
func (p *Parser) parseExpression(precedence int) Expression {
	// STEP 1: Parse prefix expression (left operand)
	// Prefix parsers handle tokens that appear at the beginning of expressions:
	// - Identifiers: x, myVar
	// - Numbers: 123, 45.67
	// - Strings: "hello"
	// - Booleans: true, false
	// - Prefix operators: -5, !true
	// - Array literals: {1, 2, 3}
	prefix := p.prefixParseFns[p.currentToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.currentToken.Type)
		return nil
	}

	// Parse the leftmost operand (prefix expression)
	// This could be a number, identifier, string, or any other prefix expression
	leftExp := prefix()

	// STEP 2: Parse infix expressions (operators between operands)
	// Continue parsing infix operators and their right-hand expressions as long as:
	// 1. The next token is not the end-of-statement marker (#)
	// 2. The next operator has higher precedence than the current precedence level
	for !p.peekTokenIs(TokenHash) && precedence < p.peekPrecedence() {
		// Get the infix parser function for the peek token
		// Infix parsers handle operators that appear between operands:
		// - Arithmetic: +, -, *, /, %
		// - Comparison: ==, !=, <, >, <=, >=
		// - Indexing: [expression]
		// - Function calls: (arg1, arg2, ...)
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		// Move from peek token to current token (consume the operator)
		p.nextToken()

		// Parse the infix expression using the current left expression as input
		// The infix parser handles the operator and parses the right operand
		// with appropriate precedence to handle operator precedence correctly
		// For example, if we see 5 + 3 * 2:
		// - When processing +, right side calls parseExpression with SUM precedence
		// - This allows * to be processed before + due to higher precedence
		leftExp = infix(leftExp)
	}

	return leftExp
}

// parseIdentifier parses a single identifier token into an Identifier expression node.
// This is a prefix parser that handles variable names, function names, etc.
func (p *Parser) parseIdentifier() Expression {
	return &Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

// parseNumberLiteral parses numeric literals, distinguishing between integer and real (float) types
// based on whether the token contains a decimal point. This is a prefix parser for numeric values.
func (p *Parser) parseNumberLiteral() Expression {
	if bytes.ContainsRune([]byte(p.currentToken.Literal), '.') {
		lit := &RealLiteral{Token: p.currentToken}
		value, err := strconv.ParseFloat(p.currentToken.Literal, 64)
		if err != nil {
			msg := fmt.Sprintf("could not parse %q as float", p.currentToken.Literal)
			p.errors = append(p.errors, msg)
			return nil
		}
		lit.Value = value
		return lit
	}

	lit := &IntegerLiteral{Token: p.currentToken}
	value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.currentToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

func (p *Parser) parseStringLiteral() Expression {
	return &StringLiteral{Token: p.currentToken, Value: p.currentToken.Literal}
}

func (p *Parser) parseBooleanLiteral() Expression {
	return &Boolean{Token: p.currentToken, Value: p.currentTokenIs(TokenTrue)}
}

// parsePrefixExpression parses prefix operators (like unary minus) where the operator
// appears before its operand. The right-hand side is parsed with PREFIX precedence level
// to ensure proper operator precedence relationships.
func (p *Parser) parsePrefixExpression() Expression {
	expression := &PrefixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

// parseInfixExpression parses infix operators (like +, -, *, etc.) where the operator
// appears between two operands. The right-hand side is parsed with the current operator's
// precedence level to ensure proper precedence relationships in expressions.
func (p *Parser) parseInfixExpression(left Expression) Expression {
	expression := &InfixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
		Left:     left,
	}
	precedence := p.currentPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}

func (p *Parser) currentTokenIs(t TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectCurrent(t TokenType) bool {
	if p.currentTokenIs(t) {
		return true
	}
	p.errors = append(p.errors, fmt.Sprintf("expected current token to be %s, got %s instead", TokenToString(t), TokenToString(p.currentToken.Type)))
	return false
}

func (p *Parser) expectPeek(t TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

// peekPrecedence returns the precedence level of the next (peek) token.
// If the token type has no defined precedence, it returns LOWEST.
func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

// currentPrecedence returns the precedence level of the current token.
// If the token type has no defined precedence, it returns LOWEST.
func (p *Parser) currentPrecedence() int {
	if p, ok := precedences[p.currentToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) parseArrayLiteral() Expression {
	array := &ArrayLiteral{Token: p.currentToken}
	array.Elements = []Expression{}

	if p.peekTokenIs(TokenRBrace) {
		p.nextToken()
		return array
	}

	p.nextToken()
	array.Elements = append(array.Elements, p.parseExpression(LOWEST))

	for p.peekTokenIs(TokenComma) {
		p.nextToken()
		p.nextToken()
		array.Elements = append(array.Elements, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(TokenRBrace) {
		return nil
	}

	return array
}

func (p *Parser) parseArraySizeLiteral() Expression {
	array := &ArraySizeLiteral{Token: p.currentToken}

	p.nextToken()
	array.Size = p.parseExpression(LOWEST)

	if !p.expectPeek(TokenRBracket) {
		return nil
	}

	return array
}

func (p *Parser) parseIndexExpression(left Expression) Expression {
	exp := &IndexExpression{Token: p.currentToken, Left: left}

	p.nextToken()
	exp.Index = p.parseExpression(LOWEST)

	if !p.expectPeek(TokenRBracket) {
		return nil
	}

	return exp
}

func (p *Parser) parseCallExpression(function Expression) Expression {
	exp := &CallExpression{Token: p.currentToken, Function: function}
	exp.Arguments = p.parseExpressionList(TokenRParen)
	return exp
}

func (p *Parser) parseExpressionList(end TokenType) []Expression {
	list := []Expression{}
	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}
	p.nextToken()
	list = append(list, p.parseExpression(LOWEST))
	for p.peekTokenIs(TokenComma) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}
	if !p.expectPeek(end) {
		return nil
	}
	return list
}

func (p *Parser) noPrefixParseFnError(t TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found (line %d, col %d)", TokenToString(t), p.currentToken.Line, p.currentToken.Column)
	p.errors = append(p.errors, msg)
}
