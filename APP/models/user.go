package models

type User struct {
	UserId   int64  `db:"user_id"`
	UserName string `db:"username"`
	Password string `db:"password"`
	Phone    string `db:"phone"`
	Gender   int    `db:"gender"`
}
