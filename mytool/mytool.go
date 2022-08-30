package mytool

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"

	"github.com/dchest/captcha"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

const (
	StdWidth  = 80
	StdHeight = 40
)

var SigKey = []byte("signature_hmac_secret_shared_key")

type PlayLoad struct {
	Uid uint
}

func GenerateToken(uid uint) string {

	signer := jwt.NewSigner(jwt.HS256, SigKey, 50*time.Minute)
	claims := PlayLoad{Uid: uid}

	token, err := signer.Sign(claims)
	if err != nil {
		fmt.Println(err)

	}

	s := string(token)

	return s

}

func TestToken() iris.Handler {
	return func(ctx iris.Context) {
		signer := jwt.NewSigner(jwt.HS256, SigKey, 50*time.Minute)
		claims := PlayLoad{Uid: 1}
		token, err := signer.Sign(claims)
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}

		ctx.Write(token)
	}
}

func IenerateToken(signer *jwt.Signer) iris.Handler {
	return func(ctx iris.Context) {
		claims := PlayLoad{Uid: 1}

		token, err := signer.Sign(claims)
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}

		ctx.Write(token)
	}
}

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

func Logout(ctx iris.Context) {
	err := ctx.Logout()
	if err != nil {
		ctx.WriteString(err.Error())
	} else {
		ret := make(map[string]interface{}, 0)
		ret["errcode"] = 1
		ret["msg"] = "请您重新登录"
		ctx.JSON(ret)
	}
}
