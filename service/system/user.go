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

// Register 用户注册
func (userservice *Userservice) Register(u *model.User) (err error, ok *model.User) {
	var user model.User
	if global.QP_db.Where("name=?", u.Name).First(&user).Error == nil {
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

// Login 用户登录
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
func (userservice *Userservice) Updateuser(u model.Updateuser) error {
	var user model.User
	u.Oldpassword = utils.Gedmd5([]byte(u.Oldpassword))
	u.Newpassword = utils.Gedmd5([]byte(u.Newpassword))
	out := global.QP_db.Model(&user).Where("uuid=? and password=?", u.Uuid, u.Oldpassword).Update("password", u.Newpassword)
	if out.RowsAffected == 1 {
		return nil
	} else {
		return errors.New("更改密码失败")
	}
}
