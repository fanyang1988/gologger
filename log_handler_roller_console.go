package log

import (
    "fmt"
)

type roller2Console struct {
}

func (self roller2Console) init() {
    return
}

func (self roller2Console) roll(msg *logMsg) {
    return
}

func (self roller2Console) close() {
    return
}

func (self roller2Console) handle(msg *logMsg) error {
    fmt.Printf("console %s\n", msg.info)
    return nil
}

func (self roller2Console) setNext(n logHandler) {
    return
}
