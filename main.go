package main

import (
	"time"

	"go.uber.org/zap"
	// "cloud.google.com/go/logging"
	// "golang.org/x/net/context"
)

func main() {
	// logger, err := zap.NewProduction(zap.WrapCore(func(c zapcore.Core) zapcore.Core {
	// 	encConfig := zapcore.EncoderConfig{
	// 		LevelKey: "severity",
	// 	}
	// 	c2 := zapcore.NewCore(zapcore.NewJSONEncoder(encConfig), nil, nil)
	// 	return c2
	// }))
	// logger, err := zap.NewProduction(zap.Fields(zap.String("levelKey", "severity")))
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// rawJSON := []byte(`{
	// 	"level": "debug",
	// 	"encoding": "json",
	// 	"outputPaths": ["stdout"],
	// 	"errorOutputPaths": ["stderr"],
	// 	"initialFields": {"foo": "bar"},
	// 	"encoderConfig": {
	// 	  "messageKey": "message",
	// 	  "levelKey": "severity",
	// 	  "levelEncoder": "lowercase"
	// 	}
	//   }`)

	// var cfg zap.Config
	// if err := json.Unmarshal(rawJSON, &cfg); err != nil {
	// 	panic(err)
	// }
	// logger2, err := cfg.Build()
	// if err != nil {
	// 	panic(err)
	// }
	// defer logger2.Sync()

	// logger2.Info("logger construction succeeded")

	cnt := 0
	for {
		if cnt < 3 {
			// fmt.Println("Normal Text Log", cnt)
			// fmt.Println(`{"level": "info","message": "Normal JSON Log"}`)
			// e := errors.New("Dummy Err")
			// strace := runtime.StartTrace()
			// fmt.Printf("{\"serviceContext\": {\"service\": \"tempus\", \"version\": \"v1234567890\"}, \"message\": \"%s\"}\n", e.Error())
			// fmt.Printf("{\"logName\": \"tempus\", \"resource\": {}}")

			// logger.Debug("DEBUG LEVEL", zap.Int("cnt", cnt))
			// logger.Info("INFO LEVEL", zap.Int("cnt", cnt))
			// logger.Warn("WARN LEVEL", zap.Int("cnt", cnt))
			// logger.Error("ERROR LEVEL", zap.Int("cnt", cnt))

			logger.Debug("DEBUG LEVEL with severity", zap.Int("cnt", cnt), zap.String("severity", "DEBUG"))
			logger.Info("INFO LEVEL with severity", zap.Int("cnt", cnt), zap.String("severity", "INFO"))
			logger.Warn("WARN LEVEL with severity", zap.Int("cnt", cnt), zap.String("severity", "WARN"))
			logger.Warn("WARNING LEVEL with severity", zap.Int("cnt", cnt), zap.String("severity", "WARNING"))
			logger.Error("ERROR LEVEL with severity", zap.Int("cnt", cnt), zap.String("severity", "ERROR"))

			// logger2.Info("Logger2 INFO LEVEL", zap.Int("cnt", cnt))
			// logger2.Warn("Logger2 WARN LEVEL", zap.Int("cnt", cnt))
			// logger2.Error("Logger2 ERROR LEVEL", zap.Int("cnt", cnt))
			cnt = cnt + 1
		}
		time.Sleep(3 * time.Second)
	}
}
