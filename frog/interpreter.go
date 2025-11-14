package frog

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Environment struct {
	store map[string]Object
}

// constructor
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
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
	case *ArrayLiteral:
		return evalArrayLiteral(node, env)
	case *ArraySizeLiteral:
		return evalArraySizeLiteral(node, env)
	case *IndexExpression:
		return evalIndexExpression(node, env)
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
	case *InputStatement:
		return evalInputStatement(node, env)
	case *Boolean:
		if node.Value {
			return TRUE
		} else {
			return FALSE
		}
	case *BreakStatement:
		return BREAK
	case *ContinueStatement:
		return CONTINUE
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
			if result == BREAK {
				return nil
			}
			if result == CONTINUE {
				break
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
		result := Eval(is.Consequence, env)
		if result != nil && (result.Type() == BREAK_OBJ || result.Type() == CONTINUE_OBJ) {
			return result
		}
		return nil
	} else if is.Alternative != nil {
		result := Eval(is.Alternative, env)
		if result != nil && (result.Type() == BREAK_OBJ || result.Type() == CONTINUE_OBJ) {
			return result
		}
		return nil
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
			if rt == "ERROR" || rt == BREAK_OBJ || rt == CONTINUE_OBJ {
				return result
			}
		}
	}
	return result
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
			return &Real{Value: float64(leftVal) / float64(rightVal)}
		}
	case "%":
		if rightVal == 0 {
			return newError(node.Token.Line, node.Token.Column, "ERROR: u can't divis per zero")
		} else {
			return &Int{Value: leftVal % rightVal}
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
	case "%":
		if rightVal == 0 {
			return newError(node.Token.Line, node.Token.Column, "ERROR: u can't divis per zero")
		} else {
			return &Real{Value: math.Mod(leftVal, rightVal)}
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

func evalIdentifier(node *Identifier, env *Environment) Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	return newError(node.Token.Line, node.Token.Column, "identifier not found: "+node.Value)
}

func evalDeclarationStatement(node *DeclarationStatement, env *Environment) Object {
	for _, ident := range node.Identifiers {
		if node.IsArray {
			env.Set(ident.Value, &Array{Elements: []Object{}})
		} else {
			env.Set(ident.Value, nil)
		}
	}
	return nil
}

func evalAssignmentStatement(node *AssignmentStatement, env *Environment) Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}
	return evalAssignmentToExpression(node.Left, val, env)
}

func evalAssignmentToExpression(left Expression, val Object, env *Environment) Object {
	switch l := left.(type) {
	case *Identifier:
		if _, ok := env.Get(l.Value); !ok {
			return newError(l.Token.Line, l.Token.Column, "cannot assign to undeclared identifier: %s", l.Value)
		}
		env.Set(l.Value, val)
		return nil
	case *IndexExpression:
		return evalIndexAssignment(l, val, env)
	default:
		return newError(0, 0, "cannot assign to %T", left)
	}
}

func evalIndexAssignment(node *IndexExpression, val Object, env *Environment) Object {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}
	index := Eval(node.Index, env)
	if isError(index) {
		return index
	}
	switch {
	case left.Type() == ARRAY_OBJ && index.Type() == INTEGER_OBJ:
		array := left.(*Array)
		idx := index.(*Int).Value
		if idx < 0 {
			return newError(node.Token.Line, node.Token.Column, "index out of bounds: %d", idx)
		}
		// Extend array if necessary
		for int64(len(array.Elements)) <= idx {
			array.Elements = append(array.Elements, &Int{Value: 0})
		}
		array.Elements[idx] = val
		return nil
	default:
		return newError(node.Token.Line, node.Token.Column, "cannot assign to index: %s[%s]", left.Type(), index.Type())
	}
}

func evalPrintStatement(node *PrintStatement, env *Environment) Object {
	for _, expr := range node.Expressions {
		val := Eval(expr, env)
		if val != nil {
			fmt.Print(val.Inspect())
		}
	}
	return nil
}

func evalInputStatement(node *InputStatement, env *Environment) Object {
	reader := bufio.NewReader(os.Stdin)

	for _, expr := range node.Expressions {
		ident, ok := expr.(*Identifier)
		if !ok {
			return newError(node.Token.Line, node.Token.Column, "input statement expects identifiers")
		}

		if _, exists := env.Get(ident.Value); !exists {
			return newError(ident.Token.Line, ident.Token.Column, "cannot input to undeclared identifier: %s", ident.Value)
		}

		input, err := reader.ReadString('\n')
		if err != nil {
			return newError(ident.Token.Line, ident.Token.Column, "error reading input: %v", err)
		}

		input = input[:len(input)-1]

		if intVal, err := strconv.ParseInt(input, 10, 64); err == nil {
			env.Set(ident.Value, &Int{Value: intVal})
		} else if realVal, err := strconv.ParseFloat(input, 64); err == nil {
			env.Set(ident.Value, &Real{Value: realVal})
		} else {
			env.Set(ident.Value, &String{Value: input})
		}
	}

	return nil
}

func evalArrayLiteral(node *ArrayLiteral, env *Environment) Object {
	elements := []Object{}
	for _, el := range node.Elements {
		evaluated := Eval(el, env)
		if isError(evaluated) {
			return evaluated
		}
		elements = append(elements, evaluated)
	}
	return &Array{Elements: elements}
}

func evalArraySizeLiteral(node *ArraySizeLiteral, env *Environment) Object {
	sizeObj := Eval(node.Size, env)
	if isError(sizeObj) {
		return sizeObj
	}
	if sizeObj.Type() != INTEGER_OBJ {
		return newError(node.Token.Line, node.Token.Column, "array size must be integer")
	}
	size := sizeObj.(*Int).Value
	if size < 0 {
		return newError(node.Token.Line, node.Token.Column, "array size cannot be negative")
	}
	elements := make([]Object, size)
	for i := int64(0); i < size; i++ {
		elements[i] = &Int{Value: 0}
	}
	return &Array{Elements: elements}
}

func evalIndexExpression(node *IndexExpression, env *Environment) Object {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}
	index := Eval(node.Index, env)
	if isError(index) {
		return index
	}
	return evalIndexExpressionWithObjects(node, left, index)
}

func evalIndexExpressionWithObjects(node *IndexExpression, left, index Object) Object {
	switch {
	case left.Type() == ARRAY_OBJ && index.Type() == INTEGER_OBJ:
		array := left.(*Array)
		idx := index.(*Int).Value
		max := int64(len(array.Elements) - 1)
		if idx < 0 || idx > max {
			return newError(node.Token.Line, node.Token.Column, "index out of bounds: %d", idx)
		}
		return array.Elements[idx]
	case left.Type() == STRING_OBJ && index.Type() == INTEGER_OBJ:
		str := left.(*String).Value
		idx := index.(*Int).Value
		if idx < 0 || idx >= int64(len(str)) {
			return newError(node.Token.Line, node.Token.Column, "index out of bounds: %d", idx)
		}
		return &String{Value: string(str[idx])}
	default:
		return newError(node.Token.Line, node.Token.Column, "index operator not supported: %s[%s]", left.Type(), index.Type())
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
