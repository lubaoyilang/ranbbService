package session
import (
	"ranbbService/models"
	"time"
	"github.com/astaxie/beego"
)


var sessioncache map[string]models.UserToken

func init() {
	sessioncache = make(map[string]models.UserToken)
}

func GetSessionByiD(sid string) string {
	var session models.UserToken

	if session,ok :=  sessioncache[sid];ok {
		return session.SessionData
	}

	session = models.UserToken{}
	sess := models.Engine.NewSession()
	defer  sess.Close()
	has,err := sess.Where("session_key = ?",sid).Get(&session)
	if err != nil || !has{
		return ""
	}

	sessioncache[session.SessionKey] = session
	return session.SessionData

}

func SetSession(sid,content string)  {
	session := models.UserToken{
		SessionKey:sid,
		SessionData:content,
		SessionExpiry:time.Now().Unix()}
	sess := models.Engine.NewSession()
	defer  sess.Close()
	sess.Begin()
	_,err := sess.Insert(session)
	if err != nil {
		defer sess.Rollback()
		return
	}
	sess.Commit()
	sessioncache[sid] = session
	return
}

func UpdateSession(sid string,content string) {
	session := models.UserToken{
		SessionKey:sid,
		SessionData:content,
		SessionExpiry:time.Now().Unix()}
	sess := models.Engine.NewSession()
	defer  sess.Close()
	sess.Begin()
	_,err := sess.Where("session_data = ?",content).Update(&session)
	if err != nil {
		sess.Rollback()
		return
	}
	sess.Commit()
	sessioncache[sid] = session
	return
}

func DeleteBySid(sid string) {
	sess := models.Engine.NewSession()
	defer  sess.Close()
	sess.Begin()
	session := models.UserToken{SessionKey:sid}
	_,err := sess.Table(new(models.UserToken)).Delete(&session)
	if err != nil {
		sess.Rollback()
		beego.Error(err.Error())
	}
	sess.Commit()
	delete(sessioncache,sid)
	beego.Debug("delete session ok")
}


func IsExist(content string) (*models.UserToken,bool ){
	session := &models.UserToken{}
	sess := models.Engine.NewSession()
	defer  sess.Close()
	has,err := sess.Where(`session_data = ?`,content).Get(session)
//	sess.Query(`SELECT session_key, session_data, session_expiry FROM session WHERE session_data = '76af69ff12e8da08fea73f6b33d52aaf' LIMIT 1 `)
	if err != nil || !has{
		beego.Error(err)
		return nil,false
	}
	sessioncache[session.SessionKey] = *session
	return session,true
}

