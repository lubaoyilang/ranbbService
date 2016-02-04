package controllers
import (
	"os"
	"ranbbService/session"
	"ranbbService/models"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

func (this * RanBaobaoController) UpLoadFile() {
	rsp := RanBaoBaoResponse{RC:1}
	sid := this.Ctx.Input.Param(":sid")
	imgtype := this.Ctx.Input.Param(":type")

	if !this.validitySession(sid) {
		rsp.RC = RC_ERR_1012
		this.Data["json"]=rsp
		this.ServeJson()
		return
	}

	uid := session.GetSessionByiD(sid)
	user := &models.User{UID:uid}
	err := models.GetUser(user)
	if err != nil {
		rsp.RC = RC_ERR_1010
		this.Data["json"]=rsp
		this.ServeJson()
		return
	}



	bb := this.Ctx.Input.CopyBody()
	beego.Debug(len(bb),this.Ctx.Input.RequestBody)
	if err != nil {
		rsp.RC =-1
		this.Data["json"]=rsp
		this.ServeJson()
		return
	}

	filepath := fmt.Sprintf("./views/image/%s-%d.%s",uid,time.Now().Unix(),imgtype)
	f,err:=  os.Create(filepath)
	if err != nil {
		rsp.RC =-2
		this.Data["json"]=rsp
		this.ServeJson()
		return
	}
	f.Write(bb)
	f.Close()

	url := beego.AppConfig.String("imageUrl")
	imgUrl := fmt.Sprintf("%s/%s",url,filepath[2:])
	user.UserImage = imgUrl;

	err = models.UpdateUserImage(user)
	if err != nil {
		rsp.RC =-3
		this.Data["json"]=rsp
		this.ServeJson()
		return
	}
	rsp.PL = imgUrl
	this.Data["json"]=rsp
	this.ServeJson()
	return
}