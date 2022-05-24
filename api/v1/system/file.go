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
	username := c.Request.Header.Get("uuid")
	filedata, filename, er := utils.Getfile(fi)
	f.Fileaddress = "work/" + username + "/" + filename
	f.Md5 = utils.Gedmd5(filedata)
	f.Useruuid = username
	if er == true {
		if fileservice.Pploadfiles(f, filedata) {
			model.FailWithMessage(nil, "ok", c)
		} else {
			model.FailWithMessage(nil, "no", c)
		}
	} else {
	}
}

// Getfile 下载文件
func (Fileapi Fileapi) Getfile(c *gin.Context) {
	var f model.FileR
	_ = c.ShouldBindJSON(&f)
	c.File(f.Fileaddress)
	return
}