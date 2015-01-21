package log

import (
    "errors"
    "fmt"
    "github.com/fanyang1988/goconfig"
    "sync"
)

type Log struct {
    loggers     map[string]*Logger
    config_chan chan string
    config      *goconfig.Config

    mutex sync.RWMutex

    config_path string
    config_name string
}

func NewLog(config_name, config_path string, config_mng *goconfig.Config) *Log {
    new_log := &Log{
        config_chan: make(chan string),
        loggers:     make(map[string]*Logger),
        config:      config_mng,
        config_path: config_path,
        config_name: config_name,
    }
    return new_log
}

func (self *Log) Init() error {
    self.config.Reg(self.config_name, self.config_path, true)
    self.config.RegNotifyChan(self.config_chan)

    config_info := self.config.Get(self.config_name)
    logger_info_arr, err := config_info.Get("logger").Array()

    if err != nil {
        return errors.New("config err : logger_info_arr not array")
    }

    fmt.Printf("%s\n", logger_info_arr)

    // init logger
    for _, e := range logger_info_arr {

        logger_info, info_ok := e.(map[string]interface{})

        if !info_ok {
            continue
        }

        logger_name := logger_info["name"]

        if logger_name == nil {
            continue
        }

        logger_name_str, name_ok := logger_name.(string)

        if !name_ok {
            continue
        }

        logger := self.loggers[logger_name_str]

        new_logger, build_err := buildLogger(logger, logger_name_str, logger_info)

        if build_err != nil {
            continue
        }

        self.loggers[logger_name_str] = new_logger

    }

    go self.logMain()
    return nil
}

func (self *Log) Close() {
    close(self.config_chan)
}

func (self *Log) logMain() {
    for {
        select {
        case config_update_name := <-self.config_chan:
            if config_update_name == self.config_name {
                self.reloadAllLogger()
                //TODO Read New Logger Defined
            }
        }
    }
}

func (self *Log) reloadAllLogger() {
    self.mutex.Lock()
    defer self.mutex.Unlock()

    for _, logger := range self.loggers {
        logger.log_conn <- nil
    }
}

func (self *Log) GetLogger(name string) *Logger {
    self.mutex.RLock()
    defer self.mutex.RUnlock()

    if self.loggers[name] != nil {
        return self.loggers[name]
    }

    return nil
}
