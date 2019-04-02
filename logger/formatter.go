package logger

type Formatter interface {
	Render(*Log) ([]byte, error)
}
