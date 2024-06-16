package tglogger

import (
	"fmt"
	"os"

	"go.uber.org/zap/zapcore"

	"github.com/yegorkir/review/internal/pkg/env"
	"github.com/yegorkir/review/internal/services/botusers"
	"github.com/yegorkir/review/internal/services/tbot"
)

const (
	envLoggerBotKey env.Key = "LOGGER_BOT_TOKEN"
)

const formatErr = "ERROR: %s\n" +
	"time:        %s\n" +
	"loggerName:  %s\n" +
	"caller:      %s\n" +
	"stack:\n%s"

func NewHook() func(entry zapcore.Entry) error {
	loggerBotKey, err := envLoggerBotKey.GetEnv()
	if err != nil {
		fmt.Printf("tg logger hook: %s", err)
		os.Exit(1)
	}

	bot, err := tbot.NewBot(
		tbot.WithBotToken(loggerBotKey.String()),
	)
	if err != nil {
		fmt.Printf("can't get bot: %s", err)
		os.Exit(1)
	}

	admin, err := botusers.GetAdminWithBot(bot)
	if err != nil {
		fmt.Printf("can't get admin: %s", err)
		os.Exit(1)
	}

	return func(entry zapcore.Entry) error {
		switch entry.Level {
		case zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel, zapcore.ErrorLevel, zapcore.WarnLevel:
			admin.Log(formatErr,
				entry.Message,
				entry.Time.Format("2006.01.02 15:04:05.000"),
				entry.LoggerName,
				entry.Caller,
				entry.Stack)
		}
		return nil
	}
}
