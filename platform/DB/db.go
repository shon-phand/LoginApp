package DB

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Start() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:mysqlpasswd@tcp(127.0.0.1:3306)/loginApp")
	//fmt.Println("db", db, "err", err)
	return db, err
}
