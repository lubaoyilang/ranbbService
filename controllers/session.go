package controllers
import (
	"github.com/go-xweb/log"
	"ranbbService/util"
	"ranbbService/models"
	"time"
	"strings"
	"github.com/henrylee2cn/pholcus/common/simplejson"
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

func (this * RanBaobaoController)GetCaptCha(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse) {
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

//	smsContent := fmt.Sprintf("您的验证码为%s \n",captcha)
//	err = msg.SendSms(pl.Mobile,smsContent)
//	if err != nil {
//		log.Info("发送失败")
//		rsp.RC = RC_ERR_1003
//		return
//	}

	rsp.RC = RC_OK
	rsp.PL = captcha
	servcache.Put(pl.Mobile,captcha,cacheTimeout)
	return
}

func (this * RanBaobaoController) Register(req * RanBaoBaoRequest,rsp * RanBaoBaoResponse) {

	pl := registerPl{}
	err := util.ConvertToModel(&req.PL,&pl)
	if err != nil {
		log.Error(err.Error())
		rsp.RC = RC_ERR_1001
		return
	}

	//验证码
	captcha := servcache.Get(pl.Mobile).(string)
	if !strings.EqualFold(captcha,pl.Captcha) {
		log.Errorf("验证码错误:%s != %s",captcha,pl.Captcha)
		rsp.RC = RC_ERR_1004
		return
	}

	if !util.ValidityIdCard(pl.RealName,pl.IdCard) {
		log.Errorf("错误的身份证:%s != %s",pl.RealName,pl.IdCard)
		rsp.RC = RC_ERR_1005
		return
	}

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
		log.Error("创建用户失败"+err.Error())
		rsp.RC = RC_ERR_1006
		return
	}
	json := simplejson.New()
	json.Set("UID",user.UID)
	rsp.PL = json
	return
}