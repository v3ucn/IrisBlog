package handler

import (
	"IrisBlog/database"
	"IrisBlog/model"
	"IrisBlog/mytool"
	"fmt"

	"github.com/kataras/iris/v12"
)

//用户管理页面模板
func Admin_user_page(ctx iris.Context) {

	ctx.View("/admin/user.html")

}

//用户列表接口
func Admin_userlist(ctx iris.Context) {

	db := database.Db()

	var users []model.User
	res := db.Find(&users)
	// 逻辑结束后关闭数据库
	defer func() {
		_ = db.Close()
	}()

	ctx.JSON(res)

}

//用户删除
func Admin_userdel(ctx iris.Context) {

	db := database.Db()

	ID := ctx.URLParamIntDefault("id", 0)

	db.Delete(&model.User{}, ID)

	defer func() {
		_ = db.Close()
	}()

	ret := map[string]string{
		"errcode": "0",
		"msg":     "删除用户成功",
	}
	ctx.JSON(ret)

}

//用户更新
func Admin_userupdate(ctx iris.Context) {

	db := database.Db()

	ID := ctx.PostValue("id")
	Password := ctx.PostValue("password")

	user := &model.User{}
	db.First(&user, ID)

	user.Password = mytool.Make_password(Password)
	db.Save(&user)

	defer func() {
		_ = db.Close()
	}()

	ret := map[string]string{
		"errcode": "0",
		"msg":     "更新密码成功",
	}
	ctx.JSON(ret)

}

//用户添加
func Admin_useradd(ctx iris.Context) {

	db := database.Db()

	defer func() {
		_ = db.Close()
	}()

	username := ctx.PostValue("username")
	password := ctx.PostValue("password")

	fmt.Println(username, password)

	md5str := mytool.Make_password(password)

	user := &model.User{Username: username, Password: md5str}
	res := db.Create(user)

	if res.Error != nil {

		fmt.Println(res.Error)

		ret := map[string]string{
			"errcode": "1",
			"msg":     "用户名不能重复",
		}
		ctx.JSON(ret)

		return

	}

	ret := map[string]string{
		"errcode": "0",
		"msg":     "ok",
	}
	ctx.JSON(ret)

}
