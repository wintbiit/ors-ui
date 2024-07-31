package internal

import (
	"context"
	"gorm.io/gorm/logger"
	"os"
	"time"

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
	ctxLogger := log.With().Str("client", name).Logger()

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

type gormLogger struct {
	*zerolog.Logger

	LogLevel logger.LogLevel
}

func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l *gormLogger) Info(ctx context.Context, format string, args ...interface{}) {
	l.Logger.Info().Ctx(ctx).Msgf(format, args...)
}

func (l *gormLogger) Warn(ctx context.Context, format string, args ...interface{}) {
	l.Logger.Warn().Ctx(ctx).Msgf(format, args...)
}

func (l *gormLogger) Error(ctx context.Context, format string, args ...interface{}) {
	l.Logger.Error().Ctx(ctx).Msgf(format, args...)
}

func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rowsAffected := fc()

	if err != nil {
		l.Logger.Error().Ctx(ctx).Err(err).Dur("elapsed", elapsed).Int64("rows affected", rowsAffected).Msg(sql)
	} else {
		l.Logger.Info().Ctx(ctx).Dur("elapsed", elapsed).Int64("rows affected", rowsAffected).Msg(sql)
	}
}
