https://powcoder.com
代写代考加微信 powcoder
Assignment Project Exam Help
Add WeChat powcoder
package main

import (
	"fmt"
	"hw2/expr"
)

func ParseAndEval(x string, y expr.Env) float64 {
	return 42.0
}

func main() {
	var result float64

	result = 2 + expr.ParseAndEval("1 + 2", expr.Env{})
	fmt.Printf("%d\n", result)

	result = expr.ParseAndEval("1 / 6", expr.Env{}) + 3.0
	fmt.Printf("%d\n", result)

	result = ParseAndEval("1 + 2", expr.Env{})
	fmt.Printf("%d\n", result)

	x := "1 + 2"
	result = expr.ParseAndEval(x, expr.Env{})
	fmt.Printf("%d\n", result)

	result = expr.ParseAndEval("1 + / / 2", expr.Env{})
	fmt.Printf("%d\n", result)
}
