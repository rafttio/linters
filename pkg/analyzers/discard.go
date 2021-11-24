package analyzers

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

var (
	DiscardAnalyzer = analysis.Analyzer{
		Name: "discardedreturn",
		Doc:  "checks for discarded return values",
		Run:  runDiscardAnalyzer,
	}
)

var (
	noCheckClosures bool
	noCheckCompound bool
	noCheckOpaque   bool
	checkBasic      bool
	checkErrors     bool
)

func init() {
	DiscardAnalyzer.Flags.BoolVar(&noCheckClosures, "no-closures", false, "ignore unused returned closures")
	DiscardAnalyzer.Flags.BoolVar(&noCheckCompound, "no-compound", false, "ignore unused structs")
	DiscardAnalyzer.Flags.BoolVar(&noCheckOpaque, "no-opaque", false, "ignore unused opaque (interface) types")
	DiscardAnalyzer.Flags.BoolVar(&checkBasic, "basic", false, "include basic types")
	DiscardAnalyzer.Flags.BoolVar(&checkErrors, "errors", false, "include returned errors")
}

func formatNode(pass *analysis.Pass, node ast.Node) string {
	writer := bytes.Buffer{}
	_ = printer.Fprint(&writer, pass.Fset, node)
	return writer.String()
}

func checkResultType(ty types.Type) bool {
	switch ty := ty.(type) {
	case *types.Interface:
		return !noCheckOpaque
	case *types.Basic:
		return checkBasic
	case *types.Signature:
		return !noCheckClosures
	case *types.Named:
		if ty.Obj().Name() == "error" {
			return checkErrors
		}
		return !noCheckOpaque
	default:
		return !noCheckCompound
	}
}

func checkResult(result *types.Tuple) bool {
	if result == nil || result.Len() == 0 {
		return false
	}

	for i := 0; i < result.Len(); i++ {
		if checkResultType(result.At(i).Type()) {
			return true
		}
	}
	return false
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
	if checkResult(f.Results()) {
		hint := "assign to an underscore to ignore"
		if _, ok := f.Results().At(0).Type().(*types.Signature); ok && f.Results().Len() == 1 {
			hint = fmt.Sprintf("did you mean to call the returned closure? %s()", formatNode(pass, call))
		}
		pass.Reportf(call.Pos(), "call discards return value (hint: %s)", hint)
	}
	return true
}

func runDiscardAnalyzer(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		switch node := node.(type) {
		case *ast.DeferStmt:
			return discardCheckCall(pass, node.Call)
		case *ast.ExprStmt:
			call, ok := node.X.(*ast.CallExpr)
			if !ok {
				return true
			}
			return discardCheckCall(pass, call)
		}
		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
