package lib

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"strconv"
	"time"
)

//create md5 string
func Strtomd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

func Pwdhash(str string) string {
	return Strtomd5(str)
}

func StringsToJson(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}

	return jsons
}

//目前采用的系统路径采用天为单位分割
func HashDatePath(id string) string {
	path := time.Now().Format("20060102")
	path = "mm" + string(os.PathSeparator) + path + string(os.PathSeparator)
	os.MkdirAll(path, 0777)
	path = path + id
	return path
}

//返回系统当前时间16进制
func HashDateName() string {
	name := strconv.FormatInt(time.Now().UnixNano(), 16)
	return name
}
