package mytool

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Make_password(password string) string {

	w := md5.New()
	io.WriteString(w, password) //将str写入到w中
	md5str := fmt.Sprintf("%x", w.Sum(nil))

	return md5str

}
