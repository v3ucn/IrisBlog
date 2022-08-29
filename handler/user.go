package handler

import (
	"IrisBlog/database"
	"IrisBlog/model"
	"IrisBlog/mytool"

	"github.com/dchest/captcha"

	"github.com/kataras/iris/v12"
)

//用户登录模板
func User_signin(ctx iris.Context) {

	ctx.View("/signin.html")

}

//登录动作
func Signin(ctx iris.Context) {

	ret := make(map[string]interface{}, 0)

	cid := ctx.PostValue("cid")
	code := ctx.PostValue("code")

	if captcha.VerifyString(cid, code) == false {

		ret["errcode"] = 2
		ret["msg"] = "登录失败,验证码错误"
		ctx.JSON(ret)
		return

	}

	db := database.Db()
	defer func() {
		_ = db.Close()
	}()

	Username := ctx.PostValue("username")
	Password := ctx.PostValue("password")

	user := &model.User{}

	db.Where(&model.User{Username: Username, Password: mytool.Make_password((Password))}).First(&user)

	if user.ID == 0 {

		ret["errcode"] = 1
		ret["msg"] = "登录失败,账号或者密码错误"
		ctx.JSON(ret)
		return

	}
	ret["errcode"] = 0
	ret["msg"] = "登录成功"
	ret["username"] = user.Username
	ctx.JSON(ret)

}
