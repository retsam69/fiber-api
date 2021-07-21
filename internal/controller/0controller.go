package controller

import (
	"github.com/phuslu/log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func Init() {
	// DBConn = ConnectDB()

}

// Example Mysql
func ConnectDB() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := viper.GetString("db_dsn")
	// log.Debug().Msg(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("error config connect database: `MOPHIC`")
	}
	if viper.GetBool("dev") {
		db = db.Debug()
	}
	return db
}
