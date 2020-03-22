package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	// 可自定义盐值
	TokenSalt = "default_salt"
)

func MD5(data []byte) string {
	_md5 := md5.New()
	_md5.Write(data)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		token := c.Request.Header.Get("token")
		if token == ""{
			c.Abort()
			c.JSON(http.StatusUnauthorized,gin.H{"message":"访问未授权"})
			return
		}


		//c.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())
		c.Next()
	}
}