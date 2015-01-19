package log

import (
    "fmt"
)

type roller2File struct {
    path string
}

func (self roller2File) init() {
    return
}

func (self roller2File) roll(msg *logMsg) {
    return
}

func (self roller2File) close() {
    return
}

func (self roller2File) handle(msg *logMsg) error {
    fmt.Printf("file %s\n", msg.info)
    return nil
}

func (self roller2File) setNext(n logHandler) {
    return
}
