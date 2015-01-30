package gologger

import (
    "errors"
    "fmt"
    "strconv"
    "strings"
    //sj "github.com/bitly/go-simplejson"
)

func getLogLevel(str string) int {
    str_upper := strings.ToUpper(str)
    switch str_upper {
    case "DEBUG":
        return DEBUG
    case "INFO":
        return INFO
    case "WARN":
        return WARN
    case "ERROR":
        return ERROR
    case "FATAL":
        return FATAL
    }

    return INFO
}

func getNewRoller(typ, path, max, format string) appender {
    typ_upper := strings.ToUpper(typ)
    max_int, int_ok := strconv.ParseInt(max, 10, 64)
    if int_ok != nil {
        max_int = 0
    }
    switch typ_upper {
    case "CONSOLE":
        return &appender2Console{}
    case "FILE":
        return &appender2File{
            path: path,
            max:  max_int,
        }
    case "DATEFILE":
        return &appender2DateFile{
            path:   path,
            max:    max_int,
            format: format,
        }
    }

    return &appender2Console{}
}

func buildLogger(logger *Logger, name string, config map[string]interface{}) (*Logger, error) {
    fmt.Printf("name : %s\n", name)

    log_level, log_level_ok := config["level"].(string)
    log_path, log_path_ok := config["path"].(string)
    log_format, log_format_ok := config["format"].(string)
    log_max, log_max_ok := config["max"].(string)
    if !log_level_ok {
        log_level = "INFO"
    }
    if !log_path_ok {
        log_path = "unpath.log"
    }
    if !log_format_ok {
        log_format = ""
    }
    if !log_max_ok {
        log_max = "0"
    }

    // log level
    log_lv_filter := &levelFilter{
        level: getLogLevel(log_level),
    }

    // roller
    log_type, log_type_ok := config["type"].(string)
    if !log_type_ok {
        return nil, errors.New("type error")
    }
    roller := getNewRoller(log_type, log_path, log_max, log_format)
    roller_init_err := roller.init()
    if roller_init_err != nil {
        roller.close()
        return nil, roller_init_err
    }
    log_lv_filter.setNext(roller)

    new_logger := logger
    if new_logger == nil {
        new_logger = NewLogger(name)
    }

    new_logger.handlers.PushBack(log_lv_filter)

    if logger == nil {
        go new_logger.startLog()
    }

    return new_logger, nil
}

func reBuildLogger(logger *Logger, config map[string]interface{}) error {
    return nil
}
