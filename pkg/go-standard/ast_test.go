package go_standard

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestAst(t *testing.T) {
	// scr is the input for which we want to inspect the AST
	src := `
package p
const c = 1.0
var X = f(3.14)*2 + c
`

	// create the ast by parsing src
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}

	// Inspect the AST and print all identifiers and literals
	ast.Inspect(f, func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		case *ast.BasicLit:
			s = x.Value
		case *ast.Ident:
			s = x.Name
		}
		if s != "" {
			t.Logf("%s:\t%s\n", fset.Position(n.Pos()), s)
		}

		return true
	})

}

func TestAst2(t *testing.T) {
	src := `
package main

var a int = 3
`

	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", src, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", f)
}
