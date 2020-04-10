package game

import (
	"github.com/gin-gonic/gin"
	"luck_game/middleware"
	"luck_game/service"
)

var GameServie service.GameService

type GameData struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Content interface{} `json:"content"`
}


func Index(c *gin.Context){

	c.JSON(200, gin.H{
		"code":1,
		"message":"ok",
		"data":middleware.UserId,
	})
	return

}
