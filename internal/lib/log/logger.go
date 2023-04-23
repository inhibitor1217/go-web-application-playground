package log

import (
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

type Logger struct {
	fx  *fxevent.ZapLogger
	zap *zap.Logger
}

func NewLogger(logger *zap.Logger) *Logger {
	return &Logger{fx: &fxevent.ZapLogger{Logger: logger}, zap: logger}
}

func (l *Logger) LogEvent(event fxevent.Event) {
	l.fx.LogEvent(event)
}

func (l *Logger) Debug(msg string, details ...Field) {
	l.zap.Debug(msg, zapFields(details)...)
}

func (l *Logger) Info(msg string, details ...Field) {
	l.zap.Info(msg, zapFields(details)...)
}

func (l *Logger) Error(err error, details ...Field) {
	fields := zapFields(details)
	fields = append(fields, zapField(Error(err)))
	l.zap.Error(err.Error(), fields...)
}

func (l *Logger) Fatal(msg string, details ...Field) {
	l.zap.Fatal(msg, zapFields(details)...)
}
