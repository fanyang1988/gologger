package log

import (
    "container/list"
    "fmt"
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

var (
    level_str = []string{"All", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
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
        handlers: list.New(),
    }
    return new_logger
}

func CloseLogger(logger *Logger) {
    if logger != nil {
        logger.endLog()
    }
}

func (self *Logger) log(level int, format string, v []interface{}) {
    info := fmt.Sprintf(format, v...)
    msg := newLogMsg(level, self.name, info)
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

func (self *Logger) Debug(format string, v ...interface{}) {
    self.log(DEBUG, format, v)
}

func (self *Logger) Info(format string, v ...interface{}) {
    self.log(INFO, format, v)
}

func (self *Logger) Warn(format string, v ...interface{}) {
    self.log(WARN, format, v)
}

func (self *Logger) Error(format string, v ...interface{}) {
    self.log(ERROR, format, v)
}

func (self *Logger) Fatal(format string, v ...interface{}) {
    self.log(FATAL, format, v)
}
