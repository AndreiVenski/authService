package logger

import (
	"authService/config"
	"go.uber.org/zap"
)

type Logger interface {
	InitLogger()
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

type ApiLogger struct {
	suggarLogger *zap.SugaredLogger
	cfg          *config.Config
}

func NewApiLogger(cfg *config.Config) *ApiLogger {
	return &ApiLogger{
		cfg: cfg,
	}
}

func (l *ApiLogger) InitLogger() {
	configL := zap.NewProductionConfig()
	configL.OutputPaths = []string{"app.log", "stdout"}

	logger, err := configL.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	l.suggarLogger = logger.Sugar()
}

func (l *ApiLogger) Info(args ...interface{}) {
	l.suggarLogger.Info(args...)
}

func (l *ApiLogger) Infof(template string, args ...interface{}) {
	l.suggarLogger.Infof(template, args...)
}

func (l *ApiLogger) Error(args ...interface{}) {
	l.suggarLogger.Error(args...)
}

func (l *ApiLogger) Errorf(template string, args ...interface{}) {
	l.suggarLogger.Errorf(template, args...)
}

func (l *ApiLogger) Fatal(args ...interface{}) {
	l.suggarLogger.Fatal(args...)
}

func (l *ApiLogger) Fatalf(template string, args ...interface{}) {
	l.suggarLogger.Fatalf(template, args...)
}
