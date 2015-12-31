package models

/*
id	int	主键
Amount	long	历史金额
Type	int	0:本金返还,1:利息结算 2:取现 3佣金增加
CreateTime	long	创建时间
Memo	String	发生事件描述

 */

type WalletLog struct {
	Wid int `xorm:"wid pk int(11) unique autoincr "`
	UID string  `xorm:"'UID' index notnull varchar(35)"`
	Amount int64 `xorm:"price"`
	Type int `xorm:"'state' tinyint(1)`
	CreateTime int64 `xorm:"createTime default 0"`
	Memo string `xorm:"memo blob"`
}