package testdata

import (
	"log/slog"
)

func correctSlogLogs() {
	slog.Info("starting server")
	slog.Info("server started")
	slog.Error("connection failed")
	slog.Info("token")
	slog.Info("user logged in")
	slog.Debug("request completed 200")
	slog.Warn("retry attempt 3 of 5")
}
