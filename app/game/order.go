package game

import (
	"github.com/gin-gonic/gin"
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
