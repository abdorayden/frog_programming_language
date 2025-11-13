package frog

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	REAL_OBJ    = "REAL"
	STRING_OBJ  = "STRING"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Int struct {
	Value int64
}

func (i *Int) Type() ObjectType { return INTEGER_OBJ }
func (i *Int) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

type Real struct {
	Value float64
}

func (r *Real) Type() ObjectType { return REAL_OBJ }
func (r *Real) Inspect() string  { return fmt.Sprintf("%f", r.Value) }

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

type Boolean struct {
	Token Token
	Value bool
}

func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}
func (b *Boolean) String() string   { return b.Token.Literal }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) expressionNode()  {} // Mark as Expression node

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

var (
	NULL  = &Null{}
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)
