package logger

import (
	"encoding/json"
	"fmt"
)

type FormatterJson struct {
	PrettyPrint bool
	FieldMap    FieldMap
}

func (f *FormatterJson) Render(l *Log) ([]byte, error) {
	bs, err := json.Marshal(l)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}
	return bs, nil
}
