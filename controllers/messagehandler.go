package controllers
import (
	"github.com/astaxie/beego"
	"encoding/json"
	_ "github.com/astaxie/beego/session/mysql"
	"github.com/astaxie/beego/cache"
)


var (
	servcache  cache.Cache
	cacheTimeout int64
)

type RanBaobaoController struct {
	beego.Controller
}


type RanBaoBaoRequest struct {
	CID int
	SID string
	PL interface{}
}

type RanBaoBaoResponse struct {
	CID int
	RC int
	PL interface{}
}

func init() {
	beego.SessionProvider = "mysql" //"memory"//
	beego.SessionSavePath = "cloudbridge:Cbcnspsp06@tcp(115.29.164.59:3306)/storedb?charset=utf8"
//	beego.SessionProvider = "file"
//	beego.SessionSavePath = "./tmp"
	beego.AppConfig.Int("captchaTimeout")
	t,_ := beego.AppConfig.Int("captchaTimeout")
	cacheTimeout = int64(t)
	servcache,_ = cache.NewCache("memory",`{"interval":1}`)
}

func (this * RanBaobaoController)Handler() {

	var req RanBaoBaoRequest
	this.bindJson(&req)

	rsp := RanBaoBaoResponse{CID:req.CID+1,RC:RC_OK}
	switch req.CID {
	case CID_CAPTCHA_REQ:
		this.GetCaptCha(&req,&rsp)
	case CID_REGISTER_REQ:
		this.Register(&req,&rsp)
	case CID_LOGIN_REQ:
		this.Login(&req,&rsp)
	case CID_LOGOUT_REQ:
		this.Logout(&req,&rsp)
	case CID_USER_INFO_REQ:
		this.GetUserInfo(&req,&rsp)
	default:
		rsp.RC = RC_ERR_1000
	}
	this.Data["json"] = rsp
	this.ServeJson()
}


func (this * RanBaobaoController) bindJson(v interface{}) {
	err := json.Unmarshal(this.Ctx.Input.RequestBody,v)
	if err != nil {
		panic(err.Error())
	}
}

func (this * RanBaobaoController) validitySession(sid string) bool  {
	if this.GetSession(sid) == nil {
		return  false
	}
	return true
}
