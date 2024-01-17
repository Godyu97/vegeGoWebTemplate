package common

type WebError struct {
	code   string
	errMsg string
}

func (err *WebError) Error() string {
	if err != nil {
		return err.errMsg
	}
	return ""
}

func (err *WebError) Code() string {
	if err != nil {
		return err.code
	}
	return ""
}

func GetWebMsgCode(errMsg string) string {
	return errCodeM[errMsg]
}

func GetWebErrorFromErrMsg(errMsg string) *WebError {
	return &WebError{
		code:   GetWebMsgCode(errMsg),
		errMsg: errMsg,
	}
}

const (
	Nil                = "" //无msg，非正常情况，需要debug
	SuccessMsg         = "success"
	CheckAuthErrMsg    = "权限不足,请先登录"
	TokenRevokedErrMsg = "已在其他地方登陆"
	TokenExpireErrMsg  = "登陆过期"
	QueryBindErrMsg    = "query参数解析错误"
	ReqIllegal         = "非法的Req"
	AccountErrMsg      = "账号或密码错误"
	InternalErrMsg     = "服务内部错误"
	TokenRevokeMsg     = "token revoked"
	CodeErrMsg         = "验证码错误"
	CodeMinuteErrMsg   = "1分钟只能发送一条短信"
	MemHasErrMsg       = "用户已存在"
	DetailMaxErrMsg    = "已达每日阅读上限"
	// req check
	ReqAccountErrMsg   = "帐号不能为空"
	ReqMobileErrMsg    = "非法的手机号"
	ReqMobileHasErrMsg = "手机号已存在"
	ReqEmailErrMsg     = "非法的邮箱"
	ReqOverLen         = "参数长度过长"
	//wechat
	WechatOfficialErrMsg = "微信服务错误"
	WechatPayErrMsg      = "微信支付失败"
	WechatOrderReqErrMsg = "非法的订单请求"
	WechatQRLoginErrMsg  = "尚未微信扫码登录"
	//	SEO
	SEOErrMsg = "SEO500"
)

var errCodeM = map[string]string{
	Nil: "000000",
	//200
	SuccessMsg:           "200200",
	ReqIllegal:           "200202",
	AccountErrMsg:        "200203",
	TokenRevokeMsg:       "200205",
	CodeErrMsg:           "200206",
	ReqAccountErrMsg:     "200207",
	ReqMobileErrMsg:      "200208",
	ReqEmailErrMsg:       "200209",
	ReqOverLen:           "200210",
	WechatOfficialErrMsg: "200211",
	WechatPayErrMsg:      "200212",
	WechatOrderReqErrMsg: "200213",
	WechatQRLoginErrMsg:  "200214",
	MemHasErrMsg:         "200215",
	DetailMaxErrMsg:      "200216",
	CodeMinuteErrMsg:     "200217",
	ReqMobileHasErrMsg:   "200218",
	//401
	CheckAuthErrMsg:    "401233",
	TokenRevokedErrMsg: "401234",
	TokenExpireErrMsg:  "401235",
	//500
	QueryBindErrMsg: "500502",
	InternalErrMsg:  "500501",
	SEOErrMsg:       "500503",
}
