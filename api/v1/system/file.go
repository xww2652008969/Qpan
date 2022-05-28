package system

import (
	"Qpan/model"
	"Qpan/utils"
	"github.com/gin-gonic/gin"
)

type Fileapi struct {
}

// Upload  传入文件，上传文件
func (fileapi *Fileapi) Upload(c *gin.Context) {
	var f model.File
	fi, _ := c.FormFile("file")
	username := c.Request.Header.Get("uuid")
	filedata, filename, er := utils.Getfile(fi)
	f.Name = filename
	f.Fileaddress = "work/" + username + "/" + filename
	f.Md5 = utils.Gedmd5(filedata)
	f.Useruuid = username
	if er == true {
		if fileservice.Pploadfiles(f, filedata) {
			model.Successres(nil, "ok", c)
		} else {
			model.Errorres(nil, "no", c)
		}
	} else {
	}
}

//获取用户所有文件
func (fileapi *Fileapi) Getfilelist(c *gin.Context) {
	uuid := c.Request.Header.Get("uuid")
	out, err := fileservice.Getfilelist(uuid)
	if err == nil {
		model.Successres(out, "ok", c)
	} else {
		model.Errorres(out, "no", c)
	}
}

// Getfile 下载文件
func (fileapi *Fileapi) Getfile(c *gin.Context) {
	fname := c.Query("name")
	username := c.Request.Header.Get("uuid")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	fileservice.Getfile("work/"+username+"/"+fname, c)
	return
}
