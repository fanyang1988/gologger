package log

import (
    "errors"
    "fmt"
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

func getNewRoller(typ, path string) roller {
    typ_upper := strings.ToUpper(typ)
    switch typ_upper {
    case "CONSOLE":
        return roller2Console{}
    case "FILE":
        return roller2File{path: path}
    case "DATEFILE":
        return roller2DateFile{path: path}
    }

    return roller2Console{}
}

func buildLogger(logger *Logger, name string, config map[string]interface{}) (*Logger, error) {
    fmt.Printf("name : %s\n", name)

    log_level, log_level_ok := config["level"].(string)
    log_path, log_path_ok := config["path"].(string)
    if !log_level_ok {
        log_level = "INFO"
    }
    if !log_path_ok {
        log_path = "unpath.log"
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
    roller := getNewRoller(log_type, log_path)
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
