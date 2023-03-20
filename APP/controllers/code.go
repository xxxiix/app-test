package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodePhoneFailed
	CodeOrPhoneFailed
	CodeInvalidPassword
	CodeServerBusy
	CodeBillNotExist

	CodeNeedLogin
	CodeInvalidToken
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodePhoneFailed:     "手机号错误",
	CodeOrPhoneFailed:   "手机号或验证码错误",
	CodeInvalidPassword: "手机号或密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeBillNotExist:    "账单不存在",

	CodeNeedLogin:    "需要登录",
	CodeInvalidToken: "无效的token",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
