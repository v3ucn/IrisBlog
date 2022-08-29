package mytool

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/dchest/captcha"
	"github.com/kataras/iris/v12"
)

const (
	StdWidth  = 80
	StdHeight = 40
)

func GetCaptchaId(ctx iris.Context) {
	m := make(map[string]interface{}, 0)
	m["errcode"] = 0
	m["msg"] = "获取成功"
	m["captchaId"] = captcha.NewLen(4)
	ctx.JSON(m)
	return
}

func GetCaptchaImg(ctx iris.Context) {
	captcha.Server(StdWidth, StdHeight).
		ServeHTTP(ctx.ResponseWriter(), ctx.Request())
}

func Make_password(password string) string {

	w := md5.New()
	io.WriteString(w, password) //将str写入到w中
	md5str := fmt.Sprintf("%x", w.Sum(nil))

	return md5str

}
