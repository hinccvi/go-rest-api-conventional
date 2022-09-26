package log

import (
	"context"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
	"gopkg.in/natefinch/lumberjack.v2"
)

type (
	Type string
	Env  string
)

const (
	ErrorLog Type = "error"
	APILog   Type = "api"
	SQLLog   Type = "sql"

	LocalEnv Env = "local"

	localPath = "./"
	devPath   = "/var/log/app/"
	prodPath  = "/www/wwwlog/"

	//	Configuration for error log
	errorLogFileName  = "error.log"
	errorLogMaxSize   = 500
	errorLogMaxBackup = 7
	errorLogMaxAge    = 7

	//	Configuration for sql log
	sqlLogFileName  = "sql.log"
	sqlLogMaxSize   = 300
	sqlLogMaxBackup = 2
	sqlLogMaxAge    = 3

	//	Configuration for api log
	accessLogFileName  = "api.log"
	accessLogMaxSize   = 400
	accessLogMaxBackup = 3
	accessLogMaxAge    = 7
)

// Logger is a logger that supports log levels, context and structured logging.
type Logger interface {
	// With returns a logger based off the root logger and decorates it with the given context and arguments.
	With(ctx context.Context, args ...interface{}) Logger

	// Debug uses fmt.Sprint to construct and log a message at DEBUG level
	Debug(args ...interface{})
	// Info uses fmt.Sprint to construct and log a message at INFO level
	Info(args ...interface{})
	// Error uses fmt.Sprint to construct and log a message at WARN level
	Warn(args ...interface{})
	// Error uses fmt.Sprint to construct and log a message at ERROR level
	Error(args ...interface{})
	// Error uses fmt.Sprint to construct and log a message at WARN level
	Fatal(args ...interface{})

	// Debugf uses fmt.Sprintf to construct and log a message at DEBUG level
	Debugf(format string, args ...interface{})
	// Infof uses fmt.Sprintf to construct and log a message at INFO level
	Infof(format string, args ...interface{})
	// Errorf uses fmt.Sprintf to construct and log a message at WARN level
	Warnf(format string, args ...interface{})
	// Errorf uses fmt.Sprintf to construct and log a message at ERROR level
	Errorf(format string, args ...interface{})
	// Errorf uses fmt.Sprintf to construct and log a message at WARN level
	Fatalf(format string, args ...interface{})
}

type logger struct {
	*zap.SugaredLogger
}

// New creates a new logger.
func New(env string, log Type) *zap.Logger {
	c := new(zapcore.Core)

	if env == string(LocalEnv) {
		*c = zapcore.NewTee(zapcore.NewCore(encoder(env), zapcore.Lock(os.Stdout), zap.InfoLevel))
	} else {
		var level zapcore.Level
		var writeSyncer zapcore.WriteSyncer
		var path string

		switch env {
		case "local":
			path = localPath
		case "dev":
			path = devPath
		case "qa":
			path = devPath
		case "prod":
			path = prodPath
		}

		switch log {
		case APILog:
			level = zap.InfoLevel
			writeSyncer = newWriteSyncer(path+accessLogFileName, accessLogMaxSize, accessLogMaxBackup, accessLogMaxAge)
		case SQLLog:
			level = zap.InfoLevel
			writeSyncer = newWriteSyncer(path+sqlLogFileName, sqlLogMaxSize, sqlLogMaxBackup, sqlLogMaxAge)
		case ErrorLog:
			level = zap.ErrorLevel
			writeSyncer = newWriteSyncer(path+errorLogFileName, errorLogMaxSize, errorLogMaxBackup, errorLogMaxAge)
		}

		*c = zapcore.NewTee(zapcore.NewCore(encoder(env), writeSyncer, level))
	}

	l := zap.New(*c, zap.AddCaller(), zap.AddCallerSkip(0))

	return l
}

// Customize log encoder.
func encoder(mode string) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Local().Format("2006-01-02T15:04:05Z0700"))
	}
	if mode == "local" {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func newWriteSyncer(fileName string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// NewWithZap creates a new logger using the preconfigured zap logger.
func NewWithZap(l *zap.Logger) Logger {
	return &logger{l.Sugar()}
}

// NewForTest returns a new logger and the corresponding observed
// logs which can be used in unit tests to verify log entries.
func newForTest() (Logger, *observer.ObservedLogs) {
	core, recorded := observer.New(zapcore.InfoLevel)
	return NewWithZap(zap.New(core)), recorded
}

// With returns a logger based off the root logger and decorates it with the given context and arguments.
//
// The arguments should be specified as a sequence of name, value pairs with names being strings.
// The arguments will also be added to every log message generated by the logger.
func (l *logger) With(ctx context.Context, args ...interface{}) Logger {
	if len(args) > 0 {
		return &logger{l.SugaredLogger.With(args...)}
	}
	return l
}
