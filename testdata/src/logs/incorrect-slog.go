package testdata

import (
	"log/slog"
)

func incorrectSlogLogs() {
	password := ""

	slog.Error("Starting server") // want "message should start with lowercase letter"

	slog.Info("запуск сервера")                    // want "message should start with lowercase letter" "message contains prohibited characters"
	slog.Info("server started! 🚀")                 // want "message contains prohibited characters"
	slog.Error("connection failed!!!")             // want "message contains prohibited characters"
	slog.Warn("warning: something went wrong...")  // want "message contains prohibited characters"
	slog.Info("user password: ")                   // want "message contains prohibited characters"
	slog.Debug("api_key=")                         // want "message contains prohibited characters"
	slog.Info("password " + password + " updated") // want "message contains sensitive variable"
}
