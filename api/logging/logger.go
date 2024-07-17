package logging

type Logger interface {
	Error(msg string, err error)
}
