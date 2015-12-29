package cmd
import (
	"github.com/go-xweb/log"
	"fmt"
	"strings"
	"ranbbService/models"
)

var (

	_HELP = "help"
	_SYNCDB = "SyncDB"
	_PASSWD = "demon"
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
