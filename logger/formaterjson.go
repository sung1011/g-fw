package logger

import (
	"encoding/json"
	"fmt"
)

type FormatterJson struct {
	PrettyPrint bool
}

func (f *FormatterJson) Render(l *Log) ([]byte, error) {
	b, err := json.Marshal(l)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}
	return b, nil
}
