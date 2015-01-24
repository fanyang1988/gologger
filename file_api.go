package log

import (
    "fmt"
    "os"
    "path"
    "strconv"
    "time"
)

func isFileExist(path_log string) bool {
    _, err := os.Stat(path_log)
    return err == nil
}

func isCanToLog(path_log string, max_size int64) bool {
    fstat, err := os.Stat(path_log)
    if err != nil {
        return true
    } else {
        return fstat.Size() <= max_size
    }
}

func getFileSize(path_log string) int64 {
    fstat, err := os.Stat(path_log)
    if err != nil {
        return 0
    } else {
        return fstat.Size()
    }
}

func newFilePath(path_log string, max_size int64) string {
    path_to_re := path_log
    num := 1
    for {
        if !isCanToLog(path_to_re, max_size) {
            path_to_re = path_log + "." + strconv.Itoa(num)
            num += 1
        } else {
            return path_to_re
        }
    }
}

func getNowDateFilePath(path_log, time_format string) string {
    now := time.Now()
    str := now.Format(time_format)
    return path_log + str
}

func getLogFile(path_log string, size_max int64, tformat string) string {
    path_to_open := path_log
    if tformat != "" {
        path_to_open = getNowDateFilePath(path_log, tformat)
    }

    if size_max > 0 {
        path_to_open = newFilePath(path_to_open, size_max)
    }

    return path_to_open

}

func createLogFile(path_log string, size_max int64, tformat string) (*os.File, error) {
    path_to_create := getLogFile(path_log, size_max, tformat)
    if !isFileExist(path_to_create) {
        path_dir := path.Dir(path_to_create)
        os.MkdirAll(path_dir, os.ModeDir)
        return os.Create(path_to_create)
    } else {
        f, err := os.OpenFile(path_to_create, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
        fmt.Printf("%s\n", f)
        return f, err
    }
}
