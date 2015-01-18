package log

type logMsg struct {
	level int
	info  string
}

func newLogMsg(level int, info string) *logMsg {
	return &logMsg{
		level: level,
		info:  info,
	}
}