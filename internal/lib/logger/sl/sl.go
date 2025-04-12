package sl

import (
	"decadecollab/internal/lib/logger/prettylog"
	"log/slog"
	"os"
	//"errors"
)

//функции *Opts принимают не только строку/ошибку, но и атрибуты

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func New(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func Debug(logger *slog.Logger, msg string, args ...any) {
	logger.Debug(msg, args...)
}

func DebugOpts(msg string, opts *slog.HandlerOptions) {
	logger := slog.New(prettylog.NewHandler(opts))

	logger.Debug(msg)
}

func Info(logger *slog.Logger, msg string, args ...any) {
	logger.Info(msg, args...)
}

func InfoOpts(msg string, opts *slog.HandlerOptions) {
	logger := slog.New(prettylog.NewHandler(opts))

	logger.Info(msg)
}

func Warn(logger *slog.Logger, msg string, args ...any) {
	logger.Warn(msg, args...)
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func ErrOpts(err error, opts *slog.HandlerOptions) {
	logger := slog.New(prettylog.NewHandler(opts))

	logger.Error("We have some problems", "error", err)
}

func Error(logger *slog.Logger, msg string, err error, args ...any) {
	logger.Error(msg, append(args, Err(err))...)
}

func Fatal(logger *slog.Logger, msg string, err error, args ...any) {
	logger.Error(msg, append(args, Err(err))...)
	os.Exit(1)
}
