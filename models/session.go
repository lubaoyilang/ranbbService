package models



type Session struct {
	SessionKey string `xorm:"'session_key' pk notnull unique varchar(64)"`
	SessionData string `xorm:"'session_data' blob"`
	SessionExpiry int `xorm:"'session_expiry'" BigInt(11)`
}
