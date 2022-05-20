package system

import (
	"Qpan/global"
	"Qpan/model"
	"Qpan/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

//
type Userservice struct {
}

//注册函数
func (userservice *Userservice) Register(u *model.User) (err error, ok *model.User) {
	var user model.User
	if errors.Is(global.QP_db.Where("name=?", u.Name).Find(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已注册"), ok
	}
	u.Password = utils.Gedmd5([]byte(u.Password))
	err = global.QP_db.Create(&u).Error
	return err, u
}

//用户登录
func (Userservice *Userservice) Login(u *model.User) (err error, ok *model.User) {
	var user model.User
	u.Password = utils.Gedmd5([]byte(u.Password))
	err = global.QP_db.Where("name=? and password=?", u.Name, u.Password).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) == true {
		fmt.Println("密码错误")
		return errors.New("密码错误"), ok
	}
	//返回空值
	return nil, u
}
