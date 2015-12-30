package util
import (
	"regexp"
	"github.com/astaxie/beego/httplib"
	"io/ioutil"
	"encoding/json"
	"github.com/go-xweb/log"
)

const (
	regular = "^(0|86|17951)?(13[0-9]|15[012356789]|17[678]|18[0-9]|14[57])[0-9]{8}$"
)

/*
{
  "error_code": 0,
"reason": "Succes",
"result":
{
"sex": "M",
"birthday": "1987-09-15",
"address": "湖北省鄂州市鄂城区"
}
}
 */

type IdCardResult struct {
	ErrorCode int `json:"error_code"`
	Reason string `json:"reason"`
	Result interface{} `json:"result"`
}

func IsPhone(s string) bool {
	reg := regexp.MustCompile(regular)
	return reg.MatchString(s)
}

func ValidityIdCard(realName,idCard string) bool {
	//http://avatardata.cn/   wm5821090	395618037
	req := httplib.Get("http://api.avatardata.cn/IdCard/LookUp")
	req.Param("key","b7a5cbe7d42444b199aa09f5778dc939")
	req.Param("id",idCard)
	rsp,err := req.SendOut()
	if err != nil {
		log.Error(err.Error())
		return false
	}
	data,err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Error(err.Error())
		return false
	}
	var result IdCardResult
	err = json.Unmarshal(data,&result)
	if err != nil {
		log.Error(err.Error())
		return false
	}

	if result.ErrorCode != 0 {
		log.Info(result.Reason)
		return false
	}

	return true
}