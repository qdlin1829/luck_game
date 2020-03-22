package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"luck_game/middleware"
	"luck_game/service"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

var UserService service.UserService

func main()  {
	gin.SetMode("debug")
	r := gin.Default()

	r.GET("/", func(g *gin.Context) {
		g.JSON(200, gin.H{
			"messag":"welcome test",
		})
	})

	api := r.Group("/api")


	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/test", Test )
		api.GET("/v1", V1)
	}
	
	r.POST("/login", func(g *gin.Context) {
			username := g.PostForm("username")
			password := g.PostForm("password")
			info, err := UserService.Login(username, password)
			if err != nil {
				g.JSON(200, gin.H{
					"code": 0,
					"data":"",
					"msg": err.Error(),
				})
				return
			}

		g.JSON(200, gin.H{
			"code": 200,
			"data":info,
			"msg":"ok",
		})

	})
	
	r.GET("/getCache", func(g *gin.Context) {
		val, err := UserService.GetCache("user")
		if err != nil {
			g.JSON(200, gin.H{
				"code": 0,
				"data":"",
				"msg": err.Error(),
			})
			return
		}

		g.JSON(200, gin.H{
			"code": 200,
			"data":val,
			"msg":"ok",
		})
	})
	
	r.GET("/cahce", func(g *gin.Context) {
		err := UserService.Cache("user", "3333", 700)
		if err != nil {
			g.JSON(200, gin.H{
				"code": 0,
				"data":"",
				"msg": err.Error(),
			})
			return
		}

		g.JSON(200, gin.H{
			"code": 200,
			"data":"",
			"msg":"ok",
		})
	})
	


	r.GET("/ws", func(g *gin.Context) {

			ws,err := upGrader.Upgrade(g.Writer, g.Request, nil )
			if err != nil {
				return
			}
			fmt.Println(&ws)
			defer ws.Close()

			for {
				//读取ws中的数据
				mt, message, err := ws.ReadMessage()
				if err != nil {
					break
				}
				if string(message) == "ping" {
					message = []byte("pong")
				}

				//写入ws数据
				err = ws.WriteMessage(mt, message)
				if err != nil {
					break
				}
			}
	})


	r.Run(":8080")
}

func Test(c *gin.Context){
	c.JSON(200, gin.H{
		"message":"hello Test",
	})
}

func V1(c *gin.Context){
	c.JSON(200, gin.H{
		"message":"hello V1",
	})
}
