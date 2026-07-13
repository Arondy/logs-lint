package testdata

import (
	"log/slog"

	"go.uber.org/zap"
)

func incorrectLogs() {
	password := ""

	slog.Error("Starting server") // want "message should start with lowercase letter"

	slog.Info("запуск сервера")                    // want "message should start with lowercase letter" "message contains prohibited characters"
	slog.Info("server started! 🚀")                 // want "message contains prohibited characters"
	slog.Error("connection failed!!!")             // want "message contains prohibited characters"
	slog.Warn("warning: something went wrong...")  // want "message contains prohibited characters"
	slog.Info("user password: ")                   // want "message contains prohibited characters"
	slog.Debug("api_key=")                         // want "message contains prohibited characters"
	slog.Info("password " + password + " updated") // want "message contains sensitive variable"

	zap.L().Debug("Smth") // want "message should start with lowercase letter"

	zap.NewExample().Debug(": smth") // want "message should start with lowercase letter" "message contains prohibited characters"

	zap.NewExample().Debug(password + "smth") // want "message contains sensitive variable"
}
