package models


/**
OrderId	int	订单Id
GoodsId	int	商品id
ShopId	int	店铺Id
TaoBaoAccount	String	完成订单使用饿淘宝账号
State	int	0:正在购买 1:提交审核 2:审核通过
ShopName	String	商家名称
price	long	商品价格
RequireLevel	int	淘宝买家等级
ShopRequire	String	淘宝卖家要求
ImageUrl	String	商品图片
BrokerAge	long	佣金
CreateTime	long	创建时间
UpdateTime	long	更新时间
Quantity	int	购买数量,默认为1
Memo	String	描述
 */
type Orders struct {
	TaskId int `xorm:"taskId pk int(11) unique autoincr "`
	GoodsId int `xorm:"goodId index int(11) notnull"`
	ShopId int	`xorm:"shopId index int(11) notnull"`
	UID string  `xorm:"'UID' index notnull varchar(35)"`
	TaoBaoAccount string `xorm:"'taobaoAccount' index varchar(35)"`
	State int `xorm:"'state' tinyint(1)  default 0"`
	ShopName string `xorm:"shopName varcha(32)"`
	Price int64 `xorm:"price"`
	RequireLevel int `xorm:"requireLevel tinyint(2) "`
	ShopRequire string `xorm:"shopRequire blob"`
	ImageUrl string `xorm:"imageUrl blob"`
	BrokerAge int64 `xorm:"brokerAge"`
	CreateTime int64 `xorm:"createTime default 0"`
	UpdateTime int64 `xorm:"updateTime default 0"`
	Quantity int `xorm:"quantity"`
	Memo string `xorm:"memo blob"`
}


func AddOrderAndSubGoods(order * Orders,goods *Goods) error {
	sess := Engine.NewSession()
	defer sess.Close()
	sess.Begin()
	_,err := Engine.Insert(order)

	if err != nil {
		sess.Rollback()
		return err
	}
	goods.Quantity -= 1
	_,err = sess.Where("goodId = ?",goods.GoodsId).Update(goods)
	if err != nil {
		sess.Rollback()
		return err
	}
	return sess.Commit()
}