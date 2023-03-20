package mysql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

type Test_table struct {
	Id       int    `db:"test_id"`
	Name     string `db:"test_name"`
	Password string `db:"test_password"`
	Date     string `db:"test_date"`
}

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return err
	}
	// 最大闲置连接
	db.SetMaxIdleConns(viper.GetInt("max_idle_conns"))
	// 最大连接
	db.SetMaxOpenConns(viper.GetInt("max_open_conns"))
	return
}

func Close() {
	_ = db.Close()
}
