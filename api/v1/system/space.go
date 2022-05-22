package system

import (
	"Qpan/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Upload(c gin.Context) {
	f, _ := c.FormFile("file")
	username := c.Request.Header.Get("name")
	filename, filedate, er := utils.Getfile(f)
	if er == true {
		fmt.Println(username, filename, filedate)
	}
}
