package game

import (
	"github.com/gin-gonic/gin"
	"luck_game/middleware"
	"luck_game/model"
	"luck_game/service"
	"luck_game/utils"

	"github.com/gin-gonic/gin/binding"
)

var OrderService  = service.OrderService{}



func GetOrderList(c *gin.Context){
	app := utils.Gin{C:c}
	data := make(map[string]interface{})
	var order model.OrderSerach
	if err := c.ShouldBindWith(&order, binding.FormPost); err != nil {
		app.Response(0, "参数错误", data)
		return
	}
	order.UserId = int(middleware.UserId)
	count, err := OrderService.GetOrderCount(order)
	if err != nil {
		app.Response(0, err.Error(), data)
		return
	}

	if count == 0 {
		data["count"] = count
		data["list"]  = make(map[int]interface{})
		app.Response(1, "暂无数据", data)
		return
	}

	list, _ := OrderService.GetOrderList(order)

	data["count"] = count
	data["list"]  = list
	app.Response(1, "ok", data)
}


type BuyData struct {
	Sq string `form:"sq" json:"sq"`
	OrderType int `form:"order_type"  json:"order_type"`
	Pos []int `form:"pos" json:"pos"`
	Odds []float32 `json:"odds" form:"odds"`
	Number []string `json:"number" form:"number"`
	Amount float32 `json:"amount" form:"amount"`
}

func Buy(c *gin.Context){
	app := utils.Gin{C:c}
	data:= make(map[string]interface{})

	buy :=  BuyData{}
	if err := c.Bind(&buy); err != nil {
		app.Response(0, "请选择", data)
		return
	}



	app.Response(1, "ok", data)
}

