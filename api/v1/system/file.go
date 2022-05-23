package system

import (
	"Qpan/model"
	"Qpan/utils"
	"github.com/gin-gonic/gin"
)

type Fileapi struct {
}

func (fileapi *Fileapi) Upload(c *gin.Context) {
	var f model.File
	fi, _ := c.FormFile("file")
	username := c.Request.Header.Get("name")
	filedata, filename, er := utils.Getfile(fi)
	f.Fileaddress = "work/" + username + "/" + filename
	f.Md5 = utils.Gedmd5(filedata)
	if er == true {
		if fileservice.Pploadfiles(f, filedata) {
			model.FailWithMessage(nil, "ok", c)
		} else {
			model.FailWithMessage(nil, "no", c)
		}
	} else {
	}
}
