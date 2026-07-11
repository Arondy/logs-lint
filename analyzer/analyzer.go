package analyzer

import (
	"go/ast"
	"go/token"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var checkedPackages = []string{"log", "slog", "zap"}
var logFunctionNames = []string{"Log", "Debug", "Info", "Warn", "Error", "Fatal"}
var sensitiveData = []string{"password", "key", "token"}

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

		pkg, ok := sel.X.(*ast.Ident)
		if !ok {
			return
		}

		packageName := pkg.Name
		method := sel.Sel.Name

		if !(isCheckedPackage(packageName) && isLogFunction(method)) {
			return
		}

		args := expr.Args

		for _, arg := range args {
			if lit, ok := arg.(*ast.BasicLit); ok && lit.Kind == token.STRING {
				checkLogMessage(pass, lit)
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

func startsWithLowercase(message string) bool {
	return len(message) > 0 && 'a' <= message[0] && message[0] <= 'z'
}

func areCharactersAllowed(message string) bool {
	for _, c := range message {
		isEnglishLetter := 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
		isNumber := '0' <= c && c <= '9'
		isSpace := c == ' '

		if !(isEnglishLetter || isNumber || isSpace) {
			return false
		}
	}

	return true
}

func hasSensitiveData(message string) bool {
	for _, data := range sensitiveData {
		if strings.Contains(message, data) {
			return true
		}
	}
	return false
}

func checkLogMessage(pass *analysis.Pass, lit *ast.BasicLit) {
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
