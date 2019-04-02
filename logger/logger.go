package logger

import (
	"io"
)

var (
	green        = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white        = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow       = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red          = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue         = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta      = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan         = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset        = string([]byte{27, 91, 48, 109})
	disableColor = false
)

type Logger struct {
	Out       io.Writer
	Hooks     todo
	Formatter Formatter
	ExitFunc  exitFunc
}

type exitFunc func(int)

type todo interface{}

func NewLogger() *Logger {
	return &Logger{
		Out:       DefaultLoggerOut,
		Hooks:     nil,
		Formatter: new(FormatterJson),
		ExitFunc:  DefaultLoggerExitFunc,
	}
}

func (this *Logger) Exit(code int) {
	if this.ExitFunc == nil {
		this.ExitFunc = DefaultLoggerExitFunc
	}
	this.ExitFunc(code)
}
