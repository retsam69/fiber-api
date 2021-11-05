package database

// import (
// 	"database/sql"

// 	"github.com/jmoiron/sqlx"
// 	_ "github.com/vertica/vertica-sql-go"
// )

// type DriverDB string

// const (
// 	DRIVER_VERTICA DriverDB = "vertica"
// )

// // https://github.com/vertica/vertica-sql-go
// // vertica://(user):(password)@(host):(port)/(database)?(queryArgs)
// func ConnectDatabase(driver DriverDB, ConnectString string) (dbConn *sqlx.DB, err error) {
// 	if d, er := sql.Open(string(driver), ConnectString); er != nil {
// 		return dbConn, er
// 	} else {
// 		dbConn = sqlx.NewDb(d, string(driver))
// 		err = dbConn.Ping()
// 	}
// 	return
// }
