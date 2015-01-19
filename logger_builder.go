package log

import (
    "fmt"
    //sj "github.com/bitly/go-simplejson"
)

func buildLogger(name string, config map[string]interface{}) (*Logger, error) {
    fmt.Printf("name : %s\n", name)
    fmt.Printf("%s\n", config["type"])
    return nil, nil
}

func reBuildLogger(logger *Logger, config map[string]interface{}) error {
    return nil
}
