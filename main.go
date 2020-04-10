package main

import (
	"github.com/gin-gonic/gin"
	"luck_game/app/game"
	"luck_game/app/shop"
	"luck_game/app/socket"
	"luck_game/middleware"
)

func main() {
	gin.SetMode("debug")
	r := gin.Default()

	//r.POST("/index", shop.Index)
	//r.POST("/login", shop.Login)
	//r.POST("/register", shop.Register)
	//r.POST("/goods_list", shop.GoodsList)
	//r.POST("/sms_send", shop.GoodsList)
	//r.POST("/goods_info", shop.GoodsInfo)
	r.GET("/ws", socket.Ws)
	r.POST("/sms_send", shop.Send)
	r.GET("/issue", shop.Issue)

	//shopApi := r.Group("")
	//shopApi.Use(middleware.AuthMiddleware())
	//{
	//	shopApi.POST("/user_info", shop.UserInfo)
	//	shopApi.POST("/user_withdraw", shop.Withdraw)
	//	shopApi.POST("/user_editpass", shop.EditPasswd)
	//	shopApi.POST("/user_info", shop.GetInfo)
	//	shopApi.POST("/user_update", shop.Update)
	//	shopApi.POST("/cart_add", shop.Update)
	//	shopApi.POST("/cart_update", shop.Update)
	//	shopApi.POST("/car_get_num", shop.Update)
	//	shopApi.POST("/buy_checkout", shop.Update)
	//	shopApi.POST("/buy_short", shop.Update)
	//
	//	//收货地址
	//	shopApi.POST("/address_list", shop.Update)
	//	shopApi.POST("/address_add", shop.Update)
	//	shopApi.POST("/address_info", shop.Update)
	//
	//}

	r.POST("/game/register", game.Register)
	r.POST("/game/login", game.Login)

	gameApi := r.Group("/game").Use(middleware.AuthMiddleware())
	{
		gameApi.GET("/index", game.Index)
		gameApi.POST("/editPwd", game.EditPwd)
		gameApi.POST("/getOrderList", game.GetOrderList)
		gameApi.POST("/buy", game.Buy)

	}

	// i := []int{1}
	//a := config.Encrypt("222A", 8, i)

	//b := config.Decrypt("222A", 8, "YmxmvwKn")

	// var a = []int{1,2,3,4,5,6}
	// var str = []string{}
	// for i:=0; i<len(a); i++ {

	// 	str  = append(str, fmt.Sprintf("%d", a[i]))
	// }
	// fmt.Println(strings.Join(str, ","))

	r.Run(":8080")
}
