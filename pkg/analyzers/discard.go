package analyzers

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

var DiscardAnalyzer = analysis.Analyzer{
	Name: "discardedreturn",
	Doc:  "todo",
	Run:  runDiscard,
}

func discardCheckCall(pass *analysis.Pass, call *ast.CallExpr) bool {
	typ := pass.TypesInfo.TypeOf(call.Fun)
	if typ == nil {
		return true
	}

	f, ok := typ.(*types.Signature)
	if !ok {
		return true
	}
	if f.Results() != nil && f.Results().Len() != 0 {
		pass.Reportf(call.Pos(), "call discards return value "+
			"(hint: assign to an underscore to ignore)")
		// 	funcDecl.Name.Name, funcDecl.Name.Name)
	}
	return true
}

func runDiscard(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		switch node := node.(type) {
		case *ast.ExprStmt:
			call, ok := node.X.(*ast.CallExpr)
			if !ok {
				return true
			}
			return discardCheckCall(pass, call)
		case *ast.DeferStmt:
			return discardCheckCall(pass, node.Call)
		}
		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
