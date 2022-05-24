package system

import (
	"Qpan/global"
	"Qpan/model"
	"Qpan/utils"
	"errors"
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
func (Fileservice Fileservice) Getfilelist(uuid string) ([]model.File, error) {
	var fout []model.File
	err := global.QP_db.Select("fileaddress").Where("useruuid=?", uuid).Find(&fout).Error
	if err == nil {
		return fout, nil
	}
	return fout, errors.New("查找失败")
}
