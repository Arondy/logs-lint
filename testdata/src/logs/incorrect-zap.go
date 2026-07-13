package testdata

import (
	"go.uber.org/zap"
)

func incorrectZapLogs() {
	password := ""
	zap.L().Debug("Smth") // want "message should start with lowercase letter"

	zap.NewExample().Debug(": smth") // want "message should start with lowercase letter" "message contains prohibited characters"

	zap.NewExample().Debug(password + "smth") // want "message contains sensitive variable"
}
