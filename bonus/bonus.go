https://powcoder.com
代写代考加微信 powcoder
Assignment Project Exam Help
Add WeChat powcoder
package bonus

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

// rewriteCalls should modify the passed AST
func rewriteCalls(node ast.Node) {
	//TODO Write the rewriteCalls function
	panic("TODO: implement this!")
}

func SimplifyParseAndEval(src string) string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	rewriteCalls(f)

	var buf bytes.Buffer
	format.Node(&buf, fset, f)
	return buf.String()
}
