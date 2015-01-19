package log

import (
	"container/list"
	"sync"
)

const (
	All = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

type Logger struct {
	handlers *list.List
	name     string
	log_conn chan *logMsg
	log_wait sync.WaitGroup
}

func NewLogger(name string) *Logger {
	new_logger := &Logger{
		name:     name,
		log_conn: make(chan *logMsg),
	}
	new_logger.reload()
	go new_logger.startLog()
	return new_logger
}

func CloseLogger(logger *Logger) {
	if logger != nil {
		logger.endLog()
	}
}

func (self *Logger) log(level int, info string) {
	msg := newLogMsg(level, info)
	if msg != nil {
		self.log_wait.Add(1)
		self.log_conn <- msg
	}
}

func (self *Logger) startLog() {
	for {
		select {
		case msg, ok := <-self.log_conn:
			if !ok {
				return
			}
			if msg == nil {
				self.reload()
			} else {
				self.logMsg(msg)
				self.log_wait.Done()
			}
		}
	}
}

func (self *Logger) endLog() {
	self.log_wait.Wait()
	close(self.log_conn)
}

func (self *Logger) logMsg(msg *logMsg) {
	for e := self.handlers.Front(); e != nil; e = e.Next() {
		hander, ok := e.Value.(logHandler)
		if ok {
			hander.handle(msg)
		}
	}
}

func (self *Logger) reload() {

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
