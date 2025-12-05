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
	IsArray     bool
	Identifiers []*Identifier
}

type FunctionDeclarationStatement struct {
	Token      Token
	Name       *Identifier
	ReturnType Token // FRG_Int, FRG_Real, FRG_Strg
	Parameters []*Parameter
	Body       *BlockStatement
}

type Parameter struct {
	Type Token // FRG_Int, FRG_Real, FRG_Strg
	Name *Identifier
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
	Left  Expression
	Value Expression
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
	Token       Token
	Expressions []Expression
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

type IfStatement struct {
	Token       Token
	Condition   Expression
	Consequence Statement
	Alternative Statement
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
	Token     Token
	Body      []Statement
	Condition Expression
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
	Token      Token
	Statements []Statement
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
	Token Token
}

func (bs *BreakStatement) statementNode() {
	// read at line 72 :)

}
func (bs *BreakStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BreakStatement) String() string       { return bs.TokenLiteral() + " #" }

type ContinueStatement struct {
	Token Token
}

func (cs *ContinueStatement) statementNode() {

	// read at line 72 :)
}
func (cs *ContinueStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *ContinueStatement) String() string       { return cs.TokenLiteral() + " #" }

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

type UseStatement struct {
	Token    Token // The FRG_Use token
	Filename *StringLiteral
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
}                                                    // Mark as Statement node
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type Identifier struct {
	Token Token // the IDENT token
	Value string
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
	Token Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {
	// mark as expression node
	// read line 362 :))
}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// RealLiteral represents a real number literal.
type RealLiteral struct {
	Token Token
	Value float64
}

func (rl *RealLiteral) expressionNode() {
	// mark as expression node
	// read line 362 :))
}
func (rl *RealLiteral) TokenLiteral() string { return rl.Token.Literal }
func (rl *RealLiteral) String() string       { return rl.Token.Literal }

// StringLiteral represents a string literal.
type StringLiteral struct {
	Token Token
	Value string
}

func (sl *StringLiteral) expressionNode() {
	// mark as expression node
	// read line 362 :))
}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return "\"" + sl.Value + "\"" }

// PrefixExpression represents a prefix operator expression (e.g., -5).
type PrefixExpression struct {
	Token    Token // The prefix token, e.g. -
	Operator string
	Right    Expression
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
	Token    Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
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
	Token    Token // the '{' token
	Elements []Expression
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
	Token Token // the '[' token
	Size  Expression
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

type CallExpression struct {
	Token     Token // The ( token
	Function  Expression
	Arguments []Expression
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
// parser : after implements all interfaces for each statement and expression
// =============================================================================
const (
	_ int = iota
	LOWEST
	EQUALS      // ==, !=
	LESSGREATER // >, <, >=, <=
	SUM         // +
	PRODUCT     // *
	INDEX       // []
	CALL        // function()
	PREFIX      // -X
)

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
	prefixParseFn func() Expression
	infixParseFn  func(Expression) Expression
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

	p.registerInfix(TokenPlus, p.parseInfixExpression)
	p.registerInfix(TokenMinus, p.parseInfixExpression)
	p.registerInfix(TokenSlash, p.parseInfixExpression)
	p.registerInfix(TokenAsterisk, p.parseInfixExpression)
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

// each token has it's own parsing method
func (p *Parser) parseStatement() Statement {
	switch p.currentToken.Type {
	case TokenFRGInt, TokenFRGReal, TokenFRGStrg:
		return p.parseDeclarationStatement()
	case TokenFRGFn:
		return p.parseFunctionDeclarationStatement()
	case TokenFRGPrint:
		return p.parsePrintStatement()
	case TokenFRGInput:
		return p.parseInputStatement()
	case TokenIf:
		return p.parseIfStatement()
	case TokenRepeat:
		return p.parseRepeatStatement()
	case TokenBegin:
		return p.parseBlockStatement()
	case TokenBreak:
		return p.parseBreakStatement()
	case TokenContinue:
		return p.parseContinueStatement()
	case TokenFRGUse:
		return p.parseUseStatementAndInclude()
	case TokenIdentifier:
		expr := p.parseExpression(LOWEST)
		if p.peekTokenIs(TokenAssign) {
			stmt := &AssignmentStatement{Left: expr}
			p.nextToken() // consume :=
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

func (p *Parser) parseDeclarationStatement() *DeclarationStatement {
	stmt := &DeclarationStatement{Token: p.currentToken}
	stmt.Identifiers = []*Identifier{}

	if p.peekTokenIs(TokenLBracket) {
		p.nextToken()
		if !p.expectPeek(TokenRBracket) {
			return nil
		}
		stmt.IsArray = true
	}

	if !p.expectPeek(TokenIdentifier) {
		return nil
	}

	stmt.Identifiers = append(stmt.Identifiers, &Identifier{Token: p.currentToken, Value: p.currentToken.Literal})

	for p.peekTokenIs(TokenComma) {
		p.nextToken()
		if !p.expectPeek(TokenIdentifier) {
			return nil
		}
		stmt.Identifiers = append(stmt.Identifiers, &Identifier{Token: p.currentToken, Value: p.currentToken.Literal})
	}

	if !p.expectPeek(TokenHash) {
		return nil
	}

	return stmt
}

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

	p.nextToken()

	expr := p.parseExpression(LOWEST)
	stmt.Expressions = append(stmt.Expressions, expr)

	for p.peekTokenIs(TokenComma) {
		p.nextToken()
		p.nextToken()
		expr := p.parseExpression(LOWEST)
		stmt.Expressions = append(stmt.Expressions, expr)
	}

	if !p.expectPeek(TokenHash) {
		return nil
	}

	return stmt
}

func (p *Parser) parseIfStatement() *IfStatement {
	stmt := &IfStatement{Token: p.currentToken}

	if !p.expectPeek(TokenLBracket) {
		return nil
	}
	p.nextToken()

	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(TokenRBracket) {
		return nil
	}

	p.nextToken()

	stmt.Consequence = p.parseStatement()

	if p.peekTokenIs(TokenElse) {
		p.nextToken() // kill else
		p.nextToken()
		stmt.Alternative = p.parseStatement()

	}

	return stmt
}

func (p *Parser) parseRepeatStatement() *RepeatStatement {
	stmt := &RepeatStatement{Token: p.currentToken}
	stmt.Body = []Statement{}

	p.nextToken()

	for !p.currentTokenIs(TokenUntil) && !p.currentTokenIs(TokenEOF) {
		s := p.parseStatement()
		if s != nil {
			stmt.Body = append(stmt.Body, s)
		}
		p.nextToken()
	}

	if !p.currentTokenIs(TokenUntil) {
		p.peekError(TokenUntil)
		return nil
	}

	if !p.expectPeek(TokenLBracket) {
		return nil
	}
	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)
	if !p.expectPeek(TokenRBracket) {
		return nil
	}
	return stmt
}

func (p *Parser) parseBlockStatement() *BlockStatement {
	block := &BlockStatement{Token: p.currentToken}
	block.Statements = []Statement{}

	p.nextToken()

	for !p.currentTokenIs(TokenEnd) && !p.currentTokenIs(TokenEOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	if !p.currentTokenIs(TokenEnd) {
		p.errors = append(p.errors, fmt.Sprintf("unterminated block statement, expected End, got %s", TokenToString(p.currentToken.Type)))
	}

	return block
}

func (p *Parser) parseBreakStatement() *BreakStatement {
	stmt := &BreakStatement{Token: p.currentToken}
	if !p.expectPeek(TokenHash) {
		return nil
	}
	return stmt
}

func (p *Parser) parseContinueStatement() *ContinueStatement {
	stmt := &ContinueStatement{Token: p.currentToken}
	if !p.expectPeek(TokenHash) {
		return nil
	}
	return stmt
}

// parseExpression with stack algorithm implementation (Shunting-yard like algorithm)
func (p *Parser) parseExpression(precedence int) Expression {
	// Stack to hold expressions and operators
	var outputStack []Expression
	var operatorStack []struct {
		token      Token
		precedence int
		infixFn    infixParseFn
	}

	// Parse the first operand (prefix expression)
	prefix := p.prefixParseFns[p.currentToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.currentToken.Type)
		return nil
	}
	outputStack = append(outputStack, prefix())

	// Continue parsing while there are tokens and precedence allows
	for precedence < p.peekPrecedence() {
		// Get the infix parser function for the current peek token
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			break // No infix function, stop parsing
		}

		// Get precedence of the current operator
		nextPrecedence := p.peekPrecedence()

		// Process operators on the stack based on precedence
		for len(operatorStack) > 0 &&
			operatorStack[len(operatorStack)-1].precedence >= nextPrecedence {

			// Pop the operator
			op := operatorStack[len(operatorStack)-1]
			operatorStack = operatorStack[:len(operatorStack)-1]

			// Pop the right operand
			if len(outputStack) == 0 {
				p.errors = append(p.errors, "missing operand for operator")
				return nil
			}
			right := outputStack[len(outputStack)-1]
			outputStack = outputStack[:len(outputStack)-1]

			// Pop the left operand
			if len(outputStack) == 0 {
				p.errors = append(p.errors, "missing left operand for operator")
				return nil
			}
			left := outputStack[len(outputStack)-1]
			outputStack = outputStack[:len(outputStack)-1]

			// Create the infix expression and push back to output
			infixExpr := &InfixExpression{
				Token:    op.token,
				Left:     left,
				Operator: op.token.Literal,
				Right:    right,
			}
			outputStack = append(outputStack, infixExpr)
		}

		// Move to the operator token
		p.nextToken()

		// Push the operator to the operator stack
		operatorStack = append(operatorStack, struct {
			token      Token
			precedence int
			infixFn    infixParseFn
		}{
			token:      p.currentToken,
			precedence: nextPrecedence,
			infixFn:    infix,
		})

		// Parse the right operand
		p.nextToken() // Move to the operand
		rightOperand := p.parseExpression(nextPrecedence)
		outputStack = append(outputStack, rightOperand)
	}

	// Process remaining operators on the stack
	for len(operatorStack) > 0 {
		// Pop the operator
		op := operatorStack[len(operatorStack)-1]
		operatorStack = operatorStack[:len(operatorStack)-1]

		// Pop the right operand
		if len(outputStack) == 0 {
			p.errors = append(p.errors, "missing operand for operator")
			return nil
		}
		right := outputStack[len(outputStack)-1]
		outputStack = outputStack[:len(outputStack)-1]

		// Pop the left operand
		if len(outputStack) == 0 {
			p.errors = append(p.errors, "missing left operand for operator")
			return nil
		}
		left := outputStack[len(outputStack)-1]
		outputStack = outputStack[:len(outputStack)-1]

		// Create the infix expression and push back to output
		infixExpr := &InfixExpression{
			Token:    op.token,
			Left:     left,
			Operator: op.token.Literal,
			Right:    right,
		}
		outputStack = append(outputStack, infixExpr)
	}

	// Final result should be a single expression on the output stack
	if len(outputStack) == 0 {
		return nil
	}
	return outputStack[0]
}

func (p *Parser) parseIdentifier() Expression {
	return &Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

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

func (p *Parser) parsePrefixExpression() Expression {
	expression := &PrefixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

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

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

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
