package service

import (
	"errors"
	"luck_game/model"
	"luck_game/utils"
	"time"
)

type UserService struct {

}

func (u *UserService) Register(username, password string) (user int64, err error) {

	tmp := model.User{}
	flag, err := Db.Table("g_user").Where("username=?", username).Get(&tmp)
	if err != nil {
		return 0,err
	}
	if flag {
		return tmp.UserId,errors.New("用户名存在")
	}

	tmp.Username = username
	tmp.Password =  utils.Md5(password)
	tmp.CreateTime = time.Now().Unix()
	tmp.UpdateTime = time.Now().Unix()

	Db.Table("g_user").Insert(&tmp)

	return tmp.UserId,err
}

func (u *UserService) Login (username, password string) (user model.User, err error) {
	tmp := model.User{}
	flag, err := Db.Table("g_user").Where("username=?", username).Get(&tmp)
	if err != nil {
		return tmp,err
	}
	if !flag {
		return tmp,errors.New("用户名不存在")
	}

	if utils.Md5(password) != tmp.Password {
		return tmp,errors.New("密码不正确")
	}

	return tmp, err
}

func (u *UserService)EditPwd (user_id int64,  passwd string)(bool,  error){
	tmp := model.User{}
	flag, err := Db.Table("g_user").Where("user_id=?", user_id).Get(&tmp)
	if err != nil {
		return false,err
	}
	if !flag {
		return false,errors.New("用户不存在")
	}

	tmp.Password = utils.Md5(passwd)
	data := map[string]string{}
	data["password"] = utils.Md5(passwd)
	_, err = Db.Table("g_user").Where("user_id=?", user_id).Update(data)
	if err != nil {
		return false,errors.New("密码更新失败")
	}

	return true,nil

}

func (u *UserService) GetOneById(user_id int64)(user model.User, err error){
	tmp := model.User{}
	flag, err := Db.Table("g_user").Where("user_id=?", user_id).Get(&tmp)
	if err != nil {
		return tmp,err
	}
	if !flag {
		return tmp,errors.New("用户不存在")
	}

	return tmp,nil
}