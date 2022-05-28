package system

import (
	"Qpan/global"
	"Qpan/model"
	"Qpan/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Space
type Fileservice struct {
}

func (Fileservice Fileservice) Pploadfiles(f model.File, filedate []byte) bool {
	var file = model.File{}
	err := global.QP_db.Where("md5= ? and useruuid=?", f.Md5, f.Useruuid).First(&file).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	} else {
		global.QP_db.Create(&f)
		utils.Writefile(f.Fileaddress, filedate)
		return true
	}
}
func (Fileservice Fileservice) Getfilelist(uuid string) ([]model.Getfile, error) {
	var a model.Getfile
	var out []model.Getfile
	var fout []model.File
	err := global.QP_db.Select("name").Where("useruuid=?", uuid).Find(&fout).Error
	for _, value := range fout {
		t := value.Name
		a.Name = t
		out = append(out, a)

	}
	if err == nil {
		return out, nil
	}
	return out, errors.New("查找失败")
}

// Getfile 下载文件
func (Fileservice Fileservice) Getfile(filedir string, c *gin.Context) {
	c.File(filedir)

}
