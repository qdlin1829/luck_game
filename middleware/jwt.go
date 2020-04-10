package middleware

import (
	"github.com/gin-gonic/gin"
	"luck_game/utils"
	"time"
)

const (
	// 可自定义盐值
	TokenSalt = "default_salt"
)
var UserId int64 = 0

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		app := utils.Gin{C:c}
		token := c.Request.Header.Get("token")
		if token == ""{
			c.Abort()
			app.Response(0, "访问未授权", nil )
			return
		}

		ret, err := utils.ParseToken(token)
		if err != nil {
			c.Abort()
			app.Response(0, "token无效", nil )
			return
		}
		if time.Now().Unix() > ret.ExpiresAt {
			c.Abort()
			app.Response(0, "token过期请重新登陆", nil )
			return
		}

		UserId = ret.UserId

		//c.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())
		c.Next()

	}
}