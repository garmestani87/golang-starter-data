package loggers

type Logger interface {
	Debug(v ...interface{})
	Warn(v ...interface{})
	Info(v ...interface{})
	Error(v ...interface{})
}
