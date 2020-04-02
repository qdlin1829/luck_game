package shop

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"luck_game/config"
	"luck_game/service"
	"math/rand"
	"strconv"
)



var UserService service.UserService


func Login (c *gin.Context){
	username := c.PostForm("username")
	password := c.PostForm("password")

	info, err := UserService.Login(username, password)
	if err != nil {
		err.Error()
		return
	}

	c.JSON(200, gin.H{
		"code" : 200,
		"data":info,
		"msg":"ok",
	})
}

func Register (c *gin.Context) {

	username := rand.Int31()
	password := config.Md5("123456")

	ok, err := UserService.Register(strconv.Itoa(int(username)), password)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(200, gin.H{
		"code" : 200,
		"data":ok,
		"msg":"ok",
	})
}