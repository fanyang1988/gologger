package log

import (
    "fmt"
)

type appender2DateFile struct {
    path   string
    max    int64
    format string
}

func (self *appender2DateFile) init() error {
    return nil
}

func (self *appender2DateFile) close() {
    return
}

func (self *appender2DateFile) handle(msg *logMsg) error {
    fmt.Printf("datefile %s\n", msg.info)
    return nil
}

func (self *appender2DateFile) setNext(n logHandler) {
    return
}
