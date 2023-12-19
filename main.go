package main

import (
	"fmt"
	"poc/aql/core"
	"poc/aql/grammar"

	"github.com/alecthomas/participle/v2"
)

func main() {
	parser, err := participle.Build[grammar.Query](
		participle.Unquote("String"),
		participle.Union[grammar.Value](grammar.String{}, grammar.Number{}),
	)
	if err != nil {
		panic(err)
	}

	alerceQuery, err := parser.ParseString("", "GET Object oid=\"ZTF20aaelulu\" from_survey=\"ZTF\"")
	if err != nil {
		panic(err)
	}
	result, err := core.MakeQuery(alerceQuery)
	if err != nil {
		panic(err)
	}
	strResult := string(result)
	fmt.Printf("result is %+v\n", strResult)
}
