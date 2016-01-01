package main

import (
	_ "ranbbService/routers"
	"github.com/astaxie/beego"
	"ranbbService/cmd"
)

func init() {
	beego.SetLogger("file", `{"filename":"logs/daily.log"}`)
}

func main() {

	if beego.RunMode == "dev" {

	}else{
		beego.BeeLogger.DelLogger("console")
	}
	beego.DirectoryIndex = true
	beego.SessionOn = true
	beego.SessionCookieLifeTime = 999999
	beego.StaticDir["/ranbaobao/views"] = "views"
	beego.StaticDir["/ranbaobao/docs"] = "views/apiviews"
	cmd.Run()
	beego.Run()
}


