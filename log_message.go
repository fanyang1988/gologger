package log

type logMsg struct {
    level int
    info  string
}

func newLogMsg(level int, typ, info string) *logMsg {
    return &logMsg{
        level: level,
        info:  format(level, typ, info),
    }
}
