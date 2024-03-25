package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DbConnect(dataSourceName string) {
	var err error
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("Error while connecting to database : ", err)
		return
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("Error ping action : ", err)
		return
	}

	fmt.Println("Database successfully connected!")
}
