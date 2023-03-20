package logic

import (
	"main/dao/mysql"
	"main/models"
	"main/pkg/jwt"
	"main/pkg/snowflake"
)

func SignUp(p *models.ParamSignUpCheckInfo) (err error) {
	// 检查各个是否相同
	if err = mysql.CheckPhoneAndCode(p.Phone, p.Code); err != nil {
		return
	}
	// 生成UID
	userID := snowflake.GenID()
	// 构造一个user实例
	user := &models.User{
		UserId:   userID,
		UserName: p.Username,
		Password: p.Password,
		Phone:    p.Phone,
	}
	// 保存进数据库
	return mysql.InsertUser(user)
}

func CheckPhone(p *models.ParamSignUp) (err error) {
	// 检测手机号是否注册
	return mysql.CheckUserExist(p.Phone)
}

func SmsCode(code string, p *models.ParamSignUp) (err error) {
	// 保存进数据库
	return mysql.InsertCode(code, p)
}

func Login(p *models.ParamLogin) (id interface{}, token string, err error) {
	user := new(models.User)
	user.Password = p.Password
	user.Phone = p.Phone
	if err := mysql.Login(user); err != nil {
		return nil, "", err
	}
	// 生成jwt
	token, err = jwt.GenToken(user.UserId, user.UserName)
	return user.UserId, token, err
}
