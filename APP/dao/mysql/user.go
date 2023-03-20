package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"main/models"
)

const secret = "xxxiix"

func CheckUserExist(phone string) (err error) {
	sqlStr := `select 
	count(user_id) 
	from user 
	where phone=?`
	var count int
	if err := db.Get(&count, sqlStr, phone); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func CheckPhoneAndCode(phone string, code string) (err error) {
	sqlStr := `select 
	count(id) 
	from code 
	where phone=? and code=?`
	var count int
	if err := db.Get(&count, sqlStr, phone, code); err != nil {
		return err
	}
	if count == 0 {
		return ErrorPhoneAndCodeNotExist
	}
	return
}

func InsertUser(user *models.User) (err error) {
	// 密码加密
	user.Password = encryptPassword(user.Password)
	// 数据存入
	sqlStr := `insert into 
	user(user_id, username, password, phone)
	value(?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserId, user.UserName, user.Password, user.Phone)
	return
}

func InsertCode(code string, p *models.ParamSignUp) (err error) {
	sqlStr := `insert into 
	code(code, phone)
	value(?, ?)`
	_, err = db.Exec(sqlStr, code, p.Phone)
	return
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

func Login(user *models.User) (err error) {
	password := user.Password
	sqlStr := `select 
	password, user_id
	from user 
	where phone=?`
	if err = db.Get(user, sqlStr, user.Phone); err != nil {
		if err == sql.ErrNoRows {
			return ErrorInvalidPassword
		}
		// 数据库查询出问题
		return err
	}
	password = encryptPassword(password)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}
