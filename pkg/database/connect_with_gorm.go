package database

// import (
// 	"fmt"
// 	"strings"

// 	"gorm.io/gorm"
// )

// // Connect DB
// // Go IMPORT
// // "github.com/attapon-th/gorm_vertica/vertica"
// // "github.com/phuslu/log"
// // "github.com/spf13/viper"
// // "gorm.io/driver/clickhouse"
// // "gorm.io/driver/mysql"
// // "gorm.io/driver/postgres"
// // "gorm.io/driver/sqlserver"
// func ConnectDB(dsn string, driverName string) (db *gorm.DB, err error) {
// 	var ErrDriverNotSupport = fmt.Errorf("Driver: `%s` not supported.", driverName)
// 	if strings.Contains(driverName, "mysql") {
// 		// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	} else if strings.Contains(driverName, "postgres") {
// 		// db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	} else if strings.Contains(driverName, "ssqlserverql") {
// 		// db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

// 	} else if strings.Contains(driverName, "clickhouse") {
// 		// db, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{})

// 		//} else if strings.Contains(driverName, "vertica") {
// 		// db, err = gorm.Open(vertica.Open(dsn), &gorm.Config{})

// 	} else {
// 		// log.Error().Err(ErrDriverNotSupport).Msgf("Driver Not supported.")
// 		err = ErrDriverNotSupport
// 	}
// 	// if err == nil && viper.GetBool("DEV") {
// 	// 	db = db.Debug()
// 	// }
// 	return
// }
