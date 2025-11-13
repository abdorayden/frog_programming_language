package frog

import (
	"bytes"
	"fmt"
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

// DeclarationStatement represents a variable declaration, e.g., FRG_Int i, j, k #
type DeclarationStatement struct {
	Token       Token // The type token (FRG_Int, FRG_Real, FRG_Strg)
	Identifiers []*Identifier
}

func (ds *DeclarationStatement) statementNode()       {}
func (ds *DeclarationStatement) TokenLiteral() string { return ds.Token.Literal }
func (ds *DeclarationStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ds.TokenLiteral() + " ")
	for i, ident := range ds.Identifiers {
		out.WriteString(ident.String())
		if i < len(ds.Identifiers)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(" #")
	return out.String()
}

type AssignmentStatement struct {
	Token      Token
	Identifier *Identifier
	Value      Expression
}

func (as *AssignmentStatement) statementNode()       {}
func (as *AssignmentStatement) TokenLiteral() string { return as.Token.Literal }
func (as *AssignmentStatement) String() string {
	var out bytes.Buffer
	out.WriteString(as.Identifier.String())
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

func (ps *PrintStatement) statementNode()       {}
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

type IfStatement struct {
	Token       Token
	Condition   Expression
	Consequence Statement
	Alternative Statement
}

func (is *IfStatement) statementNode()       {}
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

func (rs *RepeatStatement) statementNode()       {}
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

func (bs *BlockStatement) statementNode()       {}
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

// ExpressionStatement is a statement that consists of a single expression.
type ExpressionStatement struct {
	Token      Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {} // Mark as Statement node
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// Identifier represents an identifier.
type Identifier struct {
	Token Token // the IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {} // Mark as Expression node
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// IntegerLiteral represents an integer literal.
type IntegerLiteral struct {
	Token Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {} // Mark as Expression node
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// RealLiteral represents a real number literal.
type RealLiteral struct {
	Token Token
	Value float64
}

func (rl *RealLiteral) expressionNode()      {} // Mark as Expression node
func (rl *RealLiteral) TokenLiteral() string { return rl.Token.Literal }
func (rl *RealLiteral) String() string       { return rl.Token.Literal }

// StringLiteral represents a string literal.
type StringLiteral struct {
	Token Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {} // Mark as Expression node
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return "\"" + sl.Value + "\"" }

// PrefixExpression represents a prefix operator expression (e.g., -5).
type PrefixExpression struct {
	Token    Token // The prefix token, e.g. -
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {} // Mark as Expression node
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

func (ie *InfixExpression) expressionNode()      {} // Mark as Expression node
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

// =============================================================================
// Parser
// =============================================================================

const (
	_ int = iota
	LOWEST
	EQUALS      // ==, !=
	LESSGREATER // >, <, >=, <=
	SUM         // +
	PRODUCT     // *
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
}

type (
	prefixParseFn func() Expression
	infixParseFn  func(Expression) Expression
)

type Parser struct {
	lexer  *Lexer
	errors []string

	currentToken Token
	peekToken    Token

	prefixParseFns map[TokenType]prefixParseFn
	infixParseFns  map[TokenType]infixParseFn
}

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

	p.infixParseFns = make(map[TokenType]infixParseFn)
	p.registerInfix(TokenPlus, p.parseInfixExpression)
	p.registerInfix(TokenMinus, p.parseInfixExpression)
	p.registerInfix(TokenSlash, p.parseInfixExpression)
	p.registerInfix(TokenAsterisk, p.parseInfixExpression)
	p.registerInfix(TokenEqual, p.parseInfixExpression)
	p.registerInfix(TokenNotEqual, p.parseInfixExpression)
	p.registerInfix(TokenLessThan, p.parseInfixExpression)
	p.registerInfix(TokenGreaterThan, p.parseInfixExpression)
	p.registerInfix(TokenLessEqual, p.parseInfixExpression)
	p.registerInfix(TokenGreaterEqual, p.parseInfixExpression)

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) IsThereAnyErrors() bool {
	return len(p.errors) != 0
}

func (p *Parser) peekError(t TokenType) {
	msg := fmt.Sprintf(
		"ERROR: expected next token to be %s, got %s instead (line %d, col %d)",
		TokenToString(t), TokenToString(p.peekToken.Type), p.peekToken.Line, p.peekToken.Column)
	p.errors = append(p.errors, msg)
}

func (p *Parser) registerPrefix(tokenType TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) ParseProgram() *Program {
	program := &Program{}
	program.Statements = []Statement{}

	if !p.expectCurrent(TokenFRGBegin) {
		p.errors = append(p.errors, "program must start with FRG_Begin")
		return program
	}

	p.nextToken()

	var end bool = false

	for !p.currentTokenIs(TokenEOF) {
		if p.currentTokenIs(TokenFRGEnd) {
			end = true
			break
		}
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	if !end {
		p.errors = append(p.errors, "program must end with FRG_End")
	}

	return program
}

func (p *Parser) parseStatement() Statement {
	switch p.currentToken.Type {
	case TokenFRGInt, TokenFRGReal, TokenFRGStrg:
		return p.parseDeclarationStatement()
	case TokenFRGPrint:
		return p.parsePrintStatement()
	case TokenIf:
		return p.parseIfStatement()
	case TokenRepeat:
		return p.parseRepeatStatement()
	case TokenBegin:
		return p.parseBlockStatement()
	case TokenIdentifier:
		if p.peekTokenIs(TokenAssign) {
			return p.parseAssignmentStatement()
		}
	}
	return nil
}

func (p *Parser) parseDeclarationStatement() *DeclarationStatement {
	stmt := &DeclarationStatement{Token: p.currentToken}
	stmt.Identifiers = []*Identifier{}

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

func (p *Parser) parseAssignmentStatement() *AssignmentStatement {
	stmt := &AssignmentStatement{
		Identifier: &Identifier{Token: p.currentToken, Value: p.currentToken.Literal},
	}

	p.nextToken()
	stmt.Token = p.currentToken

	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)

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
		p.nextToken()
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

func (p *Parser) parseExpression(precedence int) Expression {
	prefix := p.prefixParseFns[p.currentToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.currentToken.Type)
		return nil
	}
	leftExp := prefix()

	for precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}

	return leftExp
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

func (p *Parser) noPrefixParseFnError(t TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found (line %d, col %d)", TokenToString(t), p.currentToken.Line, p.currentToken.Column)
	p.errors = append(p.errors, msg)
}
