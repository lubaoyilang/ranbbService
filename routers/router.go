
package routers

import (
	"ranbbService/controllers"
	"github.com/astaxie/beego"
	"github.com/go-xweb/log"
)

func init() {
	log.Info("beego router is init...")
	beego.Router("/ranbb/api",&controllers.RanBaobaoController{},"post:Handler");
	log.Info("beego router is init ok")
}
