package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"
	"strconv"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = analysis.Analyzer{
	Name:     "logs_lint",
	Doc:      "find logs with style/security issues",
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
		sel, ok := expr.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}

		method := sel.Sel.Name
		if !isLogFunction(method) {
			return
		}

		pkg := unwrapPkg(sel)
		if pkg == nil {
			return
		}

		obj := pass.TypesInfo.ObjectOf(pkg)
		pkgName, ok := obj.(*types.PkgName)
		if !ok {
			return
		}

		packagePath := pkgName.Imported().Path()
		if !isLogPackage(packagePath) {
			return
		}

		args := expr.Args
		if len(args) == 0 {
			return
		}

		arg := args[0]

		if lit, ok := arg.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			checkStringLog(pass, lit)
		} else if ident, ok := arg.(*ast.Ident); ok {
			if isSensitiveVar(ident.Name) {
				pass.Report(analysis.Diagnostic{
					Pos:      ident.Pos(),
					Category: "security",
					Message:  "message contains sensitive variable",
				})
			}
		}
	})

	return nil, nil
}

func unwrapPkg(sel *ast.SelectorExpr) *ast.Ident {
	for {
		pkg, ok := sel.X.(*ast.Ident)
		if ok {
			return pkg
		}

		var call *ast.CallExpr
		call, ok = sel.X.(*ast.CallExpr)
		if !ok {
			return nil
		}

		sel, ok = call.Fun.(*ast.SelectorExpr)
		if !ok {
			return nil
		}
	}
}

func checkStringLog(pass *analysis.Pass, lit *ast.BasicLit) {
	message, err := strconv.Unquote(lit.Value)
	if err != nil {
		return
	}

	if !startsWithLowercase(message) {
		pass.Report(analysis.Diagnostic{
			Pos:      lit.Pos(),
			Category: "style",
			Message:  "message should start with lowercase letter",
		})
	}
	if !areCharactersAllowed(message) {
		pass.Report(analysis.Diagnostic{
			Pos:      lit.Pos(),
			Category: "style",
			Message:  "message contains prohibited characters",
		})
	}
	if hasSensitiveData(message) {
		pass.Report(analysis.Diagnostic{
			Pos:      lit.Pos(),
			Category: "security",
			Message:  "message contains sensitive data",
		})
	}
}
