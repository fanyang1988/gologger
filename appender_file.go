package log

import (
    "bufio"
    "os"
)

type appender2File struct {
    path string
    max  int64

    fd          *os.File
    writer      *bufio.Writer
    check_count int
}

func (self *appender2File) init() error {
    return self.checkfile()
}

func (self *appender2File) checkfile() error {
    self.check_count = 100
    if self.fd == nil {
        nfd, open_err := createLogFile(self.path, self.max, "")
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

        nfd, open_err := createLogFile(self.path, self.max, "")
        if open_err != nil {
            nfd.Close()
            return open_err
        }
        self.fd = nfd
        self.writer = bufio.NewWriter(self.fd)
    }
    return nil
}

func (self *appender2File) close() {
    if self.fd != nil {
        self.writer.Flush()
        self.fd.Close()
    }
    return
}

func (self *appender2File) handle(msg *logMsg) error {
    self.check_count--
    self.writer.WriteString(msg.info)

    if self.check_count <= 0 {
        self.writer.Flush()
        err := self.checkfile()
        if err != nil {
            return err
        }
    }
    return nil
}

func (self *appender2File) setNext(n logHandler) {
    return
}
