package logger

import (
	"fmt"
	"io"
	"log/slog"

	"rover/pkg/config"
)

// TODO возмоно лучше сделать интерфейс, а реализацию через слог отдельным пакетом
// но думается что для логера это оверхед
type Logger = slog.Logger

const (
	LOG_LEVEL_DEBUG = "DEBUG"
	LOG_LEVEL_INFO  = "INFO"
	LOG_LEVEL_WARN  = "WARN"
	LOG_LEVEL_ERROR = "ERROR"
)

func New(cfg config.Logger, w io.Writer) (*Logger, error) {
	var slogLvl slog.Level

	switch cfg.Level {
	case LOG_LEVEL_DEBUG:
		slogLvl = slog.LevelDebug
	case LOG_LEVEL_INFO:
		slogLvl = slog.LevelInfo
	case LOG_LEVEL_WARN:
		slogLvl = slog.LevelWarn
	case LOG_LEVEL_ERROR:
		slogLvl = slog.LevelError
	default:
		return nil, fmt.Errorf("invalid log level: %s", cfg.Level)
	}

	return slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{Level: slogLvl})), nil
}
