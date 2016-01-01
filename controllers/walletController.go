package controllers
import (
	"ranbbService/util"
	"ranbbService/models"
	"github.com/bitly/go-simplejson"
	"github.com/astaxie/beego"
	"ranbbService/session"
)


type walletLogReq struct {
	Page int
	Size int
	Mode int
}

func (this * RanBaobaoController) GetWalletLogs(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse){
	if !this.validitySession(req.SID) {
		rsp.RC = RC_ERR_1012
		return
	}
	var pl walletLogReq
	err := util.ConvertToModel(&req.PL,&pl)
	if err != nil {
		beego.Error(err.Error())
		rsp.RC = RC_ERR_1001
		return
	}

	logs,count,err := models.GetWalletLogByMode(session.GetSessionByiD(req.SID),pl.Page,pl.Size,pl.Mode)
	if err != nil {
		beego.Error(err.Error())
		rsp.RC = RC_ERR_1030
		return
	}

	js := simplejson.New()
	js.Set("Count",count)
	js.Set("Data",logs)
	rsp.PL = js
	return
}