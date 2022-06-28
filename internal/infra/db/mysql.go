package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// TODO:
//  - configから設定できるようにする
//  - connection pool周りの設定を柔軟にする
func OpenDB(dbName string) (*sql.DB, error) {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: dbName,
	}

	db, openErr := sql.Open("mysql", cfg.FormatDSN())
	if openErr != nil {
		return nil, openErr
	}

	// 実際にコネクションが貼られるのはこのタイミング
	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	return db, nil
}
