package models
import "errors"

/**
 			  "GoodsId":"45514",
              "ShopId":"451445",
              "ShopName":"然宝宝童鞋店",
              "State":0,
              "Price":1000,
              "RequireLevel":3,
              "ShopRequire":"好多要求好多要求就是好多要求",
              "ImageUrl":"http://www.baidu.com/image/img.png",
              "BrokerAge":100,
              "CreateTime":12345677411745,
              "UpdateTime":12345678551212,
              "Quantity":100,
              "LimitPurchaseQuantity":10,
              "Memo":"好商品,好商品"
 */

type Goods struct {
	GoodsId int `xorm:"'goodId' pk int(11) unique autoincr "`
	ShopId int `xorm:"shopId index int(11) notnull"`
	ShopName string `xorm:"shopName varcha(32)"`
	State int `xorm:"'state' tinyint(1)  default 1"` //0 : 下架 1:正常
	RequireLevel int `xorm:"requireLevel tinyint(2) "`
	ShopRequire string `xorm:"shopRequire TEXT"`
	ImageUrl string `xorm:"imageUrl TEXT"`
	BrokerAge int64 `xorm:"brokerAge"`
	CreateTime int64 `xorm:"createTime default 0"`
	UpdateTime int64 `xorm:"updateTime default 0"`
	Memo string `xorm:"memo TEXT"`
}

func GetGoodsOfShopByPage(page,size ,shopId int ) ( *[]Goods ,int, error) {
	goods := make([]Goods, 0)
	if page <= 0{
		page = 0
	}else {
		page -= 1
	}

	if size <= 0 {
		 size = 20
	}

	page = page*size

	if shopId > 0 {
		err := Engine.Table(new(Goods)).Where("shopId = ?",shopId).OrderBy("createTime").Limit(size,page).Find(&goods)
		return &goods,len(goods),err
	}else{
		err := Engine.Table(new(Goods)).OrderBy("createTime").Limit(size,page).Find(&goods)
		return &goods,len(goods),err
	}

}

func GetGoods(goodsId int) (* Goods,error) {
	goods := Goods{GoodsId:goodsId}
	has,err := Engine.Get(&goods)
	if err != nil || !has{
		return nil,errors.New("can not find this goods")
	}
	return &goods,nil
}

func AddGoods(goods * Goods) error {
	sess := Engine.NewSession()
	defer sess.Close()
	sess.Begin()
	num,err := sess.InsertOne(goods)
	if err != nil {
		sess.Rollback()
		return err
	}
	if num <= 0 {
		return errors.New("添加失败")
	}
	return nil
}

//func SubOneGoods(goods * Goods) error {
//	sess := Engine.NewSession()
//	defer sess.Close()
//	sess.Begin()
//	goods.Quantity -= 1
//	_,err := sess.Where("goodId = ?",goods.GoodsId).Update(goods)
//	if err != nil {
//		sess.Rollback()
//		return err
//	}
//	return nil
//}