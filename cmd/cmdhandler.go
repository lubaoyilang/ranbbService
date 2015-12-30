package cmd
import (
	"fmt"
	"strings"
	"ranbbService/models"
	"io/ioutil"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/cache"
	"time"
	"github.com/astaxie/beego"
)

const (
	_HELP = "help"
	_SYNCDB = "SyncDB"
	_PASSWD = "demon"
	_SMS = "sms"
	_CACHE = "cache"
)

var c cache.Cache

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
		case _CACHE:
			testCache()
		default:
			fmt.Println("unknown cmd")
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
		fmt.Println("sync database .....")
		models.Sync2()
		fmt.Println("sync databse success,please restart service!")
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
		beego.Informational(err.Error())
	}
	data,err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		beego.Informational(err.Error())
	}
	fmt.Println(string(data))
}

func testCache()  {
	var input string
	fmt.Print("input some thing want to cache")
	fmt.Scanln(&input)
	c,_ = cache.NewCache("memory",`{"interval":1}`)
	c.Put("key",input,1);
	time.Sleep(time.Millisecond * 500)
	output := c.Get("key")
	fmt.Println(output)
	time.Sleep(time.Millisecond * 499)
	output = c.Get("key")
	fmt.Println(output)
}