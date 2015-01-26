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

func (self *Logger) log(level int, v []interface{}) {
    info := fmt.Sprint(v...)
    msg := newLogMsg(level, self.name, info)
    if msg != nil {
        self.log_wait.Add(1)
        self.log_conn <- msg
    }
}

func (self *Logger) logf(level int, format string, v []interface{}) {
    info := fmt.Sprintf(format, v...)
    msg := newLogMsg(level, self.name, info)
    if msg != nil {
        self.log_wait.Add(1)
        self.log_conn <- msg
    }
}

func (self *Logger) logln(level int, v []interface{}) {
    info := fmt.Sprintln(v...)
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
    self.logf(DEBUG, format, v)
}

func (self *Logger) Info(format string, v ...interface{}) {
    self.logf(INFO, format, v)
}

func (self *Logger) Warn(format string, v ...interface{}) {
    self.logf(WARN, format, v)
}

func (self *Logger) Error(format string, v ...interface{}) {
    self.logf(ERROR, format, v)
}

func (self *Logger) Fatal(v ...interface{}) {
    self.log(FATAL, v)
}

func (self *Logger) Fatalf(format string, v ...interface{}) {
    self.logf(FATAL, format, v)
}

func (self *Logger) Fatalln(v ...interface{}) {
    self.logln(FATAL, v)
}

func (self *Logger) Print(v ...interface{}) {
    self.log(INFO, v)
}
func (self *Logger) Printf(format string, v ...interface{}) {
    self.logf(INFO, format, v)
}
func (self *Logger) Println(v ...interface{}) {
    self.logln(INFO, v)
}

func (self *Logger) Panic(v ...interface{}) {
    self.log(FATAL, v)
}
func (self *Logger) Panicf(format string, v ...interface{}) {
    self.logf(FATAL, format, v)
}
func (self *Logger) Panicln(v ...interface{}) {
    self.logln(FATAL, v)
}

func (self *Logger) SetFlags(flag int) {

}
func (self *Logger) SetPrefix(prefix string) {

}
func (self *Logger) Flags() int {
    return 0
}
func (self *Logger) Output(calldepth int, s string) error {
    return nil
}
func (self *Logger) Prefix() string {
    return ""
}
