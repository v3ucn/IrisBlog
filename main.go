package main

import (
	"IrisBlog/model"
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

	app.Get("/", func(ctx iris.Context) {

		ctx.ViewData("message", "你好，女神")

		ctx.View("index.html")
	})

	return app

}
