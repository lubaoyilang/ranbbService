package msg
import (
	"github.com/astaxie/beego/httplib"
	"io/ioutil"
	"github.com/astaxie/beego"
	"encoding/xml"
	"errors"
)

type Result struct {
	Returnsms SMSResult `xml:"returnsms"`
}
type SMSResult struct {
	Returnstatus string `xml:"returnstatus"`
	Message string `xml:"message"`
	Remainpoint string `xml:"remainpoint"`
	TaskID string `xml:"taskID"`
	SuccessCounts string `xml:"successCounts"`
}

func SendSms(mobile,content string) error{
	req := httplib.Get(beego.AppConfig.String("smsUrl"))
	req.Header("apikey",beego.AppConfig.String("smsApiKey"))
	req.Param("mobile",mobile)
	beego.Info(content)
	req.Param("content",content+"【冉宝宝】")
	rsp,err := req.SendOut()
	if err != nil {
		beego.Error(err.Error())
	}
	data,err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		beego.Error(err.Error())
	}
	var result Result
	err = xml.Unmarshal(data,&result)
	if err != nil {
		beego.Error(err.Error())
		return err
	}
	if result.Returnsms.Returnstatus == "Success" {
		beego.Info(string(data))
		return errors.New("send captcha failed")
	}
	return nil
}