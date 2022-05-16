package controllers

import (
	md52 "crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

//写入文件，返回ture
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

//读取文件返回true和数据
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

//获取md5
func Gedmd5(date []byte) [16]byte {
	md5 := md52.Sum(date)
	return md5
}
