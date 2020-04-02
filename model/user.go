package model

type User struct {
	Id int64 `xorm:"pk autoincr bigint(20)" json:"id" from:"id"`
	Username string `xorm:"varchar(30)" json:"username" from:"username"`
	Password string `xorm:"varchar(30)" json:"Password" from:"Password"`
	CreateTime int64 `xorm:"int(11)" json:"create_time" from:"create_time"`
	UpdateTime int64 `xorm:"int(11)" json:"update_time" from:"update_time"`
}
