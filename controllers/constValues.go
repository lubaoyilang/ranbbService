package controllers

const (
	CID_CAPTCHA_REQ = 91001
	CID_REGISTER_REQ = 91011
)

const (
	RC_OK = 0
	RC_ERR_UNKNOWN = -1
	RC_ERR_1001 = 1001 // 参数错误
	RC_ERR_1002 = 1002 //手机号码错误
	RC_ERR_1003 = 1003 //验证码发送失败
)
