package model

type OrderGoods struct {
	OgId int64 `xorm:"pk autoincr bigint(20)" json:"og_id", form:"og_id"`
	OrderId int64  `xorm:"bigint(20)" json:"order_id", form:"order_id"`
	ColorType int `xorm:"tinyint(1)" json:"color_type" form:"color_type"`
	Sq string `xorm:"varchar(20)" json:"sq" form:"sq"`
	Pos int `xorm:"tinyint(1)" json:"pos" form:"pos"`
	GoodsNum int `xorm:"int(10)" json:"goods_num" form:"goods_num"`
	Number string `xorm:"varchar(20)" json:"number" form:"number"`
	Amount float32 `xorm:"decimal(10,2)" json:"amount" form:"amount"`
}
