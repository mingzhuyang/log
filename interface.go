package log

type XLogInterface interface {
	NewLogFromConfiguration(c *LogConfiguration)error
	Init(config map[string]string) error
	ReOpen() error
	SetLevel(level string)
	SetSkip(skip int)


	Warn(format string, a ...interface{}) error
	Fatal(format string, a ...interface{}) error
	Notice(format string, a ...interface{}) error
	Trace(format string, a ...interface{}) error
	Debug(format string, a ...interface{}) error

	Warnx(logId, format string, a ...interface{}) error
	Fatalx(logId, format string, a ...interface{}) error
	Noticex(logId, format string, a ...interface{}) error
	Tracex(logId, format string, a ...interface{}) error
	Debugx(logId, format string, a ...interface{}) error

	Close()
}