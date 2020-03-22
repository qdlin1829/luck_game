package model

type User struct {
	Id int64 `xorm:"pk autoincr bigint(20)" json:"id" from:"id"`
	Username string `xorm:"varchar(30)" json:"username" from:"username"`
	Password string `xorm:"varchar(30)" json:"Password" from:"Password"`
}
