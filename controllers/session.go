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
	"ranbbService/session"
)


type captchaPl struct {
	Mobile string
}

type changePswdPl struct {
	OldPswd string
	NewPswd string
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

type updateUserInfoPl struct {
	Mobile string
	RealName string
	IdCard string
	AliPayAccount string
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
	js := simplejson.New()
	js.Set("Captcha",captcha)
	rsp.PL = js
	rsp.RC = RC_OK

	servcache.Put(pl.Mobile,captcha,cacheTimeout)
	return
}

func (this * RanBaobaoController) Register(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse) {

	pl := registerPl{}
	err := util.ConvertToModel(&req.PL,&pl)
	if err != nil {
		beego.Error("conver err",err.Error())
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
	servcache.Delete(pl.Mobile)

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


	var SID string

	sess,ok := session.IsExist(user.UID)

	if ok {
		SID = sess.SessionKey
	}else{
		SID = util.GetGuid()
		session.SetSession(SID,user.UID)
	}

	json := simplejson.New()
	json.Set("SID",SID)
	rsp.PL = json
	return
}

func (this * RanBaobaoController) Logout(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse){
	session.DeleteBySid(req.SID)
	return
}

func (this * RanBaobaoController) GetUserInfo(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse){

	if !this.validitySession(req.SID)  {
		rsp.RC = RC_ERR_1012
		return
	}

	uid := session.GetSessionByiD(req.SID)
	user := &models.User{UID:uid}
	err := models.GetUser(user)
	if err != nil {
		rsp.RC = RC_ERR_1010
		return
	}
	rsp.PL = user
	return
}

func (this * RanBaobaoController)FindPasswd(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse) {
	pl := registerPl{}
	err := util.ConvertToModel(&req.PL,&pl)
	if err != nil {
		beego.Error("conver err",err.Error())
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
	servcache.Delete(pl.Mobile);

	if !util.ValidityIdCard(pl.RealName,pl.IdCard) {
		beego.Error("错误的身份证:%s != %s",pl.RealName,pl.IdCard)
		rsp.RC = RC_ERR_1005
		return
	}

	user := models.User{Mobile:pl.Mobile}
	err = models.GetUser(&user)
	if err != nil {
		rsp.RC = RC_ERR_1010
		return
	}
	if !strings.EqualFold(pl.IdCard,user.IdCard) {
		rsp.RC = RC_ERR_1031
		return
	}

	newpswd := util.BuildCaptcha()
	user.PassWord = util.StringMd5(newpswd)
	err = models.SetUserPasswd(&user)
	if err != nil {
		beego.Error("设置密码错误")
		rsp.RC = RC_ERR_1032
		return
	}
	if(beego.RunMode != "dev"){
		smsContent := fmt.Sprintf("您的新密码是 %s 请及时登陆并修改",newpswd)
		err = msg.SendSms(pl.Mobile,smsContent)
		if err != nil {
			beego.Info("发送失败")
			rsp.RC = RC_ERR_1003
			return
		}
	}

	beego.Debug("重置密码成功 -->",newpswd)
	return;
}

func (this * RanBaobaoController)ChangePswd(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse){
	if !this.validitySession(req.SID) {
		rsp.RC = RC_ERR_1012
		return
	}

	pl := changePswdPl{}
	err := util.ConvertToModel(&req.PL,&pl)
	if err != nil {
		beego.Error("conver err",err.Error())
		rsp.RC = RC_ERR_1001
		return
	}

	uid := session.GetSessionByiD(req.SID)
	user := &models.User{UID:uid}
	err = models.GetUser(user)
	if err != nil {
		rsp.RC = RC_ERR_1010
		return
	}

	if !strings.EqualFold(user.PassWord,util.StringMd5(pl.OldPswd)) {
		rsp.RC = RC_ERR_1011
		return
	}

	user.PassWord = util.StringMd5(pl.NewPswd)
	err = models.SetUserPasswd(user)
	if err != nil {
		beego.Error("设置密码错误")
		rsp.RC = RC_ERR_1033
		return
	}
	//删除session
	session.DeleteBySid(req.SID);
	return
}


func (this * RanBaobaoController)UpdateUserInfo(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse) {
	if !this.validitySession(req.SID) {
		rsp.RC = RC_ERR_1012
		return
	}
	pl := updateUserInfoPl{}
	err := util.ConvertToModel(&req.PL,&pl)
	if err != nil {
		beego.Error("conver err",err.Error())
		rsp.RC = RC_ERR_1001
		return
	}
	uid := session.GetSessionByiD(req.SID)
	user := &models.User{UID:uid}
	err = models.GetUser(user)
	if err != nil {
		rsp.RC = RC_ERR_1010
		return
	}

	if !util.ValidityIdCard(pl.RealName,pl.IdCard) {
		beego.Error("错误的身份证:%s != %s",pl.RealName,pl.IdCard)
		rsp.RC = RC_ERR_1005
		return
	}

	user.Mobile = pl.Mobile;
	user.RealName = pl.RealName;
	user.IdCard = pl.IdCard;
	user.AliPayAccount = pl.AliPayAccount;
	err = models.UpdateUserInfo(user)
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
			beego.Error("更新用户信息失败")
			rsp.RC = RC_ERR_1034
		}

		return
	}

	return
}