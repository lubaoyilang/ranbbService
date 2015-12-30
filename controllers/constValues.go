package controllers

const (
	CID_CAPTCHA_REQ = 91001
	CID_REGISTER_REQ = 91011
	CID_LOGIN_REQ = 91021
	CID_LOGOUT_REQ = 91031
	CID_USER_INFO_REQ = 91041
)

const (
	RC_OK = 0
	RC_ERR_UNKNOWN = -1
	RC_ERR_1000 = 1000 //未知命令
	RC_ERR_1001 = 1001 // 参数错误
	RC_ERR_1002 = 1002 //手机号码错误
	RC_ERR_1003 = 1003 //验证码发送失败
	RC_ERR_1004 = 1004 //验证码错误
	RC_ERR_1005 = 1005 //错误的身份证
	RC_ERR_1006 = 1006 //创建用户失败
	RC_ERR_1007 = 1007 //手机号已经被注册
	RC_ERR_1008 = 1008 //身份证已经被使用
	RC_ERR_1009 = 1009 //支付宝账号已被使用
	RC_ERR_1010 = 1010 //用户不存在
	RC_ERR_1011 = 1011 //密码错误
	RC_ERR_1012 = 1012 //未登录
	RC_ERR_1013 = 1013 //登录失败
)

