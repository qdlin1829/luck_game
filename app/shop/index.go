package shop

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"luck_game/service"
	"net/http"
)

var SmsService  service.SmsService

func Index(c *gin.Context)  {

}

func Send (c *gin.Context){
	mobile := c.PostForm("mobile")
	msg := c.PostForm("message")

	ok, err := SmsService.SmsSend(mobile, msg)
	if ok {
		c.String(200, "发送成功")
	}else{
		c.String(200, err.Error())
	}
}

type Mod struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Content interface{}  `json:"content"`
}

func Issue(c *gin.Context){
	url := "https://aappii.1122xxx.com/lottery-client-api/races/min/10002?issue=0"
	resp,err := http.Get(url)
	if err != nil{
		c.JSON(200, gin.H{
			"code":0,
			"msg":err,
			"data":"",
		})
		return
	}

	defer resp.Body.Close()

	// 返回数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(200, gin.H{
			"code":0,
			"msg":err,
			"data":"",
		})
		return
	}

	str := []byte(string(body))
	mod := &Mod{}
	err = json.Unmarshal(str, mod)
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"code":200,
		"msg":"ok",
		"data":mod.Content,
	})

}