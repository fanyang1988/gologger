package log

import (
    "container/list"
    "testing"
)

func TestBaseLogger(t *testing.T) {

    log_lv_filter := &levelFilter{
        level: INFO,
    }
    roll_to_console := &appender2Console{}

    log_lv_filter.setNext(roll_to_console)

    logger2test := &Logger{
        name:     "test_log",
        handlers: list.New(),
        log_conn: make(chan *logMsg),
    }

    logger2test.handlers.PushBack(log_lv_filter)

    go logger2test.startLog()

    CloseLogger(logger2test)

    return
}

func TestLoggerFunc(t *testing.T) {

    log_lv_filter := &levelFilter{
        level: INFO,
    }
    roll_to_console := &appender2Console{}

    log_lv_filter.setNext(roll_to_console)

    logger2test := &Logger{
        name:     "test_log",
        handlers: list.New(),
        log_conn: make(chan *logMsg),
    }

    logger2test.handlers.PushBack(log_lv_filter)

    go logger2test.startLog()

    logger2test.Debug("test_log Debug")
    logger2test.Info("test_log Info")
    logger2test.Warn("test_log Warn")
    logger2test.Error("test_log Error")

    CloseLogger(logger2test)

    return
}
