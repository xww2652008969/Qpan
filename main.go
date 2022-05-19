package main

import (
	"Qpan/utils"
	"fmt"
)

func main() {
	str := string("振德好屑")
	utils.Writefile("xww", []byte(str))
	fmt.Printf("%x", utils.Gedmd5([]byte(str)))

}
