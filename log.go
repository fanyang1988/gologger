package log

type Log struct {
	loggers map[string]*Logger
	level   int
}

func (self *Log) GetLogger(name string) *Logger {
	if self.loggers[name] == nil {
		new_logger := self.newLogger(name)
		self.loggers[name] = new_logger
		return new_logger
	} else {
		return self.loggers[name]
	}
}

func (self *Log) newLogger(name string) *Logger {
	return &Logger{
		level: self.level,
		name:  name,
	}
}
