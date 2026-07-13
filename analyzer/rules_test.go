package analyzer

import (
	"testing"
)

func TestLogPackage(t *testing.T) {
	tests := []struct {
		name        string
		packageName string
		expected    bool
	}{
		{"slog", "log/slog", true},
		{"zap", "go.uber.org/zap", true},
		{"http", "net/http", false},
		{"math", "math", false},
		{"zap", "github.com/uber/zap", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := isLogPackage(test.packageName)
			if res != test.expected {
				t.Fatalf("incorrect result for %s, expected %t, got %t", test.packageName, test.expected, res)
			}
		})
	}
}

func TestLogFunction(t *testing.T) {
	tests := []struct {
		function string
		expected bool
	}{
		{"Log", true},
		{"Debug", true},
		{"Info", true},
		{"Warn", true},
		{"Error", true},
		{"Fatal", true},
		{"SomeFunc", false},
		{"Check", false},
		{"log", false},
		{"DebugSmth", false},
	}

	for _, test := range tests {
		t.Run(test.function, func(t *testing.T) {
			res := isLogFunction(test.function)
			if res != test.expected {
				t.Fatalf("incorrect result for %s, expected %t, got %t", test.function, test.expected, res)
			}
		})

	}
}

func TestStartsWithLowercase(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		expected bool
	}{
		{"lowercase", "correct message", true},
		{"lowercase + mixed", "cORrEct messAge", true},
		{"uppercase", "Incorrect message", false},
		{"digit", "4 incorrect message", false},
		{"russian", "неверное сообщение", false},
		{"special symbols", ":incorrect message", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := startsWithLowercase(test.message)
			if res != test.expected {
				t.Fatalf("incorrect result for %s, expected %t, got %t", test.message, test.expected, res)
			}
		})
	}
}

func TestAreCharactersAllowed(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		expected bool
	}{
		{"lowercase english", "correctmessage", true},
		{"mixed case english", "CorrectMessage", true},
		{"english + space", " correct message", true},
		{"english + digits", "42correct42", true},
		{"digits + space", "42 420", true},
		{"english + digits + space", "correct 42 420 message", true},
		{"russian", "неверное сообщение", false},
		{"emoji", "❌ message", false},
		{"special symbols '.'", ".", false},
		{"special symbols '!'", "!", false},
		{"special symbols '?'", "?", false},
		{"special symbols ':'", ":", false},
		{"special symbols '='", "=", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := areCharactersAllowed(test.message)
			if res != test.expected {
				t.Fatalf("incorrect result for %s, expected %t, got %t", test.message, test.expected, res)
			}
		})
	}
}

func TestIsSensitiveVar(t *testing.T) {
	tests := []struct {
		varName  string
		expected bool
	}{
		{"variable", false},
		{"Variable", false},
		{"randomName", false},
		{"password", true},
		{"key", true},
		{"token", true},
	}

	for _, test := range tests {
		t.Run(test.varName, func(t *testing.T) {
			res := isSensitiveVar(test.varName)
			if res != test.expected {
				t.Fatalf("incorrect result for %s, expected %t, got %t", test.varName, test.expected, res)
			}
		})
	}
}
