package repo

// import (
// 	"github.com/phuslu/log"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var dbDefault *gorm.DB

// func GetDB() *gorm.DB {
// 	return dbDefault
// }

// func SetDB(dsn string) {
// 	myDial := mysql.Open(dsn)
// 	db, err := gorm.Open(myDial, &gorm.Config{})
// 	if err != nil {
// 		log.Fatal().Err(err).Msgf("Error Connect %s, Dsn: %s", myDial.Name(), dsn)
// 	}
// 	if loader.Dev {
// 		dbDefault = db.Debug()
// 	} else {
// 		dbDefault = db
// 	}
// }

// func Migrator(db *gorm.DB, dst ...interface{}) {
// 	if err := db.AutoMigrate(dst...); err != nil {
// 		log.Fatal().Err(err).Msg("Database `AutoMigrate` Error.")
// 	}
// }
