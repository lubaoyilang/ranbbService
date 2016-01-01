package controllers
import (
	"fmt"
)



func (this * RanBaobaoController) UpLoadFile() {
//	this.SaveToFile("")


	fmt.Println(string(this.Ctx.Input.RequestBody))
	fmt.Println(this.GetString("email"))
	this.SaveToFile("file","mysql.log")
	this.Redirect("/ranbaobao/views",302)
}