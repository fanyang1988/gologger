package log

import (
    "bufio"
    "os"
)

type appender2DateFile struct {
    path   string
    max    int64
    format string

    fd     *os.File
    writer *bufio.Writer
}

func (self *appender2DateFile) init() error {
    return self.checkfile()
}

func (self *appender2DateFile) checkfile() error {
    if self.fd == nil {
        nfd, open_err := createLogFile(self.path, self.max, self.format)
        if open_err != nil {
            nfd.Close()
            return open_err
        }
        self.fd = nfd
        self.writer = bufio.NewWriter(self.fd)
    } else {
        fstat, err := self.fd.Stat()
        if err == nil && fstat.Size() <= self.max {
            return nil
        }
        self.writer.Flush()
        self.fd.Close()
        self.fd = nil

        nfd, open_err := createLogFile(self.path, self.max, self.format)
        if open_err != nil {
            nfd.Close()
            return open_err
        }
        self.fd = nfd
        self.writer = bufio.NewWriter(self.fd)
    }
    return nil
}

func (self *appender2DateFile) close() {
    if self.fd != nil {
        self.writer.Flush()
        self.fd.Close()
    }
    return
}

func (self *appender2DateFile) handle(msg *logMsg) error {
    self.writer.WriteString(msg.info)

    self.writer.Flush()
    err := self.checkfile()
    if err != nil {
        return err
    }

    return nil
}

func (self *appender2DateFile) setNext(n logHandler) {
    return
}
