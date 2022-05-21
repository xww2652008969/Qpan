package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Writefile 写入文件，返回ture
func Writefile(filename string, chadate []byte) bool {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	defer f.Close()
	if err != nil {
		return false

	} else {
		f.Write(chadate)
		return true
	}
}

// Readfile 读取文件返回true和数据
func Readfile(filename string, b []byte) (bool, []byte) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 777)
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	if nil != err {
		fmt.Println("read", filename, "failed!")
		return false, content
	}
	return true, content
}
func Createfolder(dirname string) bool {
	err := os.MkdirAll(dirname, 777)
	if err != nil {
		return true
	} else {
		log.Print(err)
		return false
	}
}
