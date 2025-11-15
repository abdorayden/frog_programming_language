// Copyright (C) by abdenour souane
// you have a right to modify it upgrade it or do whatever you want
// but u have to keep my name on it

package frog

import (
	"fmt"
)

func PrintAST(node Node, prefix string, isLast bool) {
	var connector, childPrefix string
	if isLast {
		connector = "└── "
		childPrefix = prefix + "    "
	} else {
		connector = "├── "
		childPrefix = prefix + "│   "
	}

	fmt.Printf("%s%s", prefix, connector)

	switch n := node.(type) {
	case *Program:
		fmt.Println("Program:")
		for i, stmt := range n.Statements {
			PrintAST(stmt, childPrefix, i == len(n.Statements)-1)
		}
	case *DeclarationStatement:
		fmt.Printf("DeclarationStatement: Type=%s\n", n.TokenLiteral())
		for i, ident := range n.Identifiers {
			PrintAST(ident, childPrefix, i == len(n.Identifiers)-1)
		}
	case *AssignmentStatement:
		fmt.Println("AssignmentStatement:")
		PrintAST(n.Left, childPrefix, false)
		PrintAST(n.Value, childPrefix, true)
	case *PrintStatement:
		fmt.Println("PrintStatement:")
		for i, expr := range n.Expressions {
			PrintAST(expr, childPrefix, i == len(n.Expressions)-1)
		}
	case *IfStatement:
		fmt.Println("IfStatement:")
		PrintAST(n.Condition, childPrefix, false)
		PrintAST(n.Consequence, childPrefix, n.Alternative == nil)
		if n.Alternative != nil {
			PrintAST(n.Alternative, childPrefix, true)
		}
	case *RepeatStatement:
		fmt.Println("RepeatStatement:")
		PrintAST(n.Condition, childPrefix, false)
		for i, stmt := range n.Body {
			PrintAST(stmt, childPrefix, i == len(n.Body)-1)
		}
	case *BlockStatement:
		fmt.Println("BlockStatement:")
		for i, stmt := range n.Statements {
			PrintAST(stmt, childPrefix, i == len(n.Statements)-1)
		}
	case *FunctionDeclarationStatement:
		fmt.Printf("FunctionDeclarationStatement: %s\n", n.Name.Value)
		PrintAST(n.Body, childPrefix, true)
	case *ExpressionStatement:
		fmt.Println("ExpressionStatement:")
		PrintAST(n.Expression, childPrefix, true)
	case *Identifier:
		fmt.Printf("Identifier: %s\n", n.Value)
	case *IntegerLiteral:
		fmt.Printf("IntegerLiteral: %d\n", n.Value)
	case *RealLiteral:
		fmt.Printf("RealLiteral: %f\n", n.Value)
	case *StringLiteral:
		fmt.Printf("StringLiteral: %q\n", n.Value)
	case *PrefixExpression:
		fmt.Printf("PrefixExpression: Operator=%s\n", n.Operator)
		PrintAST(n.Right, childPrefix, true)
	case *InfixExpression:
		fmt.Printf("InfixExpression: Operator=%s\n", n.Operator)
		PrintAST(n.Left, childPrefix, false)
		PrintAST(n.Right, childPrefix, true)
	case *ArrayLiteral:
		fmt.Println("ArrayLiteral:")
		for i, el := range n.Elements {
			PrintAST(el, childPrefix, i == len(n.Elements)-1)
		}
	case *IndexExpression:
		fmt.Println("IndexExpression:")
		PrintAST(n.Left, childPrefix, false)
		PrintAST(n.Index, childPrefix, true)
	case *CallExpression:
		fmt.Println("CallExpression:")
		PrintAST(n.Function, childPrefix, false)
		for i, arg := range n.Arguments {
			PrintAST(arg, childPrefix, i == len(n.Arguments)-1)
		}
	case *Boolean:
		fmt.Printf("Boolean: %t\n", n.Value)
	default:
		fmt.Printf("Unknown Node Type: %T\n", n)
	}
}
