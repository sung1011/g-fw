package logger

import (
	"bytes"
	"fmt"
	"os"
)

type FormatterText struct {
	TimestampFormat string
	SortingFunc     func([]string)
}

func (f *FormatterText) Render(l *Log) ([]byte, error) {
	var buff bytes.Buffer
	//Level
	lb, err := l.Level.MarshalText()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	buff.Write(lb)
	buff.WriteString(" ")
	//Msg
	if len(l.Msg) > 0 {
		buff.WriteString(l.Msg)
	}
	//Kv
	for k, v := range l.Kv {
		if buff.Len() > 0 {
			buff.WriteString(" ")
		}
		buff.WriteString(k)
		buff.WriteString("=")
		s, ok := v.(string)
		if !ok {
			s = fmt.Sprint(v)
		}
		buff.WriteString(s)
	}
	buff.WriteString("\n")
	return buff.Bytes(), nil
}

// func (f *FormatterText) Render(l *Log) ([]byte, error) {
// 	v := reflect.ValueOf(*l)
// 	count := v.NumField()
// 	var bs []byte
// 	for i := 0; i < count; i++ {
// 		s := fmt.Sprintf("%v", v.Field(i))
// 		for _, b := range []byte(s) {
// 			bs = append(bs, b)
// 		}
// 	}
// 	return bs, nil
// }
