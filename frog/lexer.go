package frog

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// TODO: add tables
// FRG_Int[] a#
//	a := {1,2,3}
// FRG_Real[] a#
//	a := {1.1,2.2,3,3}
// FRG_Strg[] a#
//	a := {"a","b","c"}
// TODO: add functions
//	FRG_Fn foo(FRG_Int a) : FRG_Int
//	Begin
//		foo := a ## return value
//	End
// TODO: add system modules
//	FRG_Use "core.ifrg"

// BUG: use undeclared ID not generated error

type TokenType int

const (
	TokenEOF TokenType = iota
	TokenIllegal

	TokenIdentifier
	TokenNumber
	TokenString

	TokenAssign   // :=
	TokenPlus     // +
	TokenMinus    // -
	TokenAsterisk // *
	TokenSlash    // /
	TokenModulo   // %

	TokenEqual        // ==
	TokenNotEqual     // !=
	TokenLessThan     // <
	TokenGreaterThan  // >
	TokenLessEqual    // <=
	TokenGreaterEqual // >=

	TokenAnd // &&
	TokenOr  // ||
	TokenNot // !

	TokenComma     // ,
	TokenSemicolon // ;
	TokenColon     // :
	TokenLParen    // (
	TokenRParen    // )
	TokenLBrace    // {
	TokenRBrace    // }
	TokenLBracket  // [
	TokenRBracket  // ]
	TokenHash      // #

	TokenFRGBegin
	TokenFRGEnd
	TokenFRGInt
	TokenFRGReal
	TokenFRGStrg
	TokenFRGPrint
	TokenFRGInput
	TokenIf
	TokenElse
	TokenBegin
	TokenEnd
	TokenRepeat
	TokenUntil
	TokenBreak
	TokenContinue
	TokenTrue
	TokenFalse
)

// map of keywords and thier tokens
var keywords = map[string]TokenType{
	"FRG_Begin": TokenFRGBegin,
	"FRG_End":   TokenFRGEnd,
	"FRG_Int":   TokenFRGInt,
	"FRG_Real":  TokenFRGReal,
	"FRG_Strg":  TokenFRGStrg,
	"FRG_Print": TokenFRGPrint,
	"FRG_Input": TokenFRGInput,
	"If":        TokenIf,
	"Else":      TokenElse,
	"Begin":     TokenBegin,
	"End":       TokenEnd,
	"Repeat":    TokenRepeat,
	"Until":     TokenUntil,
	"Break":     TokenBreak,
	"Continue":  TokenContinue,
	"True":      TokenTrue,
	"False":     TokenFalse,
}

type Token struct {
	// Type export field of export Token structure that spesified the token type
	Type TokenType
	// Literal the value of this kind of token
	Literal string
	// Line of the token
	Line int
	// Column of the token
	Column int
}

func (t Token) String() string {
	// String method that return formated string of the structure fields
	return fmt.Sprintf("Token{Type: %v, Literal: %q, Line: %d, Column: %d}", t.Type, t.Literal, t.Line, t.Column)
}

type Lexer struct {
	input        string
	position     int // current position
	readPosition int // next position

	// rune in Go is a data type that stores codes that represent Unicode characters.
	// Unicode is actually the collection of all possible characters present in the whole world.
	// In Unicode, each of these characters is assigned a unique number called the Unicode code point.
	// This code point is what we store in a rune data type.
	ch     rune
	line   int
	column int
}

// constructor
func NewLexer(input string) *Lexer {
	l := &Lexer{
		input:  input,
		line:   1,
		column: 0,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch, _ = utf8.DecodeRuneInString(l.input[l.readPosition:])
	}
	l.position = l.readPosition
	l.readPosition += utf8.RuneLen(l.ch)
	if l.ch == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	ch, _ := utf8.DecodeRuneInString(l.input[l.readPosition:])
	return ch
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	line := l.line
	column := l.column

	switch l.ch {
	case ':':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TokenAssign, Literal: string(ch) + string(l.ch), Line: line, Column: column}
		} else {
			tok = Token{Type: TokenColon, Literal: string(l.ch), Line: line, Column: column}
		}
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TokenEqual, Literal: string(ch) + string(l.ch), Line: line, Column: column}
		} else {
			tok = Token{Type: TokenIllegal, Literal: string(l.ch), Line: line, Column: column}
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TokenNotEqual, Literal: string(ch) + string(l.ch), Line: line, Column: column}
		} else {
			tok = Token{Type: TokenNot, Literal: string(l.ch), Line: line, Column: column}
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TokenLessEqual, Literal: string(ch) + string(l.ch), Line: line, Column: column}
		} else {
			tok = Token{Type: TokenLessThan, Literal: string(l.ch), Line: line, Column: column}
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TokenGreaterEqual, Literal: string(ch) + string(l.ch), Line: line, Column: column}
		} else {
			tok = Token{Type: TokenGreaterThan, Literal: string(l.ch), Line: line, Column: column}
		}
	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TokenAnd, Literal: string(ch) + string(l.ch), Line: line, Column: column}
		} else {
			tok = Token{Type: TokenIllegal, Literal: string(l.ch), Line: line, Column: column}
		}
	case '|':
		if l.peekChar() == '|' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TokenOr, Literal: string(ch) + string(l.ch), Line: line, Column: column}
		} else {
			tok = Token{Type: TokenIllegal, Literal: string(l.ch), Line: line, Column: column}
		}
	case '+':
		tok = Token{Type: TokenPlus, Literal: string(l.ch), Line: line, Column: column}
	case '-':
		tok = Token{Type: TokenMinus, Literal: string(l.ch), Line: line, Column: column}
	case '*':
		tok = Token{Type: TokenAsterisk, Literal: string(l.ch), Line: line, Column: column}
	case '/':
		tok = Token{Type: TokenSlash, Literal: string(l.ch), Line: line, Column: column}
	case '%':
		tok = Token{Type: TokenModulo, Literal: string(l.ch), Line: line, Column: column}
	case ',':
		tok = Token{Type: TokenComma, Literal: string(l.ch), Line: line, Column: column}
	case ';':
		tok = Token{Type: TokenSemicolon, Literal: string(l.ch), Line: line, Column: column}
	case '(':
		tok = Token{Type: TokenLParen, Literal: string(l.ch), Line: line, Column: column}
	case ')':
		tok = Token{Type: TokenRParen, Literal: string(l.ch), Line: line, Column: column}
	case '{':
		tok = Token{Type: TokenLBrace, Literal: string(l.ch), Line: line, Column: column}
	case '}':
		tok = Token{Type: TokenRBrace, Literal: string(l.ch), Line: line, Column: column}
	case '[':
		tok = Token{Type: TokenLBracket, Literal: string(l.ch), Line: line, Column: column}
	case ']':
		tok = Token{Type: TokenRBracket, Literal: string(l.ch), Line: line, Column: column}
	case '#':
		// Check for comment (##)
		if l.peekChar() == '#' {
			l.skipComment()
			return l.NextToken()
		} else {
			tok = Token{Type: TokenHash, Literal: string(l.ch), Line: line, Column: column}
		}
	case '"':
		tok.Type = TokenString
		tok.Literal = l.readString()
		tok.Line = line
		tok.Column = column
	case 0:
		tok.Literal = ""
		tok.Type = TokenEOF
		tok.Line = line
		tok.Column = column
	default:
		if isLetter(l.ch) {
			ident := l.readIdentifier()
			tokType := TokenIdentifier
			if kw, ok := keywords[ident]; ok {
				tokType = kw
			}
			return Token{Type: tokType, Literal: ident, Line: line, Column: column}
		} else if isDigit(l.ch) {
			return Token{Type: TokenNumber, Literal: l.readNumber(), Line: line, Column: column}
		} else {
			tok = Token{Type: TokenIllegal, Literal: string(l.ch), Line: line, Column: column}
		}
	}

	l.readChar()
	return tok
}

// PRIVATE methods (helpers)
func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	l.skipWhitespace()
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	if l.ch == '.' {
		l.readChar()
		for isDigit(l.ch) {
			l.readChar()
		}
	}

	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
		if l.ch == '\\' {
			l.readChar()
		}
	}
	str := l.input[position:l.position]
	str = strings.ReplaceAll(str, "\\n", "\n")
	str = strings.ReplaceAll(str, "\\t", "\t")
	str = strings.ReplaceAll(str, "\\\"", "\"")
	str = strings.ReplaceAll(str, "\\\\", "\\")
	return str
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func TokenToString(tokenType TokenType) string {
	switch tokenType {
	case TokenEOF:
		return "EOF"
	case TokenIllegal:
		return "ILLEGAL"
	case TokenIdentifier:
		return "IDENTIFIER"
	case TokenNumber:
		return "NUMBER"
	case TokenString:
		return "STRING"
	case TokenAssign:
		return "ASSIGN"
	case TokenPlus:
		return "PLUS"
	case TokenMinus:
		return "MINUS"
	case TokenAsterisk:
		return "ASTERISK"
	case TokenSlash:
		return "SLASH"
	case TokenModulo:
		return "MODULO"
	case TokenEqual:
		return "EQUAL"
	case TokenNotEqual:
		return "NOT_EQUAL"
	case TokenLessThan:
		return "LESS_THAN"
	case TokenGreaterThan:
		return "GREATER_THAN"
	case TokenLessEqual:
		return "LESS_EQUAL"
	case TokenGreaterEqual:
		return "GREATER_EQUAL"
	case TokenAnd:
		return "AND"
	case TokenOr:
		return "OR"
	case TokenNot:
		return "NOT"
	case TokenComma:
		return "COMMA"
	case TokenSemicolon:
		return "SEMICOLON"
	case TokenColon:
		return "COLON"
	case TokenLParen:
		return "LPAREN"
	case TokenRParen:
		return "RPAREN"
	case TokenLBrace:
		return "LBRACE"
	case TokenRBrace:
		return "RBRACE"
	case TokenLBracket:
		return "LBRACKET"
	case TokenRBracket:
		return "RBRACKET"
	case TokenHash:
		return "HASH"
	case TokenFRGBegin:
		return "FRG_BEGIN"
	case TokenFRGEnd:
		return "FRG_END"
	case TokenFRGInt:
		return "FRG_INT"
	case TokenFRGReal:
		return "FRG_REAL"
	case TokenFRGStrg:
		return "FRG_STRG"
	case TokenFRGPrint:
		return "FRG_PRINT"
	case TokenFRGInput:
		return "FRG_INPUT"
	case TokenIf:
		return "IF"
	case TokenElse:
		return "ELSE"
	case TokenBegin:
		return "BEGIN"
	case TokenEnd:
		return "END"
	case TokenRepeat:
		return "REPEAT"
	case TokenUntil:
		return "UNTIL"
	case TokenBreak:
		return "BREAK"
	case TokenContinue:
		return "CONTINUE"
	case TokenTrue:
		return "TRUE"
	case TokenFalse:
		return "FALSE"
	default:
		return "UNKNOWN"
	}
}

func (l *Lexer) GetAllTokens() []Token {
	tokens := []Token{}
	for {
		token := l.NextToken()
		tokens = append(tokens, token)
		if token.Type == TokenEOF {
			break
		}
	}
	return tokens
}

func (l *Lexer) Reset() {
	l.position = 0
	l.readPosition = 0
	l.line = 1
	l.column = 0
	l.readChar()
}
