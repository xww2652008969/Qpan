package main

import (
	"Qpan/controllers"
	"fmt"
)

func main() {
	str := string("振德好屑")
	controllers.Writefile("xww", []byte(str))
	fmt.Printf("%x", controllers.Gedmd5([]byte(str)))

}
