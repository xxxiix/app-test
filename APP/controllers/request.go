package controllers

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

var (
	ErrorUserNotLogin    = errors.New("用户未登录")
	ErrorUserFailedLogin = errors.New("用户信息错误")
)

// getCurrentUserID 获取当前登录的用户ID
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func CheckUser(c *gin.Context) (int64, error) {
	userID, _ := getCurrentUserID(c)
	user_id := fmt.Sprint(userID)
	if c.Param("user_id") != user_id {
		return 0, ErrorUserFailedLogin
	}
	return userID, nil
}
