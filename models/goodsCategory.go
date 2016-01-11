package models
import "github.com/astaxie/beego"



type GoodsCategroy struct {
	CategroyId int `xorm:"'id' pk int(11) unique autoincr "`
	GoodsId int `xorm:"'goodId' index int(11)"`
	ShopId int `xorm:"'shopId' index int(11)"`
	Price int64 `xorm:"price "`
	Name string `xorm:"'name' VARCHAR(36)"`
	EnableTime int64 `xorm:"'enbaleTime' BigInt(10) default 0"`
	TotalNum int `xorm:"'totalNum'  int(11)"`
	OutNum int `xorm:"'outNum'  int(11)"`
	Memo string `xorm:"memo blob"`
	LimitPurchaseQuantity int `xorm:"limitPurchaseQuantity default 1"`
}


func GetCateGroyByGoodsId(goodsId int)([]GoodsCategroy,int,error){
	cs := make([]GoodsCategroy,0)
	sess := Engine.NewSession()
	defer sess.Close()
	err := sess.Where("goodId=?",goodsId).Find(&cs)
	if err != nil {
		return cs,0,err
	}
	return cs,len(cs),nil
}

func GetCateGroyById(cateGroyId int) (* GoodsCategroy,error) {
	sess := Engine.NewSession()
	defer sess.Close()
	categroy := GoodsCategroy{}
	has,err := sess.Where("id=?",cateGroyId).Get(&categroy)
	if err != nil || !has{
		beego.Error(err)
		return nil,err
	}
	return &categroy,nil
}

func AddCateGroyOutCount(categroy *GoodsCategroy) error {
	sess := Engine.NewSession()
	defer sess.Close()
	sess.Begin()
	_,err := sess.Query(`update goods_categroy set outNum=? where id=?`,categroy.OutNum+1,categroy.CategroyId)
	if err != nil {
		beego.Error(err)
		sess.Rollback()
		return err
	}
	return sess.Commit()
}


