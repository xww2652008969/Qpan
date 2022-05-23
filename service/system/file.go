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
	fliemd5 := utils.Gedmd5(filedate)
	err := global.QP_db.Where("Fileaddress=?", fliemd5).First(&file).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	} else {
		global.QP_db.Create(f)
		utils.Writefile(f.Fileaddress, filedate)
		return true
	}
}
