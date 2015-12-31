package controllers
import (
	"ranbbService/util"
	"ranbbService/models"
	"time"
	"strings"
	"github.com/bitly/go-simplejson"
	"github.com/astaxie/beego"
	"ranbbService/msg"
	"fmt"
)


type captchaPl struct {
	Mobile string
}

/**

{
  "PL": {
    "Mobile": "18601761051",
    "PassWord": "abcd1234",
    "RealName": "王猛",
    "IdCard": "412723156892140833",
    "AliyPayId": "89864623@qq.com",
    "AliyPayName":"demon",
    "Captcha": "456123"
  },
  "CID": "91011"
}
 */

type registerPl struct {
	Mobile string
	PassWord string
	RealName string
	IdCard string
	AliyPayId string
	AliyPayName string
	Captcha string
}

type loginPl struct  {
	Mobile string
	PassWord string
}

func (this * RanBaobaoController)GetCaptCha(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse) {
	pl := captchaPl{}

	err := util.ConvertToModel(req.PL,&pl)
	if err != nil {
		beego.Error(err.Error())
		rsp.RC = RC_ERR_1001
		return
	}

	if !util.IsPhone(pl.Mobile) {
		beego.Error("手机号码输入有误!")
		rsp.RC = RC_ERR_1002
		return
	}
	captcha := util.BuildCaptcha()

	if(beego.RunMode != "dev"){
		smsContent := fmt.Sprintf("您的验证码为%s \n",captcha)
		err = msg.SendSms(pl.Mobile,smsContent)
		if err != nil {
			beego.Info("发送失败")
			rsp.RC = RC_ERR_1003
			return
		}

	}
	rsp.PL = captcha
	rsp.RC = RC_OK

	servcache.Put(pl.Mobile,captcha,cacheTimeout)
	return
}

func (this * RanBaobaoController) Register(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse) {

	pl := registerPl{}
	err := util.ConvertToModel(&req.PL,&pl)
	if err != nil {
		beego.Error(err.Error())
		rsp.RC = RC_ERR_1001
		return
	}

	//验证码
	captcha := servcache.Get(pl.Mobile)
	if captcha == nil ||!strings.EqualFold(captcha.(string),pl.Captcha) {
		beego.Error("验证码错误:%s != %s",captcha,pl.Captcha)
		rsp.RC = RC_ERR_1004
		return
	}

	if !util.ValidityIdCard(pl.RealName,pl.IdCard) {
		beego.Error("错误的身份证:%s != %s",pl.RealName,pl.IdCard)
		rsp.RC = RC_ERR_1005
		return
	}
	//

	user := &models.User{
		UID:util.GetGuid(),
		Mobile:pl.Mobile,
		PassWord:util.StringMd5(pl.PassWord),
		RealName:pl.RealName,
		IdCard:pl.IdCard,
		AliPayAccount:pl.AliyPayId,
		AliPayName:pl.AliyPayName,
		CreateTime:time.Now().Unix(),
		UpdateTime:time.Now().Unix()}
	err = models.AddUser(user)
	if err != nil {
		errStr := err.Error()
		if strings.Contains(errStr,"UQE_user_mobile") {
			beego.Info("手机号重复注册")
			rsp.RC = RC_ERR_1007 // = 1007 //手机号已经被注册
		}else if strings.Contains(errStr,"UQE_user_idCard") {
			beego.Info("身份证重复")
			rsp.RC = RC_ERR_1008
		}else if strings.Contains(errStr,"UQE_user_aliPayAccount"){
			beego.Info("支付宝账号重复")
			rsp.RC = RC_ERR_1010
		}else{
			beego.Error("创建用户失败"+err.Error())
			rsp.RC = RC_ERR_1006
		}
		return
	}
	json := simplejson.New()
	json.Set("UID",user.UID)
	rsp.PL = json
	return
}

func (this * RanBaobaoController) Login(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse){
	pl := loginPl{}
	err := util.ConvertToModel(&req.PL,&pl)
	if err != nil {
		beego.Error(err.Error())
		rsp.RC = RC_ERR_1001
		return
	}

	user := models.User{Mobile:pl.Mobile}
	err = models.GetUser(&user)
	if err != nil {
		rsp.RC = RC_ERR_1010
		return
	}

	if !strings.EqualFold(util.StringMd5(pl.PassWord),user.PassWord) {
		rsp.RC = RC_ERR_1011
		return
	}


	SID := util.GetGuid()

	session := this.GetSession(user.UID)
	if  session != nil {
		this.DelSession(session)
	}

	this.SetSession(SID,user.UID)
	this.SetSession(user.UID,SID)

	json := simplejson.New()
	json.Set("SID",SID)
	rsp.PL = json
	return
}

func (this * RanBaobaoController) Logout(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse){
	this.DelSession(this.GetSession(req.SID))
	this.DelSession(req.SID)
	return
}

func (this * RanBaobaoController) GetUserInfo(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse){

	if !this.validitySession(req.SID)  {
		rsp.RC = RC_ERR_1012
		return
	}

	uid := this.GetSession(req.SID)

	if uid == nil {
		rsp.RC = RC_ERR_1012
		return
	}

	user := &models.User{UID:uid.(string)}
	err := models.GetUser(user)
	if err != nil {
		rsp.RC = RC_ERR_1010
		return
	}
	rsp.PL = user
	return
}