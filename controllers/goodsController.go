package controllers
import (
	"ranbbService/models"
	"github.com/bitly/go-simplejson"
	"ranbbService/util"
)

type goodsReq struct {
	ShopId string
	Page int
	Size int
}

func (this * RanBaobaoController) GetGoodsByShop(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse)  {
	if !this.validitySession(req.SID) {
		rsp.RC = RC_ERR_1012
		return
	}

	var pl goodsReq
	err := util.ConvertToModel(&req.PL,&pl)
	if err != nil {
		beego.Error(err.Error())
		rsp.RC = RC_ERR_1001
		return
	}

	goods,count ,err := models.GetShopGoodsByPage(pl.Page,pl.Size,pl.ShopId)
	if err != nil {
		rsp.RC = RC_ERR_1014 //获取列表失败
		return
	}

	json := simplejson.New()
	json.Set("Count",count)
	json.Set("data",goods)
	rsp.PL = json
}