package logger

import (
	"fmt"
	"reflect"
)

type FormatterText struct {
}

func (f *FormatterText) Render(l *Log) ([]byte, error) {
	v := reflect.ValueOf(*l)
	count := v.NumField()
	var bs []byte
	for i := 0; i < count; i++ {
		s := fmt.Sprintf("%v", v.Field(i))
		for _, b := range []byte(s) {
			bs = append(bs, b)
		}
	}
	return bs, nil
}
