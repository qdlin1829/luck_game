package model

type Xyft struct {
	Id int64 `xorm:"pk autoincr bigint(20)" json:"id" from:"id"`
	Issue string `xorm:"varchar(15)" json:"issue" from:"issue"`
	Number string `xorm:"varchar(30)" json:"number" from:"number"`
	Sumfs int `xorm:"tinyint(1)" json:"sumfs" from:"sumfs"`
	Dx string `xorm:"char(4)" json:"dx" from:"dx"`
	Ds string `xorm:"char(4)" json:"ds" from:"ds"`
	NewIssue string `xorm:"varchar(15)" json:"new_issue" from:"new_issue"`
	NewOpenTime int64 `xorm:"int(11)" json:"new_open_time" from:"new_open_time"`
	CreateTime int64 `xorm:"int(11)" json:"create_time" from:"create_time"`

}
