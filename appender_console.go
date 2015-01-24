package log

import (
    "fmt"
)

type appender2Console struct {
}

func (self *appender2Console) init() error {
    return nil
}

func (self *appender2Console) close() {
    return
}

func (self *appender2Console) handle(msg *logMsg) error {
    fmt.Printf("%s\n", msg.info)
    return nil
}

func (self *appender2Console) setNext(n logHandler) {
    return
}
