package db

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "skm"
	dbPass := "SakibMulla12!@"
	dbName := "gomysql"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(166.62.29.42:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(500)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(10)
	db.Stats()
	return db
}





