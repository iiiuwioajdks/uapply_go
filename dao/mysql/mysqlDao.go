package mysql

/*
	所有涉及db操作的业务代码都在这个包下实现
*/
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	// Connect 包括 Open 和 Ping
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return errors.Wrap(err, "connect error")
	}

	// 限流
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)

	return err
}

func Close() {
	err := db.Close()
	if err != nil {

	}
}
