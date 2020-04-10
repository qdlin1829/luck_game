package model

type OrderInfo struct {
	OrderId int64 `xorm:"pk autoincr bigint(20)" json:"order_id" from:"order_id" `
	OrderSn string `xorm:"varchar(30)  notnull unique" json:"order_sn" from:"order_sn"`
	OrderType int `xorm:"tinyint(1)" json:"order_type" from:"order_type"`
	OrderStatus int `xorm:"int(1)" json:"order_status" from:"order_status"`
	UserId int `xorm:"int(11)" json:"user_id" from:"user_id"`
	CreateTime int64 `xorm:"int(11)" json:"create_time" from:"create_time"`
	UpdateTime int64 `xorm:"int(11)" json:"update_time" from:"update_time"`
}

type OrderGoodsList struct {
	OrderInfo
	OrderGoods []OrderGoods
}
