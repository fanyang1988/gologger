package log

type logHandler interface {
	setNext(n logHandler)
	handle(msg *logMsg) error
}

type roller interface {
	logHandler
	init()
	roll(msg *logMsg)
	close()
}
