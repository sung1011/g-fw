package logger

import (
	"encoding/json"
	"fmt"
)

type FormatterJson struct {
	PrettyPrint bool
}

func (f *FormatterJson) Render(l *Log) ([]byte, error) {
	b := l.Buffer
	encoder := json.NewEncoder(b)
	if f.PrettyPrint {
		encoder.SetIndent("", "  ")
	}
	if err := encoder.Encode(l); err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}

	return b.Bytes(), nil
}
