package service

import (
	"errors"
	"log"
	"luck_game/model"
	"time"
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

func (u *UserService) Cache (key string, data interface{}, exptime time.Duration) (err error ){
	err = rdb.Set(key, data, exptime*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}


func (u *UserService) GetCache (key string) (ret string, err error ){
	val, err := rdb.Get(key).Result()
	if err != nil {
		return val, err
	}

	return val, nil
}
