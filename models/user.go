package models
import (
	"errors"
	"github.com/astaxie/beego"
)


var (
	_USER_TABLE_NAME = "user"
)

type User struct {
	UID string `xorm:"'UID' pk notnull unique varchar(35)"`
	Mobile string `xorm:"'mobile' index notnull unique varchar(15)"`
	PassWord string `xorm:"'password' notnull varchar(35)" json:"-"`
	RealName string `xorm:"'realName' notnull varchar(35)"`
	IdCard string `xorm:"'idCard' notnull unique index varchar(20)"`
	AliPayAccount string `xorm:"'aliPayAccount' index notnull unique varchar(35)"`
	AliPayName string `xorm:"'aliPayName' notnull varchar(35)"`
	UserImage string  `xorm:"userImage TEXT"`
	Active int `xorm:"'active' tinyint(1) default 1 "`
	Asset int64 `xorm:"'asset' BigInt(10) default 0"`
	Rate int64 `xorm:"'rate' BigInt(10) default 0"`
	VerifyAmount int64 `xorm:"verifyAmount BigInt(10) default 0"`
	Income int64 `xorm:"'income' BigInt(10) default 0"`
	Total int64 `xorm:"'total' BigInt(10) default 0"`
	CreateTime int64 `xorm:"'createTime' BigInt(10) default 0"`
	UpdateTime int64 `xorm:"'updateTime' BigInt(10) default 0"`
}


func (user * User) TableName() string {
	return _USER_TABLE_NAME
}

func AddUser(user * User) error {
	session := Engine.NewSession()
	defer session.Close()
	session.Begin()
	id,err := session.Insert(user)
	if err != nil {
		beego.Info(id,err.Error())
		session.Rollback()
		return err
	}
	return nil
}

func GetUser(user * User) error {
	session := Engine.NewSession()
	defer session.Close()
	session.Begin()
	has,err := session.Get(user)
	if err != nil {
		beego.Info(err.Error())
		return err
	}
	if !has {
		beego.Info("xorm can't find user ",user)
		return errors.New("xorm can't find user")
	}
	return nil
}

func SetUserPasswd(user *User) error{
	session := Engine.NewSession()
	defer session.Close()
	session.Begin()
	_,err := session.Exec(`update user set password = ? where UID = ?`,user.PassWord,user.UID)
	if err != nil {
		beego.Debug("update ok",err.Error());
		session.Rollback()
		return err;
	}
	beego.Debug("update ok");
	return session.Commit()
}

func UpdateUserInfo(user *User) error {
	session := Engine.NewSession()
	defer session.Close()
	session.Begin()
	_,err := session.Exec(`update user set realName = ?,idCard=?,aliPayAccount=?,mobile=? where UID = ?`,user.RealName,user.IdCard,user.AliPayAccount,user.Mobile,user.UID)
	if err != nil {
		beego.Debug("update ok",err.Error());
		session.Rollback()
		return err;
	}
	beego.Debug("update ok");
	return session.Commit()
}

func UpdateUserImage(user * User) error {
	session := Engine.NewSession()
	defer session.Close()
	session.Begin()
	_,err := session.Exec(`update user set userImage = ? where UID = ?`,user.UserImage,user.UID)
	if err != nil {
		beego.Debug("update ok",err.Error());
		session.Rollback()
		return err;
	}
	beego.Debug("update ok");
	return session.Commit()
}

func UpdateUserVertifyAmount(user *User)error{
	session := Engine.NewSession()
	defer session.Close()
	session.Begin()
	_,err := session.Exec(`update user set verifyAmount = ?,total=? where UID = ?`,user.VerifyAmount,user.Total,user.UID)
	if err != nil {
		beego.Debug("update ok",err.Error());
		session.Rollback()
		return err;
	}
	beego.Debug("update ok");
	return session.Commit()
}