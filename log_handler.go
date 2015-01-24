package log

type logHandler interface {
    setNext(n logHandler)
    handle(msg *logMsg) error
}

type appender interface {
    logHandler
    init() error
    close()
}
