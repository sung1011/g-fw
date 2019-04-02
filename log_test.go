package fw

import (
	"sung/g-fw/logger"
	"testing"
)

func TestLog(t *testing.T) {
	log := logger.New()
	log.Warn("haha")
}
