package mysql

import "errors"

var (
	ErrorUserExist            = errors.New("用户已存在")
	ErrorPhoneAndCodeNotExist = errors.New("手机号或验证码错误")
	ErrorInvalidPassword      = errors.New("手机号或密码错误")
	ErrorInvalidID            = errors.New("无效的ID")
	ErrorBillNotExist         = errors.New("账单不存在")
)
