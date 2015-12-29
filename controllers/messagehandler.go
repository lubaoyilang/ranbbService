package controllers
import (
	"github.com/astaxie/beego"
	"encoding/json"
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

func (this * RanBaobaoController)Handler() {
	var req RanBaoBaoRequest
	this.BindJson(&req)

	rsp := RanBaoBaoResponse{CID:req.CID+1,RC:RC_OK}

	switch req.CID {
	case CID_CAPTCHA_REQ:
		GetCaptCha(&req,&rsp)
	case CID_REGISTER_REQ:
		Register(&req,&rsp)
	}
	this.Data["json"] = rsp
	this.ServeJson()
}


func (this * RanBaobaoController) BindJson(v interface{}) {
	err := json.Unmarshal(this.Ctx.Input.RequestBody,v)
	if err != nil {
		panic(err.Error())
	}
}