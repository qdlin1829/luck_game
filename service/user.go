package service

import (
	"errors"
	"log"
	"luck_game/model"
)

type UserService struct {

}

func (u *UserService) Register(username, password string) (user model.User, err error) {

	tmp := model.User{}
	_, err = Db.Table("go_user").Where("username=?", username).Get(&tmp)
	if err != nil {
		log.Fatal(err)
		return tmp,err
	}

	if tmp.Id>0 {
		return tmp,errors.New("用户名已注册")
	}

	tmp.Username = username
	tmp.Password =  password

	_, err = Db.Table("go_user").InsertOne(&tmp)

	return tmp,err
}

func (u *UserService) Login (username, password string) (user model.User, err error) {
	tmp := model.User{}
	_, err = Db.Table("go_user").Where("username=?", username).Get(&tmp)
	if err != nil {
		log.Fatal(err)
		return tmp,err
	}

	if tmp.Id == 0 {
		return tmp,errors.New("用户名不存在")
	}

	if password != tmp.Password {
		return tmp,errors.New("密码不正确")
	}

	return tmp, err

}
