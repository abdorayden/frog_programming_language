package frog

import (
	"fmt"
)

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

type Environment struct {
	store map[string]Object
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

func Eval(node Node, env *Environment) Object {
	switch node := node.(type) {
	case *Program:
		return evalProgram(node, env)
	case *ExpressionStatement:
		return Eval(node.Expression, env)
	case *IntegerLiteral:
		return &Int{Value: node.Value}
	case *RealLiteral:
		return &Real{Value: node.Value}
	case *StringLiteral:
		return &String{Value: node.Value}
	case *PrefixExpression:
		right := Eval(node.Right, env)
		return evalPrefixExpression(node, right)
	case *InfixExpression:
		left := Eval(node.Left, env)
		right := Eval(node.Right, env)
		return evalInfixExpression(node, left, right)
	case *DeclarationStatement:
		return evalDeclarationStatement(node, env)
	case *AssignmentStatement:
		return evalAssignmentStatement(node, env)
	case *Identifier:
		return evalIdentifier(node, env)
	case *RepeatStatement:
		return evalRepeatStatement(node, env)
	case *IfStatement:
		return evalIfStatement(node, env)
	case *BlockStatement:
		return evalBlockStatement(node, env)
	case *PrintStatement:
		return evalPrintStatement(node, env)
	case *Boolean:
		if node.Value {
			return TRUE
		} else {
			return FALSE
		}
	}
	return nil
}

func evalRepeatStatement(rs *RepeatStatement, env *Environment) Object {
	for {
		for _, statement := range rs.Body {
			result := Eval(statement, env)
			if isError(result) {
				return result
			}
		}
		condition := Eval(rs.Condition, env)
		if isError(condition) {
			return condition
		}
		if isTruthy(condition) {
			break
		}
	}
	return nil
}

func evalIfStatement(is *IfStatement, env *Environment) Object {
	condition := Eval(is.Condition, env)
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		return Eval(is.Consequence, env)
	} else if is.Alternative != nil {
		return Eval(is.Alternative, env)
	} else {
		return nil
	}
}

func evalBlockStatement(block *BlockStatement, env *Environment) Object {
	var result Object
	for _, statement := range block.Statements {
		result = Eval(statement, env)
		if result != nil {
			rt := result.Type()
			if rt == "ERROR" {
				return result
			}
		}
	}
	return result
}

func isTruthy(obj Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func evalProgram(program *Program, env *Environment) Object {
	var result Object
	for _, statement := range program.Statements {
		result = Eval(statement, env)
	}
	return result
}

func evalPrefixExpression(node *PrefixExpression, right Object) Object {
	switch node.Operator {
	case "-":
		return evalMinusPrefixOperatorExpression(node, right)
	default:
		return newError(node.Token.Line, node.Token.Column, "unknown operator: %s%s", node.Operator, right.Type())
	}
}

func evalMinusPrefixOperatorExpression(node *PrefixExpression, right Object) Object {
	if right.Type() == INTEGER_OBJ {
		value := right.(*Int).Value
		return &Int{Value: -value}
	}
	if right.Type() == REAL_OBJ {
		value := right.(*Real).Value
		return &Real{Value: -value}
	}
	return newError(node.Token.Line, node.Token.Column, "unknown operator: -%s", right.Type())
}

func evalInfixExpression(node *InfixExpression, left, right Object) Object {
	switch {
	case left.Type() == INTEGER_OBJ && right.Type() == INTEGER_OBJ:
		return evalIntegerInfixExpression(node, left, right)
	case left.Type() == REAL_OBJ && right.Type() == REAL_OBJ:
		return evalRealInfixExpression(node, left, right)
	case left.Type() == STRING_OBJ && right.Type() == STRING_OBJ:
		return evalStringInfixExpression(node, left, right)
	case left.Type() != right.Type():
		return newError(node.Token.Line, node.Token.Column, "type mismatch: %s %s %s", left.Type(), node.Operator, right.Type())
	default:
		return newError(node.Token.Line, node.Token.Column, "unknown operator: %s %s %s", left.Type(), node.Operator, right.Type())
	}
}

func evalIntegerInfixExpression(node *InfixExpression, left, right Object) Object {
	leftVal := left.(*Int).Value
	rightVal := right.(*Int).Value
	switch node.Operator {
	case "+":
		return &Int{Value: leftVal + rightVal}
	case "-":
		return &Int{Value: leftVal - rightVal}
	case "*":
		return &Int{Value: leftVal * rightVal}
	case "/":
		if rightVal == 0 {
			return newError(node.Token.Line, node.Token.Column, "ERROR: u can't divis per zero")
		} else {
			return &Int{Value: leftVal / rightVal}
		}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError(node.Token.Line, node.Token.Column, "unknown operator: %s %s %s", left.Type(), node.Operator, right.Type())
	}
}

func evalRealInfixExpression(node *InfixExpression, left, right Object) Object {
	leftVal := left.(*Real).Value
	rightVal := right.(*Real).Value
	switch node.Operator {
	case "+":
		return &Real{Value: leftVal + rightVal}
	case "-":
		return &Real{Value: leftVal - rightVal}
	case "*":
		return &Real{Value: leftVal * rightVal}
	case "/":
		if rightVal == 0 {
			return newError(node.Token.Line, node.Token.Column, "ERROR: u can't divis per zero")
		} else {
			return &Real{Value: leftVal / rightVal}
		}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError(node.Token.Line, node.Token.Column, "unknown operator: %s %s %s", left.Type(), node.Operator, right.Type())
	}
}

func evalStringInfixExpression(node *InfixExpression, left, right Object) Object {
	leftVal := left.(*String).Value
	rightVal := right.(*String).Value
	if node.Operator != "+" {
		return newError(node.Token.Line, node.Token.Column, "unknown operator: %s %s %s", left.Type(), node.Operator, right.Type())
	}
	return &String{Value: leftVal + rightVal}
}

func nativeBoolToBooleanObject(input bool) *Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalIdentifier(node *Identifier, env *Environment) Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	return newError(node.Token.Line, node.Token.Column, "identifier not found: "+node.Value)
}

func evalDeclarationStatement(node *DeclarationStatement, env *Environment) Object {
	for _, ident := range node.Identifiers {
		env.Set(ident.Value, nil) // Initialize with nil
	}
	return nil
}

func evalAssignmentStatement(node *AssignmentStatement, env *Environment) Object {
	if _, ok := env.Get(node.Identifier.Value); !ok {
		return newError(node.Identifier.Token.Line, node.Identifier.Token.Column, "cannot assign to undeclared identifier: %s", node.Identifier.Value)
	}
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}
	env.Set(node.Identifier.Value, val)
	return nil
}

func evalPrintStatement(node *PrintStatement, env *Environment) Object {
	for _, expr := range node.Expressions {
		val := Eval(expr, env)
		if val != nil {
			fmt.Print(val.Inspect())
		}
	}
	fmt.Println()
	return nil
}

func newError(line, col int, format string, a ...interface{}) *Error {
	return &Error{
		Message: fmt.Sprintf(format, a...),
		Line:    line,
		Col:     col,
	}
}

type Error struct {
	Message string
	Line    int
	Col     int
}

func (e *Error) Type() ObjectType { return "ERROR" }
func (e *Error) Inspect() string {
	return fmt.Sprintf("ERROR: %s (line %d, col %d)", e.Message, e.Line, e.Col)
}

func isError(obj Object) bool {
	if obj != nil {
		return obj.Type() == "ERROR"
	}
	return false
}
