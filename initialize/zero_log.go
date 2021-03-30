package initialize

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"goblog/config"
	"goblog/consts"
	"io"
	"os"
)

func ZeroLogInit() {
	var level zerolog.Level
	switch config.ServerConfig.ZeroLog.Level { // 初始化配置文件的Level
	case "debug":
		level = zerolog.DebugLevel
	case "info":
		level = zerolog.InfoLevel
	case "warn":
		level = zerolog.WarnLevel
	case "error":
		level = zerolog.ErrorLevel
	case "fatal":
		level = zerolog.FatalLevel
	case "panic":
		level = zerolog.PanicLevel
	case "no_level":
		level = zerolog.NoLevel
	case "disabled":
		level = zerolog.Disabled
	default:
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)
	zerolog.CallerSkipFrameCount = config.ServerConfig.ZeroLog.CallerSkipFrameCount

	var logFilePath string
	if config.ServerConfig.ZeroLog.LogFilePath != "" {
		logFilePath = config.ServerConfig.ZeroLog.LogFilePath
	} else {
		logFilePath = consts.LogFilePath
	}
	logfile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR|os.O_SYNC, 0666)
	if err != nil {
		log.Panic().Err(err).Msg("打开日志文件错误")
	}

	var w io.Writer
	if config.ServerConfig.ZeroLog.LogInConsole {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		w = zerolog.MultiLevelWriter(consoleWriter, logfile)
	} else {
		w = logfile
	}

	log.Logger = zerolog.New(w).With().Timestamp().Logger()
}
