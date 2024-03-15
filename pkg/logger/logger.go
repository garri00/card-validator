package logger

import (
	"os"

	"github.com/rs/zerolog"

	"card-validator/src/config"
)

const (
	LOG_LEVEL_DEBUG   = "DEBUG"
	LOG_LEVEL_INFO    = "INFO"
	LOG_LEVEL_TRACE   = "TRACE"
	LOG_LEVEL_PANIC   = "PANIC"
	LOG_LEVEL_NOLEVEL = "NOLEVEL"
	LOG_LEVEL_ERROR   = "ERROR"
)

var Log = zerolog.New(os.Stdout).With().Timestamp().Logger()

func SetLogLevel(c config.Configs) {
	switch c.LogLevel {
	case LOG_LEVEL_DEBUG:
		Log.Level(zerolog.DebugLevel)
	case LOG_LEVEL_INFO:
		Log = Log.Level(zerolog.InfoLevel)
	case LOG_LEVEL_TRACE:
		Log = Log.Level(zerolog.TraceLevel)
	case LOG_LEVEL_PANIC:
		Log = Log.Level(zerolog.PanicLevel)
	case LOG_LEVEL_NOLEVEL:
		Log = Log.Level(zerolog.NoLevel)
	case LOG_LEVEL_ERROR:
		Log = Log.Level(zerolog.ErrorLevel)
	default:
		Log.Level(zerolog.InfoLevel)
	}

	Log.Info().Any("log level", Log.GetLevel().String()).Msg("Log level set")
}
