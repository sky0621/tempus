package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	cnt := 0
	for {
		if cnt < 10 {
			logger.Debug("DEBUG LEVEL", zap.Int("cnt", cnt))
			logger.Info("DEBUG LEVEL", zap.Int("cnt", cnt))
			logger.Warn("DEBUG LEVEL", zap.Int("cnt", cnt))
			logger.Error("DEBUG LEVEL", zap.Int("cnt", cnt))
			cnt = cnt + 1
		}
		time.Sleep(5 * time.Second)
	}
}
