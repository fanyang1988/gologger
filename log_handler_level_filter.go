package gologger

type levelFilter struct {
    level int
    next  logHandler
}

func (self *levelFilter) handle(msg *logMsg) error {
    if (msg.level >= self.level) && (self.next != nil) {
        return self.next.handle(msg)
    }

    return nil
}
func (self *levelFilter) setNext(n logHandler) {
    self.next = n
}
