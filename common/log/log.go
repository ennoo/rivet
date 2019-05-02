package log

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func InitLoggerWithDebug() {
	InitLoggerWithLevel("debug")
}

func InitLogger() {
	InitLoggerWithLevel("info")
}

func InitLoggerWithLevel(level string) {
	var logger *zap.Logger
	var err error

	if level == "debug" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	logger.Debug("Logger start success")
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	zap.S().Debug(args)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	zap.S().Info(args)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	zap.S().Warn(args)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	zap.S().Error(args)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanic(args ...interface{}) {
	zap.S().DPanic(args)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	zap.S().Panic(args)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	zap.S().Fatal(args)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	zap.S().Debugf(template, args)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	zap.S().Infof(template, args)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	zap.S().Warnf(template, args)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	zap.S().Errorf(template, args)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicf(template string, args ...interface{}) {
	zap.S().DPanicf(template, args)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	zap.S().Panicf(template, args)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	zap.S().Fatalf(template, args)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//  s.With(keysAndValues).Debug(msg)
func Debugw(msg string, keysAndValues ...interface{}) {
	zap.S().Debugw(msg, keysAndValues)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	zap.S().Infow(msg, keysAndValues)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
	zap.S().Warnw(msg, keysAndValues)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	zap.S().Errorw(msg, keysAndValues)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func DPanicw(msg string, keysAndValues ...interface{}) {
	zap.S().DPanicw(msg, keysAndValues)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...interface{}) {
	zap.S().Panicw(msg, keysAndValues)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
	zap.S().Fatalw(msg, keysAndValues)
}
