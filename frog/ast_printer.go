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
		PrintAST(n.Identifier, childPrefix, false)
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
	case *Boolean:
		fmt.Printf("Boolean: %t\n", n.Value)
	default:
		fmt.Printf("Unknown Node Type: %T\n", n)
	}
}
