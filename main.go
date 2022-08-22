package main

import (
	"IrisBlog/model"
	"IrisBlog/mytool"

	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris/v12"
)

func main() {

	db, err := gorm.Open("mysql", "root:root@(localhost)/irisblog?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
		panic("无法连接数据库")
	}
	fmt.Println("连接数据库成功")

	//单数模式
	db.SingularTable(true)

	// 创建默认表
	db.AutoMigrate(&model.User{})

	// 逻辑结束后关闭数据库
	defer func() {
		_ = db.Close()
	}()

	app := newApp(db)

	app.HandleDir("/assets", iris.Dir("./assets"))
	app.Favicon("./favicon.ico")
	app.Listen(":5000")
}

func newApp(db *gorm.DB) *iris.Application {

	app := iris.New()

	tmpl := iris.HTML("./views", ".html")
	// Set custom delimeters.
	tmpl.Delims("${", "}")
	// Enable re-build on local template files changes.
	tmpl.Reload(true)

	app.RegisterView(tmpl)

	app.Get("/admin/user/", func(ctx iris.Context) {

		ctx.View("/admin/user.html")

	})

	app.Get("/admin/userlist/", func(ctx iris.Context) {

		var users []model.User
		res := db.Find(&users)

		ctx.JSON(res)

	})

	app.Post("/admin/user_action/", func(ctx iris.Context) {

		username := ctx.PostValue("username")
		password := ctx.PostValue("password")

		fmt.Println(username, password)

		md5str := mytool.Make_password(password)

		user := &model.User{Username: username, Password: md5str}
		res := db.Create(user)

		fmt.Println(res.Error)

		ret := map[string]string{
			"errcode": "0",
			"msg":     "ok",
		}
		ctx.JSON(ret)

	})

	app.Get("/", func(ctx iris.Context) {

		ctx.ViewData("message", "你好，女神")

		//var user &model.User
		//db.First(&model.User,1)

		password := "123"

		user := &model.User{Username: "888123", Password: password}
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

		fmt.Println(res.Error)
		fmt.Println(user.ID)

		ctx.View("index.html")
	})

	return app

}
