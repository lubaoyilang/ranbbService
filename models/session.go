package models



type UserToken struct {
	SessionKey string `xorm:"'session_key' pk notnull unique varchar(64)"`
	SessionData string `xorm:"'session_data' varcha(32)"`
	SessionExpiry int64 `xorm:"'session_expiry'" BigInt(11)`
}

type Session struct  {
	SessionKey string `xorm:"'session_key' pk notnull unique varchar(64)"`
	SessionData string `xorm:"'session_data' blob"`
	SessionExpiry int64 `xorm:"'session_expiry'" BigInt(11)`
}
