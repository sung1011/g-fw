package logger

type FormatterText struct {
}

func (f *FormatterText) Render(l *Log) ([]byte, error) {
	return []byte(""), nil
}
