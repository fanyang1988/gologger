package log

import (
    "github.com/fanyang1988/goconfig"
    "testing"
)

func TestLogStart(t *testing.T) {

    configMng := goconfig.NewConfig()

    logMng := NewLog("logger", "log_config.json", configMng)

    logMng.Init()
    logMng.GetLogger("info").Info("ddddddddddd1")
    logMng.GetLogger("info").Info("ddddddddddd2")
    logMng.GetLogger("info").Info("ddddddddddd3")
    logMng.GetLogger("info").Warn("ddddddddddd warn")
    logMng.GetLogger("info").Debug("ddddddddddd debug")
    logMng.Close()

    return
}
