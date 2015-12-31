package models

/*
ShopId	int	商家id 主键
ShopName	String	店铺名称
Mobile	String	商家手机号
Email	String	商家邮箱
ShopTaoBaoId	String	商家旺旺Id
CreateTime	long	创建时间
UpdateTime	long	更新时间
Memo	String	商家描述
 */

type Shop struct {
	ShopId int `xorm:"shopId pk int(11) unique autoincr "`
	ShopName string `xorm:"shopName varcha(32)"`
	Mobile string `xorm:"'mobile' index notnull unique varchar(15)"`
	Email string `xorm:"'email' index notnull unique varchar(25)"`
	ShopTaoBaoId string `xorm:"'shopTaoBaoId' index notnull unique varchar(35)"`
	CreateTime int64 `xorm:"'createTime' BigInt(10) default 0"`
	UpdateTime int64 `xorm:"'updateTime' BigInt(10) default 0"`
	Memo string `xorm:"memo blob"`
}