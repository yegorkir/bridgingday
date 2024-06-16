package log

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/yegorkir/review/internal/pkg/env"
)

const (
	programENV     env.Key = "ENV"
	programENVDev  env.Env = "dev"
	programENVProd env.Env = "prod"

	envLogFilePath env.Key = "LOGFILE_PATH"
)

func NewLogger(hooks ...func(zapcore.Entry) error) *Logger {
	environment, err := programENV.GetEnv()
	if err != nil {
		fmt.Printf("logger env: %s", err)
		os.Exit(1)
	}

	var logger *zap.Logger

	switch environment {
	case programENVDev:
		logger, err = newDevelopmentLogger(hooks...)
	case programENVProd:
		logger, err = newProductionLogger(hooks...)
	default:
		fmt.Printf("unknown environment: %s", environment)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("failed to initialize logger: %v", err)
		os.Exit(1)
	}

	return &Logger{
		logger: logger,
	}
}

func newDevelopmentLogger(hooks ...func(zapcore.Entry) error) (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = customTimeEncoder

	logger, err := config.Build(zap.Hooks(hooks...))
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func newProductionLogger(hooks ...func(zapcore.Entry) error) (*zap.Logger, error) {
	logFilePath, err := envLogFilePath.GetEnv()
	if err != nil {
		fmt.Printf("logger file path env: %s", err)
		os.Exit(1)
	}

	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout", fmt.Sprintf("%s.%s.log", logFilePath, time.Now().Format("2006.01.02_15:04:05"))}
	config.ErrorOutputPaths = []string{"stderr"}
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	config.EncoderConfig.EncodeTime = customTimeEncoder

	logger, err := config.Build(zap.Hooks(hooks...))
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006.01.02 15:04:05.000"))
}
