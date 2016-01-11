package controllers
import (
	"github.com/astaxie/beego"
	"encoding/json"
	_ "github.com/astaxie/beego/session/mysql"
	"github.com/astaxie/beego/cache"
	"ranbbService/session"
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
	PL interface{} `json:",omitempty"`
}

func init() {
	beego.SessionProvider = "mysql" //"memory"//
	beego.SessionSavePath = "cloudbridge:Cbcnspsp06@tcp(115.29.164.59:3306)/storedb?charset=utf8"
	beego.AppConfig.Int("captchaTimeout")
	t,_ := beego.AppConfig.Int("captchaTimeout")
	cacheTimeout = int64(t)
	servcache,_ = cache.NewCache("memory",`{"interval":1}`)
}

func (this * RanBaobaoController)Handler() {

	var req RanBaoBaoRequest
	beego.Debug(string(this.Ctx.Input.RequestBody))
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
	case CID_FIND_PASSWD_REQ:
		this.FindPasswd(&req,&rsp)
	case CID_CHANGE_PASSWD_REQ:
		this.ChangePswd(&req,&rsp)
	case CID_GET_GOODS_REQ:
		this.GetGoodsList(&req,&rsp)
	case CID_ACCEPT_GOODS_REQ:
		this.AcceptTask(&req,&rsp)
	case CID_GET_GOODSCATEGROY_REQ:
		this.GetGoodsCateGroy(&req,&rsp)
	case CID_COMMIT_ORDER_REQ:
		this.CommitOrder(&req,&rsp)
	case CID_DELETE_ORDER_REQ:
		this.DeleteOrder(&req,&rsp)
	case CID_GET_ORDER_INFO_REQ :
		this.GetOrderInfo(&req,&rsp)
	case CID_GET_ORDER_LIST_REQ:
		this.GetOrderList(&req,&rsp)
	case CID_GET_TBACC_LIST_REQ:
		this.GetTaobaoAccList(&req,&rsp)
	case CID_UPDATE_TBACC_REQ :
		this.UpdateTaobaoAcc(&req,&rsp)
	case CID_DELETE_TBACC_REQ:
		this.DeleteTaobaoAcc(&req,&rsp)
	case CID_ADD_TBACC_REQ:
		this.AddTaobaoAcc(&req,&rsp)
	case CID_GET_WALLET_LOGS_REQ:
		this.GetWalletLogs(&req,&rsp)
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

	content := session.GetSessionByiD(sid)
	if len(content) <= 0 {
			return false
	}

	return true
}
