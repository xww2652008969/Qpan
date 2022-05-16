package controllers

import (
	"fmt"
	"io/ioutil"
	"os"
)

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
