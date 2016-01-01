package controllers
import (
	"ranbbService/models"
	"github.com/bitly/go-simplejson"
	"ranbbService/util"
	"time"
	"ranbbService/session"
)


type updateTaobaoAccReq struct {
	Tid int
	TaoBaoAccount string
	Memo string
}

func (this * RanBaobaoController) GetTaobaoAccList(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse){
	if !this.validitySession(req.SID) {
		rsp.RC = RC_ERR_1012
		return
	}

	accs,count,err := models.GetTaobaoAccList(session.GetSessionByiD(req.SID))
	if err != nil {
		rsp.RC = RC_ERR_1026
		return
	}

	js := simplejson.New()
	js.Set("Count",count)
	js.Set("Data",accs)

	rsp.PL = accs
}

func (this * RanBaobaoController) UpdateTaobaoAcc(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse){
	if !this.validitySession(req.SID) {
		rsp.RC = RC_ERR_1012
		return
	}


	var pl updateTaobaoAccReq
	err := util.ConvertToModel(&req.PL,&pl)
	if err != nil {
		rsp.RC = RC_ERR_1001
		return
	}


	if pl.Tid <= 0 {
		rsp.RC = RC_ERR_1001
		return
	}

	acc := &models.TaobaoAccount{Tid:pl.Tid,TaoBaoAccount:pl.TaoBaoAccount,UpdateTime:time.Now().Unix()}
	err = models.UpdateTaobaoAcc(acc)
	if err != nil {
		rsp.RC = RC_ERR_1027
		return
	}

	return
}

func (this * RanBaobaoController) DeleteTaobaoAcc(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse){
	if !this.validitySession(req.SID) {
		rsp.RC = RC_ERR_1012
		return
	}


	var pl updateTaobaoAccReq
	err := util.ConvertToModel(&req.PL,&pl)
	if err != nil {
		rsp.RC = RC_ERR_1001
		return
	}


	if pl.Tid <= 0 {
		rsp.RC = RC_ERR_1001
		return
	}

	err = models.DeleteTaobaoAcc(pl.Tid)
	if err != nil {
		rsp.RC = RC_ERR_1028
		return
	}
	return
}

func (this * RanBaobaoController) AddTaobaoAcc(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse){
	if !this.validitySession(req.SID) {
		rsp.RC = RC_ERR_1012
		return
	}


	var pl updateTaobaoAccReq
	err := util.ConvertToModel(&req.PL,&pl)
	if err != nil {
		rsp.RC = RC_ERR_1001
		return
	}

	acc := &models.TaobaoAccount{
		UID:session.GetSessionByiD(req.SID),
		TaoBaoAccount:pl.TaoBaoAccount,
		Memo:pl.Memo,
		CreateTime:time.Now().Unix(),
		UpdateTime:time.Now().Unix()}
	err = models.AddTaobaoAcc(acc)
	if err != nil {
		rsp.RC = RC_ERR_1029
		return
	}
	return
}
