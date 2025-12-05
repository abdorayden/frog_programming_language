// Copyright (C) by abdenour souane
// you have a right to modify it upgrade it or do whatever you want
// but u have to keep my name on it
package frog

import (
	"fmt"
)

func isError(obj Object) bool {
	if obj != nil {
		return obj.Type() == "ERROR"
	}
	return false
}

func newError(line, col int, format string, a ...interface{}) *Error {
	return &Error{
		Message: fmt.Sprintf(format, a...),
		Line:    line,
		Col:     col,
	}
}

func nativeBoolToBooleanObject(input bool) *Boolean {
	if input {
		return TRUE
	}
	return FALSE
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
	case TokenFRGFn:
		return "FRG_FN"
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
	case TokenFRGUse:
		return "FRG_USE"
	default:
		return "UNKNOWN"
	}
}
