package log

import (
    "fmt"
    "os"
)

type roller2File struct {
    path string
    max  int64

    fd  *os.File
}

func (self *roller2File) init() error {
    return self.checkfile()
}

func (self *roller2File) checkfile() error {
    if self.fd == nil {
        nfd, open_err := createLogFile(self.path, self.max, "")
        if open_err != nil {
            nfd.Close()
            return open_err
        }
        self.fd = nfd
    } else {
        fstat, err := self.fd.Stat()
        if err == nil && fstat.Size() <= self.max {
            return nil
        }
        self.fd.Close()
        self.fd = nil

        nfd, open_err := createLogFile(self.path, self.max, "")
        if open_err != nil {
            nfd.Close()
            return open_err
        }
        self.fd = nfd
    }
    return nil
}

func (self *roller2File) roll(msg *logMsg) {

    return
}

func (self *roller2File) close() {
    if self.fd != nil {
        self.fd.Close()
    }
    return
}

func (self *roller2File) handle(msg *logMsg) error {
    err := self.checkfile()
    if err != nil {
        return err
    }
    fmt.Printf("log info %s\n", msg.info)
    self.fd.WriteString(msg.info)
    return nil
}

func (self *roller2File) setNext(n logHandler) {
    return
}
