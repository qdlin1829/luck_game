package service

import (
	"github.com/go-xorm/xorm"
	"luck_game/model"
)

type OrderService struct {

}

type OrderArg struct {

}



func (o *OrderService)GetOrderCount(orderArg model.OrderSerach)(int64, error){
	db  := o.BuildCond(orderArg)
	count, err := db.Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}


func (o *OrderService)GetOrderList(orderArg model.OrderSerach)([]*model.OrderGoodsList, error){
	orders := []model.OrderInfo{}


	db  := o.BuildCond(orderArg)
	err := db.Limit(orderArg.PageSize, (orderArg.Page-1)*orderArg.Page).Find(&orders)
	if err != nil {
		return nil,err
	}

	var order_ids []int
	for i:=0; i<len(orders); i++ {
		order_ids = append(order_ids, int(orders[i].OrderId))
	}

	goods := []model.OrderGoods{}
	Db.Table("g_order_goods").In("order_id", order_ids).Find(&goods)

	order_goods := []*model.OrderGoodsList{}
	for _,val := range orders {
		OrderGoodsList := &model.OrderGoodsList{
			OrderInfo: val,
		}

		for _,vv := range goods  {
			if val.OrderId == vv.OrderId {
				OrderGoodsList.OrderGoods = []model.OrderGoods{vv}
				break
			}
		}

		order_goods = append(order_goods, OrderGoodsList)
	}

	return order_goods, nil
}


func (o *OrderService)BuildCond(arg model.OrderSerach)(* xorm.Session){

	t := Db.Table("g_order_info").Where("1=1")
	if 0 <len(arg.OrderSn) {
		t.And("order_sn=?",  arg.OrderSn)
	}
	if arg.OrderStatus > 0 {
		t.And("order_status=?",  arg.OrderStatus)
	}
	if arg.UserId > 0 {
		t.And("user_id=?",  arg.UserId)
	}

	return t
}
