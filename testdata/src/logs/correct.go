package testdata

import (
	"log/slog"

	"go.uber.org/zap"
)

func correctLogs() {
	slog.Info("starting server")
	slog.Info("server started")
	slog.Error("connection failed")
	slog.Info("token")
	slog.Info("user logged in")
	slog.Debug("request completed 200")
	slog.Warn("retry attempt 3 of 5")

	zap.L().Info("starting server")
	zap.L().Warn("retry attempt 3 of 5")
}
