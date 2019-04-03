package logger

type Formatter interface {
	Render(*Log) ([]byte, error)
}

type FieldMap map[string]string
