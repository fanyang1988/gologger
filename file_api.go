package log

import (
    "os"
    "strconv"
    "time"
)

func isFileExist(path string) bool {
    _, err := os.Stat(path)
    return err == nil
}

func getFileSize(path string) int64 {
    fstat, err := os.Stat(path)
    if err != nil {
        return 0
    } else {
        return fstat.Size()
    }
}

func newFilePath(path string) string {
    path_to_re := path
    num := 1
    for {
        if isFileExist(path_to_re) {
            path_to_re = path + "." + strconv.Itoa(num)
            num += 1
        } else {
            return path_to_re
        }
    }
}

func getNowDateFilePath(path, time_format string) string {
    now := time.Now()
    str := now.Format(time_format)
    return path + str
}

func getLogFile(path string, size_max int64, tformat string) string {
    path_to_open := path
    if tformat != "" {
        path_to_open = getNowDateFilePath(path, tformat)
    }

    if size_max > 0 && getFileSize(path_to_open) > size_max {
        path_to_open = newFilePath(path_to_open)
    }

    return path_to_open

}
