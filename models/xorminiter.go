package models
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/astaxie/beego"
	"os"
)


var (
	Engine *xorm.Engine
	Models []interface{}
)

func init() {
	var err error
	Engine,err = xorm.NewEngine("mysql","cloudbridge:Cbcnspsp06@tcp(115.29.164.59:3306)/storedb?charset=utf8")
	if err != nil {
		beego.Info(err.Error())
	}

	if(beego.RunMode == "dev"){
		Engine.ShowDebug = true //则会在控制台打印调试信息；
		Engine.ShowErr = true //则会在控制台打印错误信息；
		Engine.ShowWarn = true
		Engine.ShowSQL = true
	}
	initSqlLog()
}

func initSqlLog() {
	if beego.RunMode == "dev" {
		Engine.Logger = xorm.NewSimpleLogger(os.Stdout)
		return
	}else{
		f, err := os.Create("sql.log")
		if err != nil {
			println(err.Error())
			return
		}
		//	defer f.Close()
		Engine.Logger = xorm.NewSimpleLogger(f)
	}
}

func Sync2() {
	err := Engine.Sync2(new(User),new(Shop),new(WalletLog),new(UserToken),new(Goods),new(Orders),new(TaobaoAccount),new(Admin),new(Session))
	if err != nil {
		beego.Info(err.Error())
	}
}