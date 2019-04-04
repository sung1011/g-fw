package logger

import (
	"fmt"
	"os"
	"time"
)

type Log struct {
	Level  Level
	logger *Logger
	Time   time.Time
	Kv     Fields
	// Buffer *bytes.Buffer
	Msg string
	err error
}

type Fields map[string]interface{}

type Level int

const (
	PanicLevel Level = iota
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

// WithKv 自定义kv
func (l *Log) WithKv(f Fields) *Log {
	m := make(Fields, len(f))
	for k, v := range f {
		m[k] = v
	}
	l.Kv = m
	return l
}

func (l *Log) String() (string, error) {
	b, err := l.logger.Formatter.Render(l)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	return string(b), nil
}

func (l *Log) Trace(args ...interface{}) {
	l.Log(TraceLevel, args...)
}

func (l *Log) Debug(args ...interface{}) {
	l.Log(DebugLevel, args...)
}

func (l *Log) Print(args ...interface{}) {
	l.Info(args...)
}

func (l *Log) Info(args ...interface{}) {
	l.Log(InfoLevel, args...)
}

func (l *Log) Warn(args ...interface{}) {
	l.Log(WarnLevel, args...)
}

func (l *Log) Warning(args ...interface{}) {
	l.Warn(args...)
}

func (l *Log) Error(args ...interface{}) {
	l.Log(ErrorLevel, args...)
}

func (l *Log) Fatal(args ...interface{}) {
	l.Log(FatalLevel, args...)
	l.logger.Exit(1)
}

func (l *Log) Panic(args ...interface{}) {
	l.Log(PanicLevel, args...)
	panic(fmt.Sprint(args...))
}

func (l *Log) Log(lv Level, args ...interface{}) {
	l.log(lv, fmt.Sprint(args...))
}

func (l *Log) log(lv Level, msg string) {
	l.Level = lv
	l.Msg = msg
	//TODO hooks
	l.write()
}

func (l *Log) write() {
	serialized, err := l.logger.Formatter.Render(l)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	l.logger.Out.Write(serialized)
}

func (level Level) MarshalText() ([]byte, error) {
	switch level {
	case TraceLevel:
		return []byte("TRACE"), nil
	case DebugLevel:
		return []byte("DEBUG"), nil
	case InfoLevel:
		return []byte("INFO"), nil
	case WarnLevel:
		return []byte("WARNING"), nil
	case ErrorLevel:
		return []byte("ERROR"), nil
	case FatalLevel:
		return []byte("FATAL"), nil
	case PanicLevel:
		return []byte("PANIC"), nil
	}

	return nil, fmt.Errorf("not a valid logrus level %d", level)
}
