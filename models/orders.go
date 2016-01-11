package models
import (
	"github.com/astaxie/beego"
	"errors"
)


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
	OrderId int `xorm:"orderId pk int(11) unique autoincr "`
	GoodsId int `xorm:"goodId index int(11) notnull"`
	CategroyId int `xorm:"'categroyId' int(11)"`
	CategroyName string `xorm:"'categroyName' varchar(32)"`
	ShopId int	`xorm:"shopId index int(11) notnull"`
	UID string  `xorm:"'UID' index notnull varchar(35)"`
	TaoBaoAccount string `xorm:"'taobaoAccount' index varchar(35)"`
	State int `xorm:"'state' tinyint(1)  default 0"`
	ShopName string `xorm:"shopName varchar(32)"`
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


func GetOrderById(orderId int,UID string) (* Orders,error){
	sess := Engine.NewSession()
	defer sess.Close()
	order :=  &Orders{}
	has,err := sess.Where("orderId = ? AND UID = ?",orderId,UID).Get(order)
	if err != nil || !has{
		return nil,errors.New("can not find order by orderid Uid")
	}
	return order,nil
}

func AddOrderAndSubCateGroyOutNum(order * Orders,categroy *GoodsCategroy) error {

	sess := Engine.NewSession()
	defer sess.Close()
	sess.Begin()
	_,err := Engine.Insert(order)

	if err != nil {
		sess.Rollback()
		return err
	}
//	_,err =sess.Exec(`update goods set quantity = ? where goodId = ?`,goods.Quantity-1,goods.GoodsId)
	err = AddCateGroyOutCount(categroy)
	if err != nil {
		beego.Error(err.Error())
		sess.Rollback()
		return err
	}
	return sess.Commit()
}

func SetOrderState(state,orderId int,taoBaoAccount string) error {
	sess := Engine.NewSession()
	defer sess.Close()
	sess.Begin()
	_,err :=sess.Exec(`update orders set state = ? ,taoBaoAccount = ? where orderId = ?`,state,taoBaoAccount,orderId)
	if err != nil {
		sess.Rollback()
		return err
	}
	return sess.Commit()
}

func CustomDeleteOrder (orderId int)error {
	sess := Engine.NewSession()
	defer sess.Close()
	sess.Begin()
	_,err := sess.Exec(`delete from orders where orderId = ?`,orderId)
	if err != nil {
		sess.Rollback()
		return err
	}
	return sess.Commit()
}

func GetOrdersByState(state,page,size int,UID string) ([]Orders,int,error){
	sess := Engine.NewSession()
	defer sess.Close()
	orders := make([]Orders,0)

	if page <= 0{
		page = 0
	}else {
		page -= 1
	}

	if size <= 0 {
		size = 20
	}

	page = page*size

	err := sess.Where("state = ? AND UID = ?",state,UID).OrderBy("updateTime").Limit(size,page).Find(&orders)
	if err != nil {
		return nil,0,err
	}
	return orders,len(orders),nil
}