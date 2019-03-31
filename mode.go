package fw

import (
	"os"
)

const (
	DebugMode   ModeType = "debug"
	TestMode    ModeType = "test"
	ReleaseMode ModeType = "release"
)

type ModeType string

var mode ModeType

var ModeEnv = "G_FW_MODE"

func init() {
	me := os.Getenv(ModeEnv)
	SetMode(ModeType(me))
}

func SetMode(m ModeType) {
	mode = chkMode(m)
}

func Mode() ModeType {
	return mode
}

func chkMode(m ModeType) ModeType {
	switch m {
	case DebugMode, TestMode, ReleaseMode:
		return m
	case "":
		return DebugMode
	}
	panic("mode unkown: " + m)
}
