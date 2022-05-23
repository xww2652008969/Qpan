package system

import (
	"Qpan/global"
	"Qpan/model"
	"Qpan/utils"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

//
type Userservice struct {
}

//注册函数
func (userservice *Userservice) Register(u *model.User) (err error, ok *model.User) {
	var user model.User
	if global.QP_db.Where("name=?", u.Name).Find(&user).Error != nil {
		return errors.New("用户名已注册"), ok
	} else {
		u.Uuid = uuid.Must(uuid.NewRandom())
		u.Filespace = "work/" + fmt.Sprintf("%v", u.Uuid)
		u.Password = utils.Gedmd5([]byte(u.Password))
		err = global.QP_db.Create(&u).Error
		utils.Createfolder(u.Filespace)
		return err, u
	}
}

//用户登录
func (Userservice *Userservice) Login(u *model.User) (error, model.User) {
	var user model.User
	u.Password = utils.Gedmd5([]byte(u.Password))
	err := global.QP_db.Where("name=? and password=?", u.Name, u.Password).First(&user).Error
	if err != nil {
		fmt.Println("密码错误")
		return errors.New("密码错误"), model.User{}
	}
	//返回空值
	return nil, user
}

//
