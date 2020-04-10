package game

import (
	"github.com/gin-gonic/gin"
	"luck_game/middleware"
	"luck_game/model"
	"luck_game/service"
	"luck_game/utils"
)

type LoginData struct {
	Id int `form:"id" json:"id"`
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var UserService = service.UserService{}
/**
 * 注册
 */
func Register(c *gin.Context)  {
	app := utils.Gin{C:c}
	data := make(map[string]interface{})
	var reg LoginData
	if err := c.Bind(&reg); err != nil {
		app.Response(0, err.Error(), data)
		return
	}

	_, err := UserService.Register(reg.Username, reg.Password)
	if err != nil {
		app.Response(0, err.Error(), data)
		return
	}

	app.Response(1, "ok", data)
}
/**
 * 登陆
 */
func Login(c *gin.Context)  {
	app := utils.Gin{C:c}
	data := make(map[string]interface{})
	var login LoginData
	if err := c.Bind(&login); err != nil {
		app.Response(0, err.Error(), data)
		return
	}

	user := model.User{}
	user, err := UserService.Login(login.Username, login.Password)
	if err != nil {
		app.Response(0, err.Error(), data)
		return
	}

	token, err := utils.GenerateToken(user.UserId, user.Username, user.Password)
	if err != nil {
		app.Response(0, err.Error(), data)
		return
	}

	data["token"] = token
	app.Response(1, "ok", data)
}

func Logout(c *gin.Context){

}

type Editpwd struct {
	OldPwd string `json:"old_pwd" form:"old_pwd" binding:"required"`
	NewPwd string `json:"new_pwd" form:"new_pwd" binding:"required"`
	ConfirmPwd string `json:"confirm_pwd" form:"confirm_pwd" binding:"required"`

}
/**
 * 更新密码
 */
func EditPwd (c *gin.Context){
	app := utils.Gin{C:c}
	data := make(map[string]interface{})
	var editPwd Editpwd
	if err := c.Bind(&editPwd); err != nil {
		app.Response(0, err.Error(), nil)
		return
	}
	if editPwd.NewPwd != editPwd.ConfirmPwd {
		app.Response(0, "两次确认密码不一制", data)
		return
	}

	// 验证原始密码
	user := model.User{}
	user, err := UserService.GetOneById(middleware.UserId)
	if err != nil {
		app.Response(0, err.Error(), data)
		return
	}
	if utils.Md5(editPwd.OldPwd) != user.Password {
		app.Response(0, "原始密码不正确", data)
		return
	}

	// 更新操作
	flag,err := UserService.EditPwd(middleware.UserId, editPwd.NewPwd)
	if err != nil {
		app.Response(0, err.Error(), data)
		return
	}
	if !flag {
		app.Response(0, err.Error(), data)
		return
	}

	app.Response(1, "更新成功", data)
}
