
package routers

import (
	"ranbbService/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/ranbb/api",&controllers.RanBaobaoController{},"post:Handler");
	beego.Router("/ranbb/upload/?:sid/?:type",&controllers.RanBaobaoController{},"post:UpLoadFile");
}
