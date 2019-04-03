package logger

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

type Log struct {
	Level  level
	logger *Logger
	Time   time.Time
	Data   Fields
	Buffer *bytes.Buffer
	Msg    string
	err    error
}

type Fields map[string]interface{}

type level int

const (
	PanicLevel level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

func New() *Log {
	return &Log{
		logger: NewLogger(),
		Time:   time.Now(),
	}
}

func (this *Log) String() (string, error) {
	b, err := this.logger.Formatter.Render(this)
	if err != nil {

	}
	return string(b), nil
}

func (this *Log) Trace(args ...interface{}) {
	this.Log(TraceLevel, args...)
}

func (this *Log) Debug(args ...interface{}) {
	this.Log(DebugLevel, args...)
}

func (this *Log) Print(args ...interface{}) {
	this.Info(args...)
}

func (this *Log) Info(args ...interface{}) {
	this.Log(InfoLevel, args...)
}

func (this *Log) Warn(args ...interface{}) {
	this.Log(WarnLevel, args...)
}

func (this *Log) Warning(args ...interface{}) {
	this.Warn(args...)
}

func (this *Log) Error(args ...interface{}) {
	this.Log(ErrorLevel, args...)
}

func (this *Log) Fatal(args ...interface{}) {
	this.Log(FatalLevel, args...)
	this.logger.Exit(1)
}

func (this *Log) Panic(args ...interface{}) {
	this.Log(PanicLevel, args...)
	panic(fmt.Sprint(args...))
}

func (this *Log) Log(lv level, args ...interface{}) {
	this.log(lv, fmt.Sprint(args...))
}

func (this *Log) log(lv level, msg string) {
	this.Level = lv
	this.Msg = msg
	//TODO hooks
	this.write()
}

func (this *Log) write() {
	serialized, err := this.logger.Formatter.Render(this)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	this.logger.Out.Write(serialized)
}
