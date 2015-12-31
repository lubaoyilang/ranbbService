package models


/*
Tid	int	主键
TaoBaoAccount	String	刷单所用淘宝账号
CreateTime	long	创建时间
UpdateTime	long	更新时间
Memo	String	描述
 */

type TaobaoAccount struct {
	Tid int `xorm:"tid pk int(11) unique autoincr "`
	UID string `xorm:"'UID' index notnull varchar(35)"`
	TaoBaoAccount string  `xorm:"'taobaoAccount' index notnull  varchar(35)"`
	CreateTime int64 `xorm:"createTime default 0"`
	UpdateTime int64 `xorm:"updateTime default 0"`
	Memo string  `xorm:"memo blob"`
}