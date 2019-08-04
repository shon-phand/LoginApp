package DB

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Start() *sql.DB {
	db, err := sql.Open("mysql", "root:Shon@2544@tcp(127.0.0.1:3306)/loginApp")
	fmt.Println("db", db, "err", err)
	return db
}
