package utils

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm/logger"
	"time"
)

type GormLog struct {
	Logger zerolog.Logger
	Level  zerolog.Level
}

var Gormlogger = &GormLog{}

func (l *GormLog) LogMode(ll logger.LogLevel) logger.Interface {
	var level zerolog.Level
	switch ll {
	case logger.Silent:
		level = zerolog.Disabled
	case logger.Error:
		level = zerolog.ErrorLevel
	case logger.Warn:
		level = zerolog.WarnLevel
	case logger.Info:
		level = zerolog.InfoLevel
	}
	l.Logger = log.With().Logger().Level(level)
	return l
}

func (l *GormLog) Info(ctx context.Context, msg string, vals ...interface{}) {
	tl := l.Logger.Info()
	for i, val := range vals {
		tl.Interface(fmt.Sprintf("val%d", i), val)
	}
	tl.Msg(msg)
}

func (l *GormLog) Warn(ctx context.Context, msg string, vals ...interface{}) {
	tl := l.Logger.Warn()
	for i, val := range vals {
		tl.Interface(fmt.Sprintf("val%d", i), val)
	}
	tl.Msg(msg)
}

func (l *GormLog) Error(ctx context.Context, msg string, vals ...interface{}) {
	tl := l.Logger.Error()
	for i, val := range vals {
		tl.Interface(fmt.Sprintf("val%d", i), val)
	}
	tl.Msg(msg)
}

func (l *GormLog) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	rawSql, rowsAffected := fc()
	l.Logger.Trace().Time("begin", begin).Str("raw_sql", rawSql).Int64("rows_affected", rowsAffected).Err(err).Msg("")
}
