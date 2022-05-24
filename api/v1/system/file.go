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
func (fileapi *Fileapi) Getfilelist(c *gin.Context) {
	uuid := c.Request.Header.Get("uuid")
	out, err := fileservice.Getfilelist(uuid)
	if err == nil {
		model.FailWithMessage(out, "ok", c)
	} else {
		model.FailWithMessage(nil, "no", c)
	}
}

// Getfile 下载文件
func (fileapi *Fileapi) Getfile(c *gin.Context) {
	var f model.FileR
	_ = c.ShouldBindJSON(&f)
	c.File(f.Fileaddress)
	return
}
