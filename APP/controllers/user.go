package controllers

import (
	"errors"
	"fmt"
	"main/dao/mysql"
	"main/logic"
	"main/models"
	"main/pkg/sms"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 处理注册请求
func SignUpHandler(c *gin.Context) {
	// 1. 参数处理
	p := new(models.ParamSignUpCheckInfo)
	if err := c.ShouldBindJSON(p); err != nil {
		// 参数有误，直接返回
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是否是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2. 业务逻辑处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorPhoneAndCodeNotExist) {
			ResponseError(c, CodeOrPhoneFailed)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil, nil)
	fmt.Println(p)
}

func SmsHandler(c *gin.Context) {
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 参数有误，直接返回
		zap.L().Error("Sms with invalid param", zap.Error(err))
		// 判断err是否是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 手机号检验
	reg := "^1[3-9]{1}\\d{9}$"
	result := regexp.MustCompile(reg)
	ok := result.MatchString(p.Phone)
	if !ok {
		zap.L().Error("Phone number failed")
		ResponseError(c, CodePhoneFailed)
		return
	}
	if err := logic.CheckPhone(p); err != nil {
		zap.L().Error("Check phone failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 发送验证码
	IsTest := true
	code, err := sms.SmsResult(p, IsTest)
	if err != nil {
		zap.L().Error("Send code failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 保存进数据库
	if err := logic.SmsCode(code, p); err != nil {
		zap.L().Error("Keep data failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	if IsTest {
		ResponseSuccess(c, "测试阶段，就不直接向手机发送验证码了，验证码是"+code+"。用data代替一下吧（毕竟验证码有条数限制）", nil)
	} else {
		ResponseSuccess(c, nil, nil)
	}
	fmt.Println(p)
}

func LoginHandler(c *gin.Context) {
	// 1.
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 参数有误，直接返回
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是否是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2.
	id, token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("Login failed", zap.String("phone", p.Phone), zap.Error(err))
		if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.
	ResponseSuccess(c, token, id)
	fmt.Println(p)
}
