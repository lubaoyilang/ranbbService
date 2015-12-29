package models
import "github.com/go-xweb/log"


var (
	_USER_TABLE_NAME = "user"
)

type User struct {
	UID string `xorm:"'UID' pk notnull unique varchar(35)"`
	Mobile string `xorm:"'mobile' index notnull unique varchar(15)"`
	PassWord string `xorm:"'password' not null -> varchar(35)" json:"-"`
	RealName string `xorm:"'realName' notnull varchar(35)"`
	IdCard string `xorm:"'idCard' notnull unique index varchar(20)"`
	AliPayAccount string `xorm:"'aliPayAccount' index notnull unique varchar(35)"`
	AliPayName string `xorm:"'aliPayName' notnull varchar(35)"`
	Active int `xorm:"'active' tinyint(1) default 1 "`
	Asset int64 `xorm:"'asset' BigInt(10) default 0"`
	Rate int64 `xorm:"'rate' BigInt(10) default 0"`
	Income int64 `xorm:"'income' BigInt(10) default 0"`
	Total int64 `xorm:"'total' BigInt(10) default 0"`
	CreateTime int64 `xorm:"'createTime' BigInt(10) default 0"`
	UpdateTime int64 `xorm:"'updateTime' BigInt(10) default 0"`
}


func (user * User) TableName() string {
	return _USER_TABLE_NAME
}

func AddUser(user * User) {
	session := Engine.NewSession()
	defer session.Close()
	session.Begin()
	id,err := session.Insert(user)
	if err != nil {
		log.Info(id,err.Error())
		session.Rollback()
	}
	return
}