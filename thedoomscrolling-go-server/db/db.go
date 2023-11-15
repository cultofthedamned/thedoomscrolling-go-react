package db

import (
	"database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:meowmeow@tcp(127.0.0.1:3306)/thedoomscrollingdb")

	if err != nil {
		panic(err.Error())
	}

	DB = db

	fmt.Println("Connected to Database")
}