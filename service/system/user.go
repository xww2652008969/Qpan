package system

import (
	"Qpan/global"
	"Qpan/model"
	"Qpan/utils"
	"errors"
	"gorm.io/gorm"
)

//
type Userservice struct {
}

func (userservice *Userservice) Register(u *model.User) (err error, ok *model.User) {
	var user model.User
	if !errors.Is(global.QP_db.Where("name=?", u.Name).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已注册"), ok
	}
	u.Password = utils.Gedmd5([]byte(u.Password))
	err = global.QP_db.Create(&u).Error
	return err, u
}
