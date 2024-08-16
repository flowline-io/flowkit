package flog

import (
	"fmt"
	"github.com/flowline-io/flowkit/internal/pkg/constant"
	"github.com/flowline-io/flowkit/internal/pkg/preferences"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"io"
	"log"
	"os"
	"time"
)

var l zerolog.Logger

func Init() {
	// error stack
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	// json marshaling
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	zerolog.InterfaceMarshalFunc = json.Marshal

	var writer []io.Writer
	// log file
	logFileName := fmt.Sprintf("%s.log", constant.AppId)
	logPath := preferences.AppConfig().LogPath
	if _, err := os.Stat(logPath); !os.IsNotExist(err) {
		logFilePath := fmt.Sprintf("%s/%s", logPath, logFileName)
		var logFile *os.File
		_, err = os.Stat(logFilePath)
		if os.IsNotExist(err) {
			logFile, err = os.Create(logFilePath)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			logFile, err = os.OpenFile(logFilePath, os.O_RDWR|os.O_APPEND, 0666)
		}
		console := zerolog.ConsoleWriter{Out: logFile, NoColor: true, TimeFormat: zerolog.TimeFieldFormat}
		writer = append(writer, console)
	}
	// console
	console := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.DateTime,
		NoColor:    true,
		FormatLevel: func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		},
	}
	writer = append(writer, console)

	multi := zerolog.MultiLevelWriter(writer...)
	l = zerolog.New(multi).With().Timestamp().Logger()
}

func GetLogger() zerolog.Logger {
	return l
}

const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
	FatalLevel = "fatal"
	PanicLevel = "panic"
)

// SetLevel sets the global logging level based on the provided level.
//
// level: The logging level to set. Valid values are "debug", "info", "warn",
//
//	"error", "fatal", "panic". If an invalid level is provided, the
//	default level is set to "info".
func SetLevel(level string) {
	switch level {
	case DebugLevel:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case InfoLevel:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case WarnLevel:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case ErrorLevel:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case FatalLevel:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case PanicLevel:
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

func Debug(format string, a ...any) {
	l.Debug().Caller(1).Msgf(format, a...)
}

func Info(format string, a ...any) {
	l.Info().Caller(1).Msgf(format, a...)
}

func Warn(format string, a ...any) {
	l.Warn().Caller(1).Msgf(format, a...)
}

func Error(err error) {
	l.Error().Caller(1).Err(err).Stack().Msg(err.Error())
}

func Fatal(format string, a ...any) {
	l.Fatal().Caller(1).Stack().Msgf(format, a...)
}

func Panic(format string, a ...any) {
	l.Panic().Caller(1).Stack().Msgf(format, a...)
}
