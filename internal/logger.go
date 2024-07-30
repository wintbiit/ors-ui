package internal

import (
	"os"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/rs/zerolog/log"
)

func bootLogger() {
	lumberjackLogger := lumberjack.Logger{
		Filename:   "logs/ors.log",
		MaxSize:    10, // megabytes
		MaxBackups: 10,
		MaxAge:     28, // days
		Compress:   true,
	}

	log.Logger = log.Output(zerolog.MultiLevelWriter(
		zerolog.ConsoleWriter{Out: os.Stderr},
		zerolog.SyncWriter(&lumberjackLogger),
	))

	if DEBUG {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	log.Info().Msg("Logger initialized")
}

type serverLogger struct {
	*zerolog.Logger
}

func newServerLogger(name string) *serverLogger {
	ctxLogger := log.With().Str("server", name).Logger()

	return &serverLogger{
		Logger: &ctxLogger,
	}
}

func (l *serverLogger) Infof(format string, args ...interface{}) {
	l.Info().Msgf(format, args...)
}

func (l *serverLogger) Errorf(format string, args ...interface{}) {
	l.Error().Msgf(format, args...)
}

func (l *serverLogger) Warnf(format string, args ...interface{}) {
	l.Warn().Msgf(format, args...)
}

func (l *serverLogger) Debugf(format string, args ...interface{}) {
	l.Debug().Msgf(format, args...)
}
