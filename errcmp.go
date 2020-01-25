package errcmp

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name: "errcmp",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const Doc = "errcmp is ..."

func run(pass *analysis.Pass) (interface{}, error) {
	err := errors.New("error")
	if err != nil {
		fmt.Println("error!")
	}

	for _, f := range pass.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			if binary, ok := n.(*ast.BinaryExpr); ok {
				if binary.Op == token.NEQ {
					switch binary.X.(type) {
					case *ast.Ident:
						errName := binary.X.(*ast.Ident).Name
						if errName == "err" {
							pass.Reportf(binary.OpPos, "'err != ...' should be 'erros.Is(err, ...)'")
						}
					}
				}
			}
			return true
		})
	}
	return nil, nil
}
