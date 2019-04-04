package logger

import "os"

var DefaultLoggerOut = os.Stdout
var DefaultLoggerExitFunc = os.Exit
var DefaultLoggerTimestampFormat = "2006-01-02 15:04:05"
