package analyzer

import (
	"slices"
)

var logPackages = []string{"log/slog", "go.uber.org/zap"}
var logFunctionNames = []string{"Log", "Debug", "Info", "Warn", "Error", "Fatal"}
var sensitiveData = []string{"password", "key", "token"}

func isLogPackage(packagePath string) bool {
	return slices.Contains(logPackages, packagePath)
}

func isLogFunction(function string) bool {
	return slices.Contains(logFunctionNames, function)
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

func isSensitiveVar(name string) bool {
	return slices.Contains(sensitiveData, name)
}
