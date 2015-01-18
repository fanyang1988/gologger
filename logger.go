package log

const (
	All = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

type Logger struct {
	level    int
	handlers []logHandler
	name     string
}

func (self *Logger) log(level int, info string) {
	msg := newLogMsg(level, info)
	self.logMsg(msg)
}

func (self *Logger) logMsg(msg *logMsg) error {
	for _, hander := range self.handlers {
		hander.handle(msg)
	}

	return nil
}

func (self *Logger) Debug(info string) {
	self.log(DEBUG, info)
}

func (self *Logger) Info(info string) {
	self.log(INFO, info)
}

func (self *Logger) Warn(info string) {
	self.log(WARN, info)
}

func (self *Logger) Error(info string) {
	self.log(ERROR, info)
}

func (self *Logger) Fatal(info string) {
	self.log(FATAL, info)
}
