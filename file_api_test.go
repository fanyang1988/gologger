package log

import (
    "testing"
)

const (
    test_file_path1 = "./test/logger_test.log"
    test_file_path2 = "./test/logger_test1.log"
    test_file_path3 = "./test/logger_test1.log.1"
    test_file_path4 = "./test/logger_testN.log"
)

func TestFileExist(t *testing.T) {
    must_true := isFileExist(test_file_path1)
    must_false := isFileExist(test_file_path4)
    if !must_true {
        t.Errorf("%s is %s", test_file_path1, must_true)
        t.Fail()
    }
    if must_false {
        t.Errorf("%s is %s", test_file_path4, must_true)
        t.Fail()
    }
    return
}

func TestNewFilePath(t *testing.T) {
    new_file_1 := newFilePath(test_file_path1)
    t.Logf("new_file_1 %s", new_file_1)
    new_file_2 := newFilePath(test_file_path2)
    t.Logf("new_file_2 %s", new_file_2)
    new_file_3 := newFilePath(test_file_path3)
    t.Logf("new_file_3 %s", new_file_3)
    new_file_4 := newFilePath(test_file_path4)
    t.Logf("new_file_4 %s", new_file_4)
    return
}

func TestGetNowDateFilePath(t *testing.T) {
    time_format := "-2006-01-02.log"
    new_file_1 := getNowDateFilePath(test_file_path1, time_format)
    t.Logf("new_file_1 %s", new_file_1)
    new_file_2 := getNowDateFilePath(test_file_path2, time_format)
    t.Logf("new_file_2 %s", new_file_2)
    new_file_3 := getNowDateFilePath(test_file_path3, time_format)
    t.Logf("new_file_3 %s", new_file_3)
    new_file_4 := getNowDateFilePath(test_file_path4, time_format)
    t.Logf("new_file_4 %s", new_file_4)
    return
}
