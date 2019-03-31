package fw

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv(ModeEnv, string(TestMode))
}

func TestSetMode(t *testing.T) {
	assert.Equal(t, TestMode, Mode()) // unexpect Why?
	os.Unsetenv(ModeEnv)

	SetMode(DebugMode)
	assert.Equal(t, DebugMode, Mode())

	SetMode(TestMode)
	assert.Equal(t, TestMode, Mode())

	SetMode(ReleaseMode)
	assert.Equal(t, ReleaseMode, Mode())

	assert.Panics(t, func() { SetMode("unkown") })
}
