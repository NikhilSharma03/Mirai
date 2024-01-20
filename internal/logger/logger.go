package logger

import (
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
)

type Logger interface {
	Info(string)
	Debug(string, string)
	Error(error)
	Panic(error)
}

type logger struct {
	log zerolog.Logger
}

func (l logger) Info(msg string) {
	l.log.Info().Msg(msg)
}

func (l logger) Debug(key, value string) {
	l.log.Debug().Str(key, value).Send()
}

func (l logger) Error(err error) {
	l.log.Error().Msg(err.Error())
}

func (l logger) Panic(err error) {
	l.log.Panic().Msg(err.Error())
}

func NewLogger(log zerolog.Logger) Logger {
	return &logger{log}
}

func GetConfiguredLogger() zerolog.Logger {
	buildInfo, _ := debug.ReadBuildInfo()

	return zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Int("pid", os.Getpid()).
		Str("go_version", buildInfo.GoVersion).
		Logger()
}
