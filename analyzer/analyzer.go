package analyzer

import (
	"go/ast"
	"slices"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var checkedPackages = []string{"log", "slog", "zap"}
var logFunctionNames = []string{"Log", "Debug", "Info", "Warn", "Error", "Fatal"}

var Analyzer = analysis.Analyzer{
	Name:     "logs_lint",
	Doc:      "finds logs with issues",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(node ast.Node) {
		expr := node.(*ast.CallExpr)

		if sel, ok := expr.Fun.(*ast.SelectorExpr); ok {
			if pkg, ok := sel.X.(*ast.Ident); ok {
				packageName := pkg.Name
				if !isCheckedPackage(packageName) {
					return
				}

				method := sel.Sel.Name
				if !isLogFunction(method) {
					return
				}

				pass.Report(analysis.Diagnostic{
					Pos:      sel.Pos(),
					Category: "logs",
					Message:  "logs found",
				})
			}
		}
	})

	return nil, nil
}

func isCheckedPackage(packageName string) bool {
	return slices.Contains(checkedPackages, packageName)
}

func isLogFunction(method string) bool {
	return slices.Contains(logFunctionNames, method)
}
