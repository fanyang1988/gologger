package log

import (
    "fmt"
)

type roller2DateFile struct {
    path string
}

func (self roller2DateFile) init() {
    return
}

func (self roller2DateFile) roll(msg *logMsg) {
    return
}

func (self roller2DateFile) close() {
    return
}

func (self roller2DateFile) handle(msg *logMsg) error {
    fmt.Printf("datefile %s\n", msg.info)
    return nil
}

func (self roller2DateFile) setNext(n logHandler) {
    return
}
