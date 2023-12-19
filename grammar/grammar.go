package grammar

import (
	"fmt"
)

type Query struct {
	Verb    string        `@Ident`
	Entity  string        `@Ident`
	Filters []*Expression `@@*`
}

type Value interface{ value() }

type String struct {
	String string `@String`
}

func (String) value() {}

type Number struct {
	Number float64 `@Float | @Int`
}

func (Number) value() {}

type Expression struct {
	Field string `@Ident`
	Op    string `@( "|" "|" | "&" "&" | "!" "=" | ("!"|"="|"<"|">") "="?  )`
	Value Value  `@@`
}

func (ee *Expression) String() string {
	return fmt.Sprintf("%s %s %v", ee.Field, ee.Op, ee.Value)
}
