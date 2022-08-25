package main

import (
	"IrisBlog/handler"
	"github.com/kataras/iris/v12"
)

func main() {

	app := newApp()

	app.HandleDir("/assets", iris.Dir("./assets"))
	app.Favicon("./favicon.ico")
	app.Listen(":5000")
}

func newApp() *iris.Application {

	app := iris.New()

	tmpl := iris.HTML("./views", ".html")

	tmpl.Delims("${", "}")

	tmpl.Reload(true)

	app.RegisterView(tmpl)

	adminhandler := app.Party("/admin")
	{
		adminhandler.Use(iris.Compression)
		adminhandler.Get("/user/", handler.Admin_user_page)
		adminhandler.Get("/userlist/", handler.Admin_userlist)
		adminhandler.Delete("/user_action/", handler.Admin_userdel)
		adminhandler.Put("/user_action/", handler.Admin_userupdate)
		adminhandler.Post("/user_action/", handler.Admin_useradd)

	}

	app.Get("/", func(ctx iris.Context) {

		ctx.ViewData("message", "你好，女神")

		ctx.View("index.html")
	})

	return app

}
