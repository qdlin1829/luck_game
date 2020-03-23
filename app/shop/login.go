package shop

import(
	"github.com/gin-gonic/gin"
	"luck_game/service"
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

}