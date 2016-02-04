package controllers

const (
	CID_CAPTCHA_REQ = 91001
	CID_REGISTER_REQ = 91011
	CID_LOGIN_REQ = 91021
	CID_LOGOUT_REQ = 91031
	CID_USER_INFO_REQ = 91041
	CID_FIND_PASSWD_REQ = 91051
	CID_CHANGE_PASSWD_REQ = 91061
	CID_UPDATE_USERINFO_REQ = 91071

	CID_GET_GOODS_REQ = 91101
	CID_ACCEPT_GOODS_REQ = 91111
	CID_GET_GOODSCATEGROY_REQ = 91161

	CID_COMMIT_ORDER_REQ = 91121
	CID_DELETE_ORDER_REQ = 91131
	CID_GET_ORDER_INFO_REQ = 91141
	CID_GET_ORDER_LIST_REQ = 91151

	CID_GET_TBACC_LIST_REQ = 91201
	CID_UPDATE_TBACC_REQ = 91211
	CID_DELETE_TBACC_REQ = 91221
	CID_ADD_TBACC_REQ = 91231

	CID_GET_WALLET_LOGS_REQ =  91301
	CID_OUT_MOUNY_TO_ALP_REQ = 91311
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
	RC_ERR_1014 = 1014 //获取商品列表失败
	RC_ERR_1015 = 1015 //商品不存在 分类不存在
	RC_ERR_1016 = 1016 //商品已下架
	RC_ERR_1017 = 1017 //购买失败
	RC_ERR_1018 = 1018 //库存不足
	RC_ERR_1019 = 1019 //超过购买限额
	RC_ERR_1020 = 1020 //未找到订单
	RC_ERR_1021 = 1021 //订单状态不是正在刷单
	RC_ERR_1022 = 1022 //设置订单状态失败
	RC_ERR_1023 = 1023 //普通用户无法删除正在审核和有问题的订单
	RC_ERR_1024 = 1024 //删除订单失败
	RC_ERR_1025 = 1025 //获取订单列表失败
	RC_ERR_1026 = 1026 //获取淘宝账号列表
	RC_ERR_1027 = 1027 //更新淘宝账号失败
	RC_ERR_1028 = 1028 //删除淘宝账号失败
	RC_ERR_1029 = 1029 //添加淘宝账号失败
	RC_ERR_1030 = 1030 //获取钱包历史失败
	RC_ERR_1031 = 1031 //身份证号码不符
	RC_ERR_1032 = 1032 //查找密码错误
	RC_ERR_1033 = 1033 //设置新密码错误
	RC_ERR_1034 = 1034 //更新信息失败
	RC_ERR_1035 = 1035 //金额不足
	RC_ERR_1036 = 1036 //取现失败
)

