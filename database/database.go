package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)




func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ".aA1451418"
	dbName := "webtest_1"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db

}

