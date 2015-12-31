package models

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
	State int `xorm:"'state' tinyint(1)  default 0"`
	Price int64 `xorm:"price "`
	RequireLevel int `xorm:"requireLevel tinyint(2) "`
	ShopRequire string `xorm:"shopRequire blob"`
	ImageUrl string `xorm:"imageUrl blob"`
	BrokerAge int64 `xorm:"brokerAge"`
	CreateTime int64 `xorm:"createTime default 0"`
	UpdateTime int64 `xorm:"updateTime default 0"`
	Quantity int `xorm:"quantity"`
	LimitPurchaseQuantity int `xorm:"limitPurchaseQuantity default 1"`
	Memo string `xorm:"memo blob"`
}

func GetShopGoodsByPage(page,size ,shopId int ) ( *[]Goods ,int, error) {
	goods := make([]Goods, 0)
	if page <= 0{
		page = 0
	}else {
		page -= 1
	}

	page = page*size

	err := Engine.Table(new(Goods)).Where("shopId = ?",shopId).Limit(size,page).Find(&goods)
	return &goods,len(goods),err
}