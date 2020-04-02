package game

import (
	"github.com/gin-gonic/gin"
	"luck_game/service"
)

func Index(c *gin.Context)  {

}

var GameServie service.GameService

type GameData struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Content interface{} `json:"content"`
}


func Set(c *gin.Context){
	ret,_ := GameServie.Set()


	c.JSON(200, gin.H{
		"code":1,
		"message":"ok",
		"data":ret,
	})
	return

}
