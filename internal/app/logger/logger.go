package logger

import (
	"uszn-go-test/internal/app/config"

	"github.com/sirupsen/logrus"
)

// Logger ...
type Logger struct {
	Logger *logrus.Logger
	config *config.Config
}

// New ...
func New(config *config.Config) *Logger {
	return &Logger{
		Logger: logrus.New(),
		config: config,
	}
}

// InitLogger ...
func (l *Logger) InitLogger() error {
	level, err := logrus.ParseLevel(l.config.LogLevel)
	if err != nil {
		return err
	}

	l.Logger.SetLevel(level)

	return nil
}
