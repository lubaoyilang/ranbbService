package models
import "fmt"

/*
id	int	主键
Amount	long	历史金额
Type	int	0:本金返还,1:利息结算 2:取现 3佣金增加
CreateTime	long	创建时间
Memo	String	发生事件描述

 */

type WalletLog struct {
	Wid int `xorm:"wid pk int(11) unique autoincr "`
	UID string  `xorm:"'UID' index notnull VARCHAR(35)"`
	Amount int64 `xorm:"price"`
	Categroy int `xorm:"'categroy' TINYINT(1)`
	CreateTime int64 `xorm:"createTime default 0"`
	NowValue int64 `xorm:"nowValue"`
	Memo string `xorm:"memo TEXT"`
}

func GetWalletLogByMode(UID string ,page,size,mode int) ([]WalletLog,int,error){
	logs := make([]WalletLog,0)
	w := fmt.Sprintf("UID = '%s' and categroy = %d",UID,mode)

	if page <= 0{
		page = 0
	}else {
		page -= 1
	}

	if size <= 0 {
		size = 20
	}

	page = page*size
	sess := Engine.NewSession()
	defer sess.Close()
	err := sess.Where(w).Limit(size,page).OrderBy("createTime desc").Find(&logs)
	return logs,len(logs),err
}

func AddWalletLog(log * WalletLog) error{
	sess := Engine.NewSession()
	defer sess.Close()
	sess.Begin()
	_,err := sess.InsertOne(log)
	if err != nil {
		 sess.Rollback()
		return err
	}
	return sess.Commit()
}