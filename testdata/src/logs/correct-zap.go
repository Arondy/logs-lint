package testdata

import "go.uber.org/zap"

func correctZapLogs() {
	zap.L().Info("starting server")
	zap.L().Warn("retry attempt 3 of 5")
}
