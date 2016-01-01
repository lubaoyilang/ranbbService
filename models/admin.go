package models
import (
	"ranbbService/util"
	"github.com/astaxie/beego"
)

/*
UID	String	管理员id主键
Name	String	管理员名称
Mobile	String	管理员手机号码
PassWord	String	管理员密码
Privilege	String	权限list 逗号分隔
 */

type Admin struct  {
	UID string `xorm:"'UID' pk notnull unique varchar(35)"`
	Name string `xorm:"'name' varcha(32)"`
	Mobile string `xorm:"'mobile' index notnull unique varchar(15)"`
	PassWord string `xorm:"'password' notnull varchar(35)" json:"-"`
	Privilege string `xorm:"'privilege' notnull varchar(100)" json:"-"`
}


func AddAdmin(mobile, password string) {
	admin := &Admin{
		UID:util.GetGuid(),
		Mobile:mobile,
		Name:"admin",
		PassWord:util.StringMd5(password),
		Privilege: "123456,456789,2132546"}
	sess := Engine.NewSession()
	defer sess.Close()
	_,err := sess.InsertOne(admin)
	if err != nil {
		beego.Error(err)
		sess.Rollback()
	}
	sess.Commit()
}