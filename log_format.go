package log

import (
    "bytes"
    "fmt"
    "time"
)

const (
    date_format = "2006-01-02 15:04:05"
)

func format(level int, typ string, info string) string {
    //[2013-01-23 01:00:01][INFO][player] - something error!
    buf := &bytes.Buffer{}
    time_str := time.Now().Format(date_format)

    fmt.Fprintf(buf,
        "[%s][%s][%s] - %s",
        time_str,
        level_str[level],
        typ,
        info)
    return buf.String()
}
