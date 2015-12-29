package util
import (
	caprand "math/rand"
	"strconv"
	"crypto/md5"
	"io"
	"encoding/hex"
	"encoding/base64"
	"crypto/rand"
	"encoding/json"
)


func BuildCaptcha() string {
	v := caprand.Intn(899999)
	return strconv.Itoa(v+100000)
}

func StringMd5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}

func GetGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return StringMd5(base64.URLEncoding.EncodeToString(b))
}

func ConvertToModel(src interface{},dst interface{}) (err error) {
	data,err := json.Marshal(src)
	if err != nil {
		return
	}
	err = json.Unmarshal(data,dst)
	if err != nil {
		return
	}
	return
}