package fw

import (
	"sung/g-fw/logger"
	"testing"
)

func TestLog(t *testing.T) {
	logger.New().WithKv(logger.Fields{
		"foo": 123,
		"bar": "xxx",
	}).Warn("hahaha")
}
