package controllers
import (
	"ranbbService/util"
	"ranbbService/models"
	"github.com/bitly/go-simplejson"
	"github.com/astaxie/beego"
	"ranbbService/session"
	"strings"
	"time"
)


type walletLogReq struct {
	Page int
	Size int
	Mode int
}

type OutToAlipay struct {
	Amount int64
	PassWord string
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

func (this *RanBaobaoController)OutToAlipayAcc(req *RanBaoBaoRequest, rsp *RanBaoBaoResponse) {
	if !this.validitySession(req.SID) {
		rsp.RC = RC_ERR_1012
		return
	}
	var pl OutToAlipay
	err := util.ConvertToModel(&req.PL,&pl)
	if err != nil {
		beego.Error(err.Error())
		rsp.RC = RC_ERR_1001
		return
	}

	user := &models.User{UID:session.GetSessionByiD(req.SID)}
	err = models.GetUser(user)
	if err != nil {
		rsp.RC = RC_ERR_1010
		return
	}

	if pl.Amount > user.Asset+user.Income+user.Rate {
		rsp.RC = RC_ERR_1035
		return
	}
	if !strings.EqualFold(util.StringMd5(pl.PassWord),user.PassWord) {
		rsp.RC = RC_ERR_1011
		return
	}
	
	en := &models.Enchashment{UID:user.UID,Amount:pl.Amount,CreateTime:time.Now().Unix(),UpdateTime:time.Now().Unix(),State:0,Memo:"申请提现"}
	err = models.AddEnchashment(en)
	if err != nil {
		beego.Error(err.Error())
		rsp.RC = RC_ERR_1036
		return
	}
	return;
}