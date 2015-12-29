package controllers
import (
	"github.com/go-xweb/log"
	"ranbbService/util"
	"fmt"
	"ranbbService/msg"
	"ranbbService/models"
	"time"
)


type captchaPl struct {
	Mobile string
}

func GetCaptCha(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse) {
	pl := captchaPl{}

	err := util.ConvertToModel(req.PL,&pl)
	if err != nil {
		log.Error(err.Error())
		rsp.RC = RC_ERR_1001
		return
	}

	if !util.IsPhone(pl.Mobile) {
		log.Error("手机号码输入有误!")
		rsp.RC = RC_ERR_1002
		return
	}
	captcha := util.BuildCaptcha()
	smscontent := fmt.Sprintf("您的验证码为%s \n",captcha)

	err = msg.SendSms(pl.Mobile,smscontent)
	if err != nil {
		log.Info("发送失败")
		rsp.RC = RC_ERR_1003
		return
	}
	rsp.RC = RC_OK
	return
}

func Register(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse) {
	user := &models.User{
		UID:util.GetGuid(),
		Mobile:"18601761051",
		PassWord:util.StringMd5("123456"),
		RealName:"demon",
		IdCard:"412723198906060833",
		AliPayAccount:"838532366@qq.com",
		AliPayName:"demon",
		CreateTime:time.Now().Unix(),
		UpdateTime:time.Now().Unix()}
	models.AddUser(user)
}