package cmd
import (
	"github.com/go-xweb/log"
	"fmt"
	"strings"
	"ranbbService/models"
	"io/ioutil"
	"github.com/astaxie/beego/httplib"
)

var (

	_HELP = "help"
	_SYNCDB = "SyncDB"
	_PASSWD = "demon"
	_SMS = "sms"
)

func Run() {
	go handleCmd()
}

func handleCmd() {
	var cmd string
	for {
		fmt.Scanln(&cmd)
		switch cmd {
		case _HELP,"?" :
			printUseage()
		case _SYNCDB :
			syncDB()
		case _SMS:
			textSms()
		}
	}
}


func printUseage() {
	fmt.Println("======================================")

	fmt.Println("SyncDB : 同步表结构")

	fmt.Println("======================================")
}

func syncDB() {
	var passwd string
	fmt.Print("输入管理员密码:")
	fmt.Scanln(&passwd)
	if strings.EqualFold(passwd,_PASSWD){
		log.Info("sync database .....")
		models.Sync2()
		log.Info("sync databse success,please restart service!")
	}else{
		fmt.Print("密码错误")
	}
}

func textSms() {
	req := httplib.Get("http://apis.baidu.com/kingtto_media/106sms/106sms")
	req.Header("apikey","e3b793a1d9278be02bf49a0f3f80b68a")
	req.Param("mobile","18601761051")
	req.Param("content","123456 【demon】")
	rsp,err := req.SendOut()
	if err != nil {
		log.Info(err.Error())
	}
	data,err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Info(err.Error())
	}
	log.Info(string(data))
}