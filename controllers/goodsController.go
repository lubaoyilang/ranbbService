package controllers
import (
	"ranbbService/models"
	"github.com/bitly/go-simplejson"
	"ranbbService/util"
	"github.com/astaxie/beego"
	"time"
	"ranbbService/session"
)

type goodsReq struct {
	ShopId int
	Page int
	Size int
}

/*
     "GoodsId":"45514",
     "UID":"asdfasdfasdf42121245454",
     "TaoBaoAccId":"82342347987@qq.com",
     "Count":6
 */
type AcceptReq struct {
	GoodsId int
	CategroyId int
	Count int
}

type commitReq struct {
	OrderId int
	TaoBaoAccount string
}
type orderListReq struct {
	State int
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

	goods,count ,err := models.GetGoodsOfShopByPage(pl.Page,pl.Size,pl.ShopId)
	if err != nil {
		rsp.RC = RC_ERR_1014 //获取列表失败
		return
	}

	json := simplejson.New()
	json.Set("Count",count)
	json.Set("data",goods)
	rsp.PL = json
}

func (this * RanBaobaoController) GetGoodsList(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse){
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

	goods,count ,err := models.GetGoodsOfShopByPage(pl.Page,pl.Size,pl.ShopId)
	if err != nil {
		rsp.RC = RC_ERR_1014 //获取列表失败
		return
	}

	json := simplejson.New()
	json.Set("Count",count)
	json.Set("data",goods)
	rsp.PL = json
}

func (this * RanBaobaoController) AcceptTask(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse){
	if !this.validitySession(req.SID) {
		rsp.RC = RC_ERR_1012
		return
	}
	var pl AcceptReq
	err := util.ConvertToModel(&req.PL,&pl)
	if err != nil {
		rsp.RC = RC_ERR_1001
		return
	}

	categroy,err := models.GetCateGroyById(pl.CategroyId)
	if err != nil {
		rsp.RC = RC_ERR_1015
		return
	}

	goods,err := models.GetGoods(categroy.GoodsId)
	if err != nil {
		rsp.RC = RC_ERR_1015
		return
	}
	if goods.State == 0 {
		rsp.RC = RC_ERR_1016
		return
	}

	if categroy.TotalNum <= categroy.OutNum {
		rsp.RC = RC_ERR_1018
		return
	}

	if pl.Count > categroy.LimitPurchaseQuantity {
		rsp.RC = RC_ERR_1019
		return
	}

	order := models.Orders{
		UID:session.GetSessionByiD(req.SID),
		GoodsId:goods.GoodsId,
		CategroyId:categroy.CategroyId,
		CategroyName:categroy.Name,
		ShopId:goods.ShopId,
		State:0,
		ShopName:goods.ShopName,
		Price:categroy.Price,
		RequireLevel:goods.RequireLevel,
		ShopRequire:goods.ShopRequire,
		ImageUrl:goods.ImageUrl,
		BrokerAge:goods.BrokerAge,
		CreateTime:time.Now().Unix(),
		UpdateTime:time.Now().Unix(),
		Quantity:pl.Count,
		Memo:"接了一单"}

	err = models.AddOrderAndSubCateGroyOutNum(&order,categroy)
	if err != nil {
		rsp.RC = RC_ERR_1017
		return
	}

	rsp.PL = order
}


func  (this * RanBaobaoController)GetGoodsCateGroy(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse) {
	if !this.validitySession(req.SID) {
		rsp.RC = RC_ERR_1012
		return
	}
	var pl AcceptReq
	err := util.ConvertToModel(&req.PL,&pl)
	if err != nil {
		rsp.RC = RC_ERR_1001
		return
	}

	categroy,count,err := models.GetCateGroyByGoodsId(pl.GoodsId)
	js := simplejson.New()
	js.Set("Count",count)
	js.Set("Data",categroy)
	if err != nil {
		beego.Error(err)
	}
	rsp.PL = js
}

func (this * RanBaobaoController) CommitOrder(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse){
	if !this.validitySession(req.SID) {
		rsp.RC = RC_ERR_1012
		return
	}
	var pl commitReq
	err := util.ConvertToModel(&req.PL,&pl)
	if err != nil {
		rsp.RC = RC_ERR_1001
		return
	}

	if len(pl.TaoBaoAccount) < 0 {
		rsp.RC = RC_ERR_1001
		return
	}

	order,err := models.GetOrderById(pl.OrderId,session.GetSessionByiD(req.SID))
	if err != nil {
		rsp.RC = RC_ERR_1020
		return
	}
	if order.State != 0 {
		rsp.RC = RC_ERR_1021
		return
	}

	err = models.SetOrderState(1,order.OrderId,pl.TaoBaoAccount)
	if err != nil {
		rsp.RC = RC_ERR_1022
		return
	}
	return

}

func (this * RanBaobaoController) DeleteOrder(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse){
	if !this.validitySession(req.SID) {
		rsp.RC = RC_ERR_1012
		return
	}
	var pl commitReq
	err := util.ConvertToModel(&req.PL,&pl)
	if err != nil {
		rsp.RC = RC_ERR_1001
		return
	}

	order,err := models.GetOrderById(pl.OrderId,session.GetSessionByiD(req.SID))
	if err != nil {
		rsp.RC = RC_ERR_1020
		return
	}
	if order.State != 0 && order.State != 2 {
		rsp.RC = RC_ERR_1023
		return
	}

	err = models.CustomDeleteOrder(order.OrderId)
	if err != nil {
		rsp.RC = RC_ERR_1024
		return
	}
	return
}

func (this * RanBaobaoController) GetOrderInfo(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse){
	if !this.validitySession(req.SID) {
		rsp.RC = RC_ERR_1012
		return
	}
	var pl commitReq
	err := util.ConvertToModel(&req.PL,&pl)
	if err != nil {
		rsp.RC = RC_ERR_1001
		return
	}

	order,err := models.GetOrderById(pl.OrderId,session.GetSessionByiD(req.SID))
	if err != nil {
		rsp.RC = RC_ERR_1020
		return
	}
	rsp.PL = order
	return
}

func (this * RanBaobaoController) GetOrderList(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse){
	if !this.validitySession(req.SID) {
		rsp.RC = RC_ERR_1012
		return
	}
	
	var pl orderListReq
	err := util.ConvertToModel(&req.PL,&pl)
	if err != nil {
		rsp.RC = RC_ERR_1001
		return
	}

	orders,count,err := models.GetOrdersByState(pl.State,pl.Page,pl.Size,session.GetSessionByiD(req.SID))
	if err != nil {
		rsp.RC = RC_ERR_1025
		return
	}

	js := simplejson.New()
	js.Set("Count",count)
	js.Set("Data",orders)
	rsp.PL = js
	return
}