package log

import (
	"testing"
)

func TestBaseLogger(t *testing.T) {

	log_lv_filter := &levelFilter{
		level: INFO,
	}
	roll_to_console := roller2Console{}

	log_lv_filter.setNext(roll_to_console)

	logger2test := &Logger{
		name:     "test_log",
		level:    INFO,
		handlers: []logHandler{log_lv_filter},
	}

	logger2test.log(DEBUG, "test_log DEBUG")
	logger2test.log(INFO, "test_log INFO")
	logger2test.log(FATAL, "test_log FATAL")

	return
}

func TestLoggerFunc(t *testing.T) {

	log_lv_filter := &levelFilter{
		level: INFO,
	}
	roll_to_console := roller2Console{}

	log_lv_filter.setNext(roll_to_console)

	logger2test := &Logger{
		name:     "test_log",
		level:    INFO,
		handlers: []logHandler{log_lv_filter},
	}

	logger2test.Debug("test_log Debug")
	logger2test.Info("test_log Info")
	logger2test.Warn("test_log Warn")
	logger2test.Error("test_log Error")

	return
}
