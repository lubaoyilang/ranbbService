package main

import (
	_ "ranbbService/routers"
	"github.com/astaxie/beego"
	"ranbbService/cmd"
)

func init() {

}

func main() {

	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.SessionOn = true
		beego.SessionCookieLifeTime = 999999
		beego.StaticDir["/ranbaobao/views"] = "views"
	}
	cmd.Run()
	beego.Run()
}
