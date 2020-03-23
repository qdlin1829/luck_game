package main

import (
	"github.com/gin-gonic/gin"
	"luck_game/app/shop"
	"luck_game/app/websocket"
	"luck_game/middleware"
)


func main()  {
	gin.SetMode("debug")
	r := gin.Default()

	//r.POST("/index", shop.Index)
	r.POST("/login", shop.Login)
	r.POST("/register", shop.Register)
	//r.POST("/goods_list", shop.GoodsList)
	//r.POST("/sms_send", shop.GoodsList)
	//r.POST("/goods_info", shop.GoodsInfo)
	r.GET("/ws", websocket.Ws)

	shopApi := r.Group("")
	shopApi.Use(middleware.AuthMiddleware())
	{
		shopApi.POST("/user_info", shop.Update)
		//shopApi.POST("/user_info", shop.GetInfo)
		//shopApi.POST("/user_update", shop.Update)
		//shopApi.POST("/cart_add", shop.Update)
		//shopApi.POST("/cart_update", shop.Update)
		//shopApi.POST("/car_get_num", shop.Update)
		//shopApi.POST("/buy_checkout", shop.Update)
		//shopApi.POST("/buy_short", shop.Update)

		// 收货地址
		//shopApi.POST("/address_list", shop.Update)
		//shopApi.POST("/address_add", shop.Update)
		//shopApi.POST("/address_info", shop.Update)

	}


	r.Run(":8080")
}