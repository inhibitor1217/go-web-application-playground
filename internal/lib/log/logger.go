package log

import (
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

type Logger struct {
	zap *fxevent.ZapLogger
}

func NewLogger(logger *zap.Logger) *Logger {
	return &Logger{zap: &fxevent.ZapLogger{Logger: logger}}
}

func (l *Logger) LogEvent(event fxevent.Event) {
	l.zap.LogEvent(event)
}
