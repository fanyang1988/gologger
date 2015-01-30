package log

import (
    "github.com/fanyang1988/goconfig"
    "testing"
)

func TestLogStart(t *testing.T) {

    configMng := goconfig.NewConfig()

    logMng := NewLog("logger", "log_config.json", configMng)

    logMng.GetLogger("info").Info("ddddddddddd1 %s %s", "infodd", "sssss")
    logMng.GetLogger("info").Info("ddddddddddd2")
    logMng.GetLogger("info").Info("ddddddddddd3")
    logMng.GetLogger("info").Warn("ddddddddddd warn")
    logMng.GetLogger("info").Debug("ddddddddddd debug")
    logMng.Close()

    return
}

func TestLogFile(t *testing.T) {

    configMng := goconfig.NewConfig()

    logMng := NewLog("logger", "log_config.json", configMng)

    logMng.GetLogger("test").Info("ddddddddddd1 %s %s", "infodd", "sssss")
    logMng.GetLogger("test").Info("ddddddddddd2")
    logMng.GetLogger("test").Info("ddddddddddd3")
    logMng.GetLogger("test").Warn("ddddddddddd warn")
    logMng.GetLogger("test").Debug("ddddddddddd debug")

    num := 1
    for {
        num += 1
        if num >= 3000 {
            break
        }
        logMng.GetLogger("test").Info("11111111111111 debug %d", num)
        logMng.GetLogger("info").Info("12121212 debug %d", num)
    }

    logMng.Close()

    return
}
