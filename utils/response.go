package utils

import "github.com/gin-gonic/gin"

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode int, Msg string, data interface{}) {
	g.C.JSON(200, Response{
		Code: httpCode,
		Msg:  Msg,
		Data: data,
	})

	return
}