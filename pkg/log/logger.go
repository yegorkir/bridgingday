package log

import (
	"github.com/google/uuid"
	"github.com/mymmrac/telego"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func (l *Logger) Info(message string) {
	l.logger.Info(message)
}

func (l *Logger) Error(err error, msg string) string {
	errId := uuid.New().String()

	l.logger.Error(msg, zap.Error(err), zap.String("err_id", errId))
	return errId
}

func (l *Logger) ErrorWithMessage(message telego.Message, err error, msg string) string {
	errId := uuid.New().String()
	l.logger.Error(msg, zap.Error(err),
		zap.Int64("chat_id", message.Chat.ID),
		zap.String("username", message.Chat.Username),
		zap.String("first_name", message.From.FirstName),
		zap.String("text", message.Text),
		zap.String("err_id", errId))
	return errId
}

func (l *Logger) Fatal(err error, msg string) string {
	errId := uuid.New().String()
	l.logger.Fatal(msg, zap.Error(err), zap.String("err_id", errId))

	return errId
}

func (l *Logger) Debug(message string) {
	l.logger.Debug(message)
}

func (l *Logger) DebugWithMessage(message telego.Message, msg string) {
	l.logger.Debug(msg,
		zap.Int64("chat_id", message.Chat.ID),
		zap.String("username", message.Chat.Username),
		zap.String("first_name", message.From.FirstName),
		zap.String("text", message.Text),
	)
}

func (l *Logger) Warn(message string) {
	l.logger.Warn(message)
}

func (l *Logger) WarnWithMessage(message telego.Message, msg string) {
	l.logger.Warn(msg,
		zap.Int64("chat_id", message.Chat.ID),
		zap.String("username", message.Chat.Username),
		zap.String("first_name", message.From.FirstName),
		zap.String("text", message.Text),
	)
}
